package main

import (
	"fmt"
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

// Equal returns true if the receiver and the arg are both builtins and have the
// same name.
func (b Builtin) Equal(o Sexpr) bool {
	if o, ok := o.(Builtin); ok {
		return b.Name == o.Name
	}
	return false
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
				n, ok := arg.(Number)
				if !ok {
					return nil, fmt.Errorf("expected number, got '%s'", arg)
				}
				sum = sum.Add(n)
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
			sum, ok := args[0].(Number)
			if !ok {
				return nil, fmt.Errorf("expected number, got '%s'", args[0])
			}
			if len(args) == 1 {
				return args[0].(Number).Neg(), nil
			}
			for _, arg := range args[1:] {
				n, ok := arg.(Number)
				if !ok {
					return nil, fmt.Errorf("expected number, got '%s'", arg)
				}
				sum = sum.Sub(n)
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
				n, ok := arg.(Number)
				if !ok {
					return nil, fmt.Errorf("expected number, got '%s'", arg)
				}
				prod = prod.Mul(n)
			}
			return prod, nil
		},
	},
	"/": {
		Name: "/",
		Fn: func(args []Sexpr) (Sexpr, error) {
			if len(args) < 2 {
				return nil, fmt.Errorf("missing argument")
			}
			quot, ok := args[0].(Number)
			if !ok {
				return nil, fmt.Errorf("expected number, got '%s'", args[0])
			}
			for _, arg := range args[1:] {
				if arg.Equal(Num(0)) {
					return nil, fmt.Errorf("division by zero")
				}
				n, ok := arg.(Number)
				if !ok {
					return nil, fmt.Errorf("expected number, got '%s'", arg)
				}
				quot = quot.Div(n)
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
				return nil, fmt.Errorf("'%s' is not a list", args[0])
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
				return nil, fmt.Errorf("'%s' is not a list", args[0])
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
			if args[0].Equal(args[1]) {
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
