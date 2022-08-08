package main

import (
	"fmt"
	"strings"
)

type lambdaFn struct {
	args []string
	body *ConsCell
	env  *env
}

func mkLambda(cdr *ConsCell, e *env) *lambdaFn {
	args := []string{}
	argList := cdr.car.(*ConsCell)
	for ; argList != Nil; argList = argList.cdr.(*ConsCell) {
		args = append(args, argList.car.(Atom).s)
	}
	return &lambdaFn{args, cdr.cdr.(*ConsCell), e}
}

func (f *lambdaFn) String() string {
	return fmt.Sprintf("<lambda(%s)>", strings.Join(f.args, " "))
}

func (f *lambdaFn) Equal(o Sexpr) bool {
	return false
}
