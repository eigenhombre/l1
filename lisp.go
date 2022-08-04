package main

import (
	"fmt"
	"strings"

	"github.com/eigenhombre/lexutil"
)

// Sexpr is a general-purpose data structure for representing
// S-expressions.
type Sexpr interface {
	String() string
	Eval(*env) (Sexpr, error)
	Equal(Sexpr) bool
}

// ConsCell is a cons cell.  Use Cons to create one.
type ConsCell struct {
	car Sexpr
	cdr Sexpr
}

// Nil is the empty list / cons cell.  Cons with Nil to create a list
// of one item.
var Nil *ConsCell = nil

func (c *ConsCell) String() string {
	ret := "("
	for car := c; car != Nil; car = car.cdr.(*ConsCell) {
		ret += car.car.String()
		if car.cdr != Nil {
			ret += " "
		}
	}
	return ret + ")"
}

// Cons creates a cons cell.
func Cons(i Sexpr, cdr *ConsCell) *ConsCell {
	return &ConsCell{i, cdr}
}

// Equal returns true iff the two S-expressions are equal cons-wise
func (c *ConsCell) Equal(o Sexpr) bool {
	_, ok := o.(*ConsCell)
	if !ok {
		return false
	}
	if c == Nil {
		return o == Nil
	}
	if o == Nil {
		return c == Nil
	}
	return c.car.Equal(o.(*ConsCell).car) && c.cdr.Equal(o.(*ConsCell).cdr)
}

func evList(expr *ConsCell, e *env) ([]Sexpr, error) {
	ret := []Sexpr{}
	for ; expr != Nil; expr = expr.cdr.(*ConsCell) {
		ee, err := expr.car.Eval(e)
		if err != nil {
			return nil, err
		}
		ret = append(ret, ee)
	}
	return ret, nil
}

func evCond(pairList *ConsCell, e *env) (Sexpr, error) {
	if pairList == Nil {
		return Nil, nil
	}
	pair := pairList.car.(*ConsCell)
	ev, err := pair.car.Eval(e)
	if err != nil {
		return nil, err
	}
	if ev == Nil {
		return evCond(pairList.cdr.(*ConsCell), e)
	}
	return pair.cdr.(*ConsCell).car.Eval(e)
}

func evDef(pair *ConsCell, e *env) Sexpr {
	name := pair.car.(Atom).s
	val, err := pair.cdr.(*ConsCell).car.Eval(e)
	if err != nil {
		panic(err)
	}
	e.Set(name, val)
	return val
}

func stringFromList(l *ConsCell) string {
	ret := []string{}
	for ; l != Nil; l = l.cdr.(*ConsCell) {
		ret = append(ret, l.car.String())
	}
	return strings.Join(ret, " ")
}

func evErrors(args *ConsCell, e *env) (Sexpr, error) {
	if args == Nil {
		return nil, fmt.Errorf("no error spec given")
	}
	sigExpr, ok := args.car.(*ConsCell)
	if !ok {
		return nil, fmt.Errorf("error signature must be a list")
	}
	sigEvaled, err := sigExpr.Eval(e)
	if err != nil {
		return nil, err
	}
	sigList, ok := sigEvaled.(*ConsCell)
	if !ok {
		return nil, fmt.Errorf("error signature must be a list")
	}
	errorStr := stringFromList(sigList)
	bodyArgs := args.cdr.(*ConsCell)
	for {
		if bodyArgs == Nil {
			return nil, fmt.Errorf("error not found")
		}
		toEval := bodyArgs.car
		_, err := toEval.Eval(e)
		if err != nil {
			if strings.Contains(err.Error(), errorStr) {
				return Nil, nil
			}
			return nil, fmt.Errorf("error '%s' not found in '%s'",
				errorStr, err.Error())
		}
		bodyArgs = bodyArgs.cdr.(*ConsCell)
	}
}

