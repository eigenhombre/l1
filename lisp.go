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

// consCell is a cons cell.  Use Cons to create one.
type consCell struct {
	car sexpr
	cdr sexpr
}

// Nil is the empty list / cons cell.  Cons with Nil to create a list
// of one item.
var Nil *consCell = nil

func (j *consCell) String() string {
	ret := "("
	for car := j; car != Nil; car = car.cdr.(*consCell) {
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

// number is a big.Int, but narrower in its string representation.
type number struct {
	big.Int
}

// String returns the string representation of the number.
func (n number) String() string {
	return n.Text(10)
}

func Num(ob interface{}) number {
	var n number
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

func Cons(i sexpr, cdr *consCell) *consCell {
	return &consCell{i, cdr}
}

type atom struct {
	s string
}

// Atom makes an atom from a string.  It doesn't check for lexical correctness
// -- the lexer should do that.
func Atom(s string) atom {
	return atom{s}
}

func (a atom) String() string {
	return a.s
}

func balancedParenPoints(tokens []item) (int, int) {
	level := 0
	start := 0
	for i, token := range tokens[start:] {
		switch token.typ {
		case itemLeftParen:
			level++
		case itemRightParen:
			level--
			if level == 0 {
				return 0, i
			}
		}
	}
	panic(fmt.Sprintf("balancedParenPoints: unbalanced parens in %q", tokens))
}

func mkList(xs []sexpr) *consCell {
	if len(xs) == 0 {
		return Nil
	}
	return Cons(xs[0], mkList(xs[1:]))
}

// parse returns a list of sexprs parsed from a list of tokens.
func parse(tokens []item) []sexpr {
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
			ret = append(ret, Atom(token.val))
			i++
		case itemLeftParen:
			start, end := balancedParenPoints(tokens[i:])
			ret = append(ret, mkList(parse(tokens[i+start+1:i+end])))
			i = i + end + 1
		default:
			panic(fmt.Sprintf("strToSexpr: bad token type %d ", token.typ))
		}
	}
	return ret
}

func lexAndParse(s string) []sexpr {
	return parse(lexItems(s))
}

func eval(s sexpr) sexpr {
	switch s.(type) {
	case *consCell:
		// Work to be done here
		return Nil
	case number:
		return s
	// case atom:
	// Finish this when we have environments
	default:
		panic(fmt.Sprintf("eval: unknown type %T\n", s))
	}
}
