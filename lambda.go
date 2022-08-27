package main

import (
	"fmt"
	"strings"
)

type lambdaFn struct {
	args    []string
	restArg string
	body    *ConsCell
	doc     *ConsCell
	isMacro bool
	env     *env
}

var noRestArg string = ""

func mkLambda(cdr *ConsCell, isMacro bool, e *env) (*lambdaFn, error) {
	args := []string{}
	restArg := noRestArg
	// look for fn name
	if cdr == Nil {
		return nil, fmt.Errorf("missing arguments")
	}
	fnNameAtom, ok := cdr.car.(Atom)
	var fnName string
	if ok {
		fnName = fnNameAtom.s
		cdr = cdr.cdr.(*ConsCell)
	}
	if cdr == Nil {
		return nil, fmt.Errorf("missing arguments")
	}
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
	body := cdr.cdr.(*ConsCell)
	// Find `doc` form and save it if found:
	doc := Nil
	if body != Nil && body.car != Nil {
		doc2, ok := body.car.(*ConsCell)
		if ok && doc2.car.Equal(Atom{"doc"}) {
			doc = doc2.cdr.(*ConsCell)
			body = body.cdr.(*ConsCell) // Skip `doc` part.
		}
	}
	f := lambdaFn{args,
		restArg,
		body,
		doc,
		isMacro,
		e}
	if fnName != "" {
		// Monkey-patch the environment the lambda is created in, so the
		// lambda can invoke itself if the name is available:
		e.Set(fnName, &f)
	}
	return &f, nil
}

func (f *lambdaFn) String() string {
	restArgsRepr := ""
	if f.restArg != noRestArg {
		restArgsRepr = fmt.Sprintf(" . %s", f.restArg)
	}
	return fmt.Sprintf("<lambda(%s%s)>",
		strings.Join(f.args, " "), restArgsRepr)
}

func (f *lambdaFn) Equal(o Sexpr) bool {
	return false
}