// applyFn applies a function to an already-evaluated list of arguments.
func applyFn(fnCar Sexpr, args []Sexpr) (Sexpr, error) {
	// User-defined functions:
	lambda, ok := fnCar.(*lambdaFn)
	if ok {
		return lambda.evLambda(args)
	}
	// Built-in functions:
	builtin, ok := fnCar.(*Builtin)
	if !ok {
		return nil, fmt.Errorf("%s is not a function", fnCar)
	}
	biResult, err := builtin.Fn(args)
	if err != nil {
		return nil, err
	}
	return biResult, nil

}

// Eval a list expression
func (c *ConsCell) Eval(e *env) (Sexpr, error) {
	if c == Nil {
		return Nil, nil
	}
	// special forms:
	if carAtom, ok := c.car.(Atom); ok {
		switch {
		case carAtom.s == "quote":
			return c.cdr.(*ConsCell).car, nil
		case carAtom.s == "cond":
			return evCond(c.cdr.(*ConsCell), e)
		case carAtom.s == "def":
			return evDef(c.cdr.(*ConsCell), e), nil
		case carAtom.s == "errors":
			return evErrors(c.cdr.(*ConsCell), e)
		case carAtom.s == "lambda":
			return mkLambda(c.cdr.(*ConsCell), e), nil
		}
	}
	// functions / normal order of evaluation:
	evalCar, err := c.car.Eval(e)
	if err != nil {
		return nil, err
	}
	evaledList, err := evList(c.cdr.(*ConsCell), e)
	if err != nil {
		return nil, err
	}
	return applyFn(evalCar, evaledList)
}

func balancedParenPoints(tokens []lexutil.LexItem) (int, int, error) {
	level := 0
	start := 0
	for i, token := range tokens[start:] {
		switch token.Typ {
		case itemLeftParen:
			level++
		case itemRightParen:
			level--
			if level == 0 {
				return 0, i, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("unbalanced parens")
}

func mkList(xs []Sexpr) *ConsCell {
	if len(xs) == 0 {
		return Nil
	}
	return Cons(xs[0], mkList(xs[1:]))
}

// parse returns a list of sexprs parsed from a list of tokens.
func parse(tokens []lexutil.LexItem) ([]Sexpr, error) {
	ret := []Sexpr{}
	i := 0
	for {
		if i >= len(tokens) {
			break
		}
		token := tokens[i]
		switch token.Typ {
		case itemNumber:
			ret = append(ret, Num(token.Val))
			i++
		case itemAtom:
			ret = append(ret, Atom{token.Val})
			i++
		case itemForwardQuote:
			i++ // skip ' token
			if i >= len(tokens) {
				return nil, fmt.Errorf("unexpected end of input")
			}
			var quoted Sexpr
			var delta int
			// quoted and delta depend on whether the quoted expression is
			// a list or an atom/num:
			if tokens[i].Typ != itemLeftParen {
				inner, err := parse(tokens[i : i+1])
				if err != nil {
					return nil, err
				}
				quoted = inner[0]
				delta = 1
			} else {
				start, end, err := balancedParenPoints(tokens[i:])
				if err != nil {
					return nil, err
				}
				inner, err := parse(tokens[i+start+1 : i+end])
				if err != nil {
					return nil, err
				}
				quoted = mkList(inner)
				delta = end - start + 1
			}
			i += delta
			quoteList := []Sexpr{Atom{"quote"}}
			quoteList = append(quoteList, quoted)
			ret = append(ret, mkList(quoteList))
		case itemLeftParen:
			start, end, err := balancedParenPoints(tokens[i:])
			if err != nil {
				return nil, err
			}
			inner, err := parse(tokens[i+start+1 : i+end])
			if err != nil {
				return nil, err
			}
			ret = append(ret, mkList(inner))
			i = i + end + 1
		case itemRightParen:
			return nil, fmt.Errorf("unexpected right paren")
		default:
			return nil, fmt.Errorf(token.Val)
		}
	}
	return ret, nil
}

func lexAndParse(s string) ([]Sexpr, error) {
	return parse(lexItems(s))
}
