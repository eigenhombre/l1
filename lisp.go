package main

import (
	"fmt"

	"github.com/eigenhombre/lexutil"
)

// Sexpr is a general-purpose data structure for representing
// S-expressions; for now, it has String only, but it may have
// evaluable or other methods added later.
type Sexpr interface {
	String() string
	Eval(env env) Sexpr
}

// ConsCell is a cons cell.  Use Cons to create one.
type ConsCell struct {
	car Sexpr
	cdr Sexpr
}

type env map[string]Sexpr

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

func evlist(expr *ConsCell, e env) []Sexpr {
	ret := []Sexpr{}
	for ; expr != Nil; expr = expr.cdr.(*ConsCell) {
		ret = append(ret, expr.car.Eval(e))
	}
	return ret
}

// Eval a list expression
func (c *ConsCell) Eval(e env) Sexpr {
	if c == Nil {
		return Nil
	}
	if carAtom, ok := c.car.(Atom); ok {
		switch {
		case carAtom.s == "quote":
			return c.cdr.(*ConsCell).car
		case builtins[carAtom.s] != nil:
			return builtins[carAtom.s](evlist(c.cdr.(*ConsCell), e))
		default:
			panic("unimplemented")
		}
	}
	return Nil
}

// Atom is the primitive symbolic type.
type Atom struct {
	s string
}

func (a Atom) String() string {
	return a.s
}

// Eval for atom returns the atom if it's the truth value; otherwise, it looks
// up the value in the environment.
func (a Atom) Eval(e env) Sexpr {
	if a.s == "t" {
		return a
	}
	return e[a.s]
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

var builtins = map[string]func([]Sexpr) Sexpr{
	"+": func(args []Sexpr) Sexpr {
		if len(args) == 0 {
			return Num(0)
		}
		sum := Num(0)
		for _, arg := range args {
			sum = sum.Add(arg.(Number))
		}
		return sum
	},
	"-": func(args []Sexpr) Sexpr {
		if len(args) == 0 {
			panic("Handle me!")
		}
		if len(args) == 1 {
			return args[0].(Number).Neg()
		}
		sum := args[0].(Number)
		for _, arg := range args[1:] {
			sum = sum.Sub(arg.(Number))
		}
		return sum
	},
	"*": func(args []Sexpr) Sexpr {
		if len(args) == 0 {
			return Num(1)
		}
		prod := Num(1)
		for _, arg := range args {
			prod = prod.Mul(arg.(Number))
		}
		return prod
	},
	"/": func(args []Sexpr) Sexpr {
		if len(args) == 0 {
			panic("Handle me!")
		}
		if len(args) == 1 {
			return Num(1)
		}
		quot := args[0].(Number)
		for _, arg := range args[1:] {
			quot = quot.Div(arg.(Number))
		}
		return quot
	},
	"car": func(args []Sexpr) Sexpr {
		if len(args) != 1 {
			panic("Handle me!")
		}
		carCons, ok := args[0].(*ConsCell)
		if !ok {
			panic("Handle me!")
		}
		return carCons.car
	},
	"cdr": func(args []Sexpr) Sexpr {
		if len(args) != 1 {
			panic("Handle me!")
		}
		cdrCons, ok := args[0].(*ConsCell)
		if !ok {
			panic("Handle me!")
		}
		return cdrCons.cdr
	},
	"cons": func(args []Sexpr) Sexpr {
		if len(args) != 2 {
			panic("Handle me!")
		}
		return Cons(args[0], args[1].(*ConsCell))
	},
	"atom": func(args []Sexpr) Sexpr {
		if len(args) != 1 {
			panic("Handle me!")
		}
		_, ok := args[0].(Atom)
		if ok {
			return Atom{"t"}
		}
		return Nil
	},
	"eq": func(args []Sexpr) Sexpr {
		if len(args) != 2 {
			panic("Handle me!")
		}
		if args[0] == args[1] {
			return Atom{"t"}
		}
		return Nil
	},
}
