package main

import (
	"fmt"
	"strings"
)

type lambdaFn struct {
	args    []string
	restArg string
	body    *ConsCell
	env     *env
}

var noRestArg string = ""

func mkLambda(cdr *ConsCell, e *env) (*lambdaFn, error) {
	args := []string{}
	restArg := noRestArg
	argList, ok := cdr.car.(*ConsCell)
	if !ok {
		return nil, fmt.Errorf("lambda requires an argument list")
	}
	emptyArgList := false
top:
	for argList != Nil && !emptyArgList {
		if argList.car == Nil {
			emptyArgList = true
		} else {
			arg, ok := argList.car.(Atom)
			if !ok {
				return nil, fmt.Errorf("argument list item is not an atom")
			}
			args = append(args, arg.s)
		}
		switch t := argList.cdr.(type) {
		case Atom:
			restArg = t.s
			break top
		case *ConsCell:
			argList = t
		default:
			// I was unable to reach this with a test:
			panic("unknown type in lambda arg list")
		}
	}
	if emptyArgList && restArg == noRestArg {
		return nil, fmt.Errorf("lambda with () argument requires a rest argument")
	}
	return &lambdaFn{args, restArg, cdr.cdr.(*ConsCell), e}, nil
}

func (f *lambdaFn) String() string {
	//restArgsRepr := ""
	// if f.restArg != noRestArgs {
	// 	restArgsRepr = "BOO" // fmt.Sprintf(" & %s", f.restArg)
	// }
	return fmt.Sprintf("<lambda(%s)>",
		strings.Join(f.args, " "),
	//	restArgsRepr
	)
}

func (f *lambdaFn) Equal(o Sexpr) bool {
	return false
}
