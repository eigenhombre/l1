package main

import (
	"fmt"

	"github.com/eigenhombre/lexutil"
)

// Sexpr is a general-purpose data structure for representing
// S-expressions.
type Sexpr interface {
	String() string
	Eval(*env) (Sexpr, error)
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
		}
	}
	// functions / normal order of evaluation:
	evalCar, err := c.car.Eval(e)
	if err != nil {
		return nil, err
	}
	el, err := evList(c.cdr.(*ConsCell), e)
	if err != nil {
		return nil, err
	}
	biResult, err := evalCar.(*Builtin).Fn(el)
	if err != nil {
		return nil, err
	}
	return biResult, nil
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
			return nil, fmt.Errorf("unexpected token %v", token)
		}
	}
	return ret, nil
}

func lexAndParse(s string) ([]Sexpr, error) {
	return parse(lexItems(s))
}
