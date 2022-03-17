package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Builtin represents a function with a native (Go) implementation.
type Builtin struct {
	Name string
	Fn   func([]Sexpr) (Sexpr, error)
}

func (b Builtin) String() string {
	return fmt.Sprintf("<builtin: %s>", b.Name)
}

// Eval for builtin returns itself.
func (b Builtin) Eval(e *env) (Sexpr, error) {
	return b, nil
}

var builtins = map[string]*Builtin{
	"+": {
		Name: "+",
		Fn: func(args []Sexpr) (Sexpr, error) {
			if len(args) == 0 {
				return Num(0), nil
			}
			sum := Num(0)
			for _, arg := range args {
				sum = sum.Add(arg.(Number))
			}
			return sum, nil
		},
	},
	"-": {
		Name: "-",
		Fn: func(args []Sexpr) (Sexpr, error) {
			if len(args) == 0 {
				return nil, fmt.Errorf("missing argument")
			}
			if len(args) == 1 {
				return args[0].(Number).Neg(), nil
			}
			sum := args[0].(Number)
			for _, arg := range args[1:] {
				sum = sum.Sub(arg.(Number))
			}
			return sum, nil
		},
	},
	"*": {
		Name: "*",
		Fn: func(args []Sexpr) (Sexpr, error) {
			if len(args) == 0 {
				return Num(1), nil
			}
			prod := Num(1)
			for _, arg := range args {
				prod = prod.Mul(arg.(Number))
			}
			return prod, nil
		},
	},
	"/": {
		Name: "/",
		Fn: func(args []Sexpr) (Sexpr, error) {
			if len(args) == 0 {
				return nil, fmt.Errorf("missing argument")
			}
			if len(args) == 1 {
				return Num(1), nil
			}
			quot := args[0].(Number)
			for _, arg := range args[1:] {
				if reflect.DeepEqual(arg, Num(0)) {
					return nil, fmt.Errorf("division by zero")
				}
				quot = quot.Div(arg.(Number))
			}
			return quot, nil
		},
	},
	"car": {
		Name: "car",
		Fn: func(args []Sexpr) (Sexpr, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("missing argument")
			}
			carCons, ok := args[0].(*ConsCell)
			if !ok {
				panic("Handle me!")
			}
			return carCons.car, nil
		},
	},
	"cdr": {
		Name: "cdr",
		Fn: func(args []Sexpr) (Sexpr, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("missing argument")
			}
			cdrCons, ok := args[0].(*ConsCell)
			if !ok {
				panic("Handle me!")
			}
			return cdrCons.cdr, nil
		},
	},
	"cons": {
		Name: "cons",
		Fn: func(args []Sexpr) (Sexpr, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("missing argument")
			}
			return Cons(args[0], args[1].(*ConsCell)), nil
		},
	},
	"atom": {
		Name: "atom",
		Fn: func(args []Sexpr) (Sexpr, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("missing argument")
			}
			_, ok := args[0].(Atom)
			if ok {
				return Atom{"t"}, nil
			}
			return Nil, nil
		},
	},
	"eq": {
		Name: "eq",
		Fn: func(args []Sexpr) (Sexpr, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("missing argument")
			}
			if args[0] == args[1] {
				return Atom{"t"}, nil
			}
			return Nil, nil
		},
	},
	"print": {
		Name: "print",
		Fn: func(args []Sexpr) (Sexpr, error) {
			strArgs := []string{}
			for _, arg := range args {
				strArgs = append(strArgs, arg.String())
			}
			fmt.Println(strings.Join(strArgs, " "))
			return Nil, nil
		},
	},
}
