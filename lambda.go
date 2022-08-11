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

func mkLambda(cdr *ConsCell, e *env) *lambdaFn {
	args := []string{}
	argList := cdr.car.(*ConsCell)
	for ; argList != Nil; argList = argList.cdr.(*ConsCell) {
		args = append(args, argList.car.(Atom).s)
	}
	return &lambdaFn{args, noRestArg, cdr.cdr.(*ConsCell), e}
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
