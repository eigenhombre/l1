package main

import (
	"fmt"
	"math/big"
)

// sexpr is a general-purpose data structure for representing
// S-expressions; for now, it has String only, but it may have
// evaluable or other methods added later.
type sexpr interface {
	String() string
}

// ConsCell is a cons cell.  Use Cons to create one.
type ConsCell struct {
	car sexpr
	cdr sexpr
}

type env map[string]sexpr

// Nil is the empty list / cons cell.  Cons with Nil to create a list
// of one item.
var Nil *ConsCell = nil

func (j *ConsCell) String() string {
	ret := "("
	for car := j; car != Nil; car = car.cdr.(*ConsCell) {
		if car.car == Nil {
			return ret + ")"
		}
		ret += car.car.String()
		if car.cdr != Nil {
			ret += " "
		}
	}
	return ret + ")"
}

// Number is a big.Int, but narrower in its string representation.
type Number struct {
	big.Int
}

// String returns the string representation of the number.
func (n Number) String() string {
	return n.Text(10)
}

// Num is a `number` constructor, which can take a string or a
// ("normal" number).
func Num(ob interface{}) Number {
	var n Number
	switch s := ob.(type) {
	case string:
		n.SetString(s, 10)
	case int:
		n.SetInt64(int64(s))
	default:
		panic(fmt.Sprintf("Num: unknown type %T\n", ob))
	}
	return n
}

// Cons creates a cons cell.
func Cons(i sexpr, cdr *ConsCell) *ConsCell {
	return &ConsCell{i, cdr}
}

// Atom is the primitive symbolic type.
type Atom struct {
	s string
}

func (a Atom) String() string {
	return a.s
}

func balancedParenPoints(tokens []item) (int, int, error) {
	level := 0
	start := 0
	for i, token := range tokens[start:] {
		switch token.typ {
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

func mkList(xs []sexpr) *ConsCell {
	if len(xs) == 0 {
		return Nil
	}
	return Cons(xs[0], mkList(xs[1:]))
}

// parse returns a list of sexprs parsed from a list of tokens.
func parse(tokens []item) ([]sexpr, error) {
	ret := []sexpr{}
	i := 0
	for {
		if i >= len(tokens) {
			break
		}
		token := tokens[i]
		switch token.typ {
		case itemNumber:
			ret = append(ret, Num(token.val))
			i++
		case itemAtom:
			ret = append(ret, Atom{token.val})
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

func lexAndParse(s string) ([]sexpr, error) {
	return parse(lexItems(s))
}

func evalCons(expr *ConsCell, e env) sexpr {
	if expr == Nil {
		return Nil
	}
	if carAtom, ok := expr.car.(Atom); ok {
		if carAtom.s == "quote" {
			return expr.cdr.(*ConsCell).car
		}
	}
	return Nil
}

func eval(expr sexpr, e env) sexpr {
	switch expr := expr.(type) {
	case *ConsCell:
		return evalCons(expr, e)
	case Number:
		return expr
	case Atom:
		if expr.s == "t" {
			return expr
		}
		return e[expr.s]
	default:
		panic(fmt.Sprintf("eval: unknown type %T\n", expr))
	}
}
