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
	sigEvaled, err := eval(sigExpr, e)
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
		_, err := eval(toEval, e)
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

// FIXME: This is only used by `(apply f ...)`.  DRY it out with
// respect to `eval`.
// applyFn applies a function to an already-evaluated list of arguments.
func applyFn(evalCar Sexpr, evaledList []Sexpr) (Sexpr, error) {
	// User-defined functions:
	lambda, ok := evalCar.(*lambdaFn)
	if ok {
		newEnv := mkEnv(lambda.env)
		if len(lambda.args) != len(evaledList) {
			return nil, fmt.Errorf("wrong number of args: %d != %d",
				len(lambda.args), len(evaledList))
		}
		for i, arg := range lambda.args {
			newEnv.Set(arg, evaledList[i])
		}
		if lambda.body == Nil {
			return Nil, nil
		}

		for {
			ret, err := eval(lambda.body.car, &newEnv)
			if err != nil {
				return nil, err
			}
			// Tail call, in sheep's clothing...
			if lambda.body.cdr == Nil {
				return ret, nil
			}
			lambda.body = lambda.body.cdr.(*ConsCell)
		}
	}
	// Built-in functions:
	builtin, ok := evalCar.(*Builtin)
	if !ok {
		return nil, fmt.Errorf("%s is not a function", evalCar)
	}
	biResult, err := builtin.Fn(evaledList)
	if err != nil {
		return nil, err
	}
	return biResult, nil

}

func evAtom(a Atom, e *env) (Sexpr, error) {
	if a.s == "t" {
		return a, nil
	}
	ret, ok := e.Lookup(a.s)
	if ok {
		return ret, nil
	}
	ret, ok = builtins[a.s]
	if ok {
		return ret, nil
	}
	return nil, fmt.Errorf("unknown symbol: %s", a.s)
}

func evDef(args *ConsCell, e *env) (Sexpr, error) {
	name := args.car.(Atom).s
	val, err := eval(args.cdr.(*ConsCell).car, e)
	if err != nil {
		panic(err)
	}
	e.Set(name, val)
	return val, nil
}

func eval(expr Sexpr, e *env) (Sexpr, error) {
top:
	switch t := expr.(type) {
	case Atom:
		return evAtom(t, e)
	case Number:
		return expr, nil
	case *ConsCell:
		if t == Nil {
			return Nil, nil
		}
		// special forms:
		if carAtom, ok := t.car.(Atom); ok {
			switch {
			case carAtom.s == "quote":
				return t.cdr.(*ConsCell).car, nil
			case carAtom.s == "cond":
				pairList := t.cdr.(*ConsCell)
				if pairList == Nil {
					return Nil, nil
				}
				for {
					if pairList == Nil {
						return Nil, nil
					}
					pair := pairList.car.(*ConsCell)
					ev, err := eval(pair.car, e)
					if err != nil {
						return nil, err
					}
					if ev == Nil {
						pairList = pairList.cdr.(*ConsCell)
						continue
					}
					// TAIL CALL!!!
					expr = pair.cdr.(*ConsCell).car
					goto top
				}
			case carAtom.s == "def":
				return evDef(t.cdr.(*ConsCell), e)
			case carAtom.s == "defn":
				return evDefn(t.cdr.(*ConsCell), e)
			case carAtom.s == "errors":
				return evErrors(t.cdr.(*ConsCell), e)
			case carAtom.s == "let":
				args := t.cdr.(*ConsCell)
				if args == Nil {
					return nil, fmt.Errorf("let requires a binding list")
				}
				bindings, ok := args.car.(*ConsCell)
				if !ok {
					return nil, fmt.Errorf("let bindings must be a list")
				}
				body := args.cdr.(*ConsCell)
				newEnv := mkEnv(e)
				for ; bindings != Nil; bindings = bindings.cdr.(*ConsCell) {
					binding, ok := bindings.car.(*ConsCell)
					if !ok {
						return nil, fmt.Errorf("a let binding must be a list")
					}
					name := binding.car.(Atom).s
					val, err := eval(binding.cdr.(*ConsCell).car, e)
					if err != nil {
						return nil, err
					}
					newEnv.Set(name, val)
				}

				var ret Sexpr = Nil
				for {
					var err error
					if body == Nil {
						return ret, nil
					}
					// Implement TCO for `let`:
					if body.cdr == Nil {
						expr = body.car
						e = &newEnv
						goto top
					}
					ret, err = eval(body.car, &newEnv)
					if err != nil {
						return nil, err
					}
					body = body.cdr.(*ConsCell)
				}
			case carAtom.s == "lambda":
				return mkLambda(t.cdr.(*ConsCell), e), nil
			}
		}
		// Functions / normal order of evaluation.  Get function to use first:
		evalCar, err := eval(t.car, e)
		if err != nil {
			return nil, err
		}
		// In normal function application, evaluate the arguments executing the
		// function:
		evaledList := []Sexpr{}
		start := t.cdr.(*ConsCell)
		for {
			if start == Nil {
				break
			}
			ee, err := eval(start.car, e)
			if err != nil {
				return nil, err
			}
			evaledList = append(evaledList, ee)
			start = start.cdr.(*ConsCell)
		}
		// User-defined functions:
		lambda, ok := evalCar.(*lambdaFn)
		if ok {
			newEnv := mkEnv(lambda.env)
			if len(lambda.args) != len(evaledList) {
				return nil, fmt.Errorf("wrong number of args: %d != %d",
					len(lambda.args), len(evaledList))
			}
			for i, arg := range lambda.args {
				newEnv.Set(arg, evaledList[i])
			}
			var ret Sexpr = Nil
			for {
				if lambda.body == Nil {
					return ret, nil
				}
				// TCO:
				if lambda.body.cdr == Nil {
					expr = lambda.body.car
					e = &newEnv
					goto top
				}
				ret, err = eval(lambda.body.car, &newEnv)
				if err != nil {
					return nil, err
				}
				lambda.body = lambda.body.cdr.(*ConsCell)
			}
		}
		// Built-in functions:
		builtin, ok := evalCar.(*Builtin)
		if !ok {
			return nil, fmt.Errorf("%s is not a function", evalCar)
		}
		biResult, err := builtin.Fn(evaledList)
		if err != nil {
			return nil, err
		}
		return biResult, nil
	default:
		panic(fmt.Sprintf("unknown type to eval: %T", t))
	}
}

func evDefn(args *ConsCell, e *env) (Sexpr, error) {
	if args == Nil {
		return nil, fmt.Errorf("defn requires a function name")
	}
	name, ok := args.car.(Atom)
	if !ok {
		return nil, fmt.Errorf("defn name must be an atom")
	}
	args = args.cdr.(*ConsCell)
	if args == Nil {
		return nil, fmt.Errorf("defn requires an argument list")
	}
	fn := mkLambda(args, e)
	e.Set(name.s, fn)
	return Nil, nil
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
