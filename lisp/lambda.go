package lisp

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
	env     *Env
}

var noRestArg string = ""

func mkLambda(cdr *ConsCell, isMacro bool, e *Env) (*lambdaFn, error) {
	args := []string{}
	restArg := noRestArg
	// look for fn name
	if cdr == Nil {
		return nil, baseError("missing arguments")
	}
	fnNameAtom, ok := cdr.car.(Atom)
	var fnName string
	if ok {
		fnName = fnNameAtom.s
		cdr = cdr.cdr.(*ConsCell)
	}
	if cdr == Nil {
		return nil, baseError("missing arguments")
	}
	argList, ok := cdr.car.(*ConsCell)
	if !ok {
		return nil, baseError("lambda requires an argument list")
	}
	emptyArgList := false
top:
	for argList != Nil && !emptyArgList {
		if argList.car == Nil {
			emptyArgList = true
		} else {
			arg, ok := argList.car.(Atom)
			if !ok {
				return nil, baseError("argument list item is not an atom")
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
			return nil, baseError("unknown type in lambda arg list")
		}
	}
	if emptyArgList && restArg == noRestArg {
		return nil, baseError("lambda with () argument requires a rest argument")
	}
	body := cdr.cdr.(*ConsCell)
	// Find `doc` form and save it if found:
	doc := Nil
	if body != Nil && body.car != Nil {
		doc2, ok := body.car.(*ConsCell)
		if ok && doc2 != Nil && doc2.car.Equal(Atom{"doc"}) {
			cdrCons, ok := doc2.cdr.(*ConsCell)
			if !ok {
				return nil, baseError("doc form is not a list")
			}
			// Omit explicitly undocumented functions:
			if cdrCons != Nil && cdrCons.car != Nil {
				doc = cdrCons
			}
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
