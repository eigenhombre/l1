package main

import (
	"fmt"
	"strings"
)

// Sexpr is a general-purpose data structure for representing
// S-expressions.
type Sexpr interface {
	String() string
	Equal(Sexpr) bool
}

func stringFromList(l *ConsCell) string {
	ret := []string{}
	for ; l != Nil; l = l.cdr.(*ConsCell) {
		ret = append(ret, l.car.String())
	}
	return strings.Join(ret, " ")
}

func evAtom(a Atom, e *env) (Sexpr, error) {
	if a.s == "t" {
		return a, nil
	}
	ret, ok := e.Lookup(a.s)
	if ok {
		return ret, nil
	}
	ret, ok = builtins[a.s]
	if ok {
		return ret, nil
	}
	return nil, fmt.Errorf("unknown symbol: %s", a.s)
}

func evDef(args *ConsCell, e *env) (Sexpr, error) {
	name := args.car.(Atom).s
	val, err := eval(args.cdr.(*ConsCell).car, e)
	if err != nil {
		return nil, err
	}
	err = e.Set(name, val)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func evDefn(args *ConsCell, isMacro bool, e *env) (Sexpr, error) {
	errPreamble := "defn"
	if isMacro {
		errPreamble = "defmacro"
	}

	if args == Nil {
		return nil, fmt.Errorf("%s requires a function name", errPreamble)
	}
	name, ok := args.car.(Atom)
	if !ok {
		return nil, fmt.Errorf("%s name must be an atom", errPreamble)
	}
	args = args.cdr.(*ConsCell)
	if args == Nil {
		return nil, fmt.Errorf("%s requires an argument list", errPreamble)
	}
	fn, err := mkLambda(args, isMacro, e)
	if err != nil {
		return nil, err
	}
	err = e.Set(name.s, fn)
	if err != nil {
		return nil, err
	}
	return Nil, nil
}

func evErrors(args *ConsCell, e *env) (Sexpr, error) {
	if args == Nil {
		return nil, fmt.Errorf("no error spec given")
	}
	sigExpr, ok := args.car.(*ConsCell)
	if !ok {
		return nil, fmt.Errorf("error signature must be a list")
	}
	sigEvaled, err := eval(sigExpr, e)
	if err != nil {
		return nil, err
	}
	sigList, ok := sigEvaled.(*ConsCell)
	if !ok {
		return nil, fmt.Errorf("error signature must be a list")
	}
	errorStr := stringFromList(sigList)
	bodyArgs := args.cdr.(*ConsCell)
	for {
		if bodyArgs == Nil {
			return nil, fmt.Errorf("error not found in %s", args)
		}
		toEval := bodyArgs.car
		_, err := eval(toEval, e)
		if err != nil {
			if strings.Contains(err.Error(), errorStr) {
				return Nil, nil
			}
			return nil, fmt.Errorf("error '%s' not found in '%s'",
				errorStr, err.Error())
		}
		bodyArgs = bodyArgs.cdr.(*ConsCell)
	}
}

// Both eval and apply use this to bind lambda arguments in the
// supplied environment:
func setLambdaArgsInEnv(newEnv *env, lambda *lambdaFn, evaledList []Sexpr) error {
	var err error
	if lambda.restArg != noRestArg {
		if len(lambda.args) > len(evaledList) {
			return fmt.Errorf("not enough arguments for function")
		}
		err = newEnv.Set(lambda.restArg,
			mkListAsConsWithCdr(evaledList[len(lambda.args):],
				Nil))
		if err != nil {
			return err
		}
	} else {
		if len(lambda.args) < len(evaledList) {
			return fmt.Errorf("too many arguments for function")
		} else if len(lambda.args) > len(evaledList) {
			return fmt.Errorf("not enough arguments for function")
		}
	}
	for i, arg := range lambda.args {
		err := newEnv.Set(arg, evaledList[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func isMacroCall(args Sexpr, e *env) bool {
	if args == Nil {
		return false
	}
	argsCons, ok := args.(*ConsCell)
	if !ok {
		return false
	}
	fn, ok := argsCons.car.(Atom)
	if !ok {
		return false
	}
	item, found := e.Lookup(fn.s)
	if !found {
		return false
	}
	f, ok := item.(*lambdaFn)
	if !ok {
		return false
	}
	return f.isMacro
}

func macroexpand1(expr Sexpr, e *env) (Sexpr, error) {
	if !isMacroCall(expr, e) {
		return expr, nil
	}
	fn, _ := e.Lookup(expr.(*ConsCell).car.(Atom).s)
	c, ok := expr.(*ConsCell)
	if !ok {
		return nil, fmt.Errorf("macro call must be a list")
	}
	if err := setLambdaArgsInEnv(e,
		fn.(*lambdaFn),
		consToExprs(c.cdr)); err != nil {
		return nil, err
	}
	ast := fn.(*lambdaFn).body
	toEval := ast.car
	ret, err := eval(toEval, e)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func macroexpand(expr Sexpr, e *env) (Sexpr, error) {
	var ret Sexpr = expr
	var err error
	for {
		ret, err = macroexpand1(ret, e)
		if err != nil {
			return nil, err
		}
		if !isMacroCall(ret, e) {
			return ret, nil
		}
	}
}

func listStartsWith(expr *ConsCell, s string) bool {
	if expr == Nil {
		return false
	}
	car, ok := expr.car.(Atom)
	if !ok {
		return false
	}
	return car.s == s
}

// Adapted from https://github.com/kanaka/mal/blob/master/impls/go/src/step7_quote/step7_quote.go#L36:
// FIXME: Rather than converting to slice and back, could recursively build up from the bottom.
func transformSyntaxQuoteList(l *ConsCell) (*ConsCell, error) {
	tList := consToExprs(l)
	ret := Nil
	for i := len(tList) - 1; 0 <= i; i-- {
		elt := tList[i]
		switch t := elt.(type) {
		case *ConsCell:
			if listStartsWith(t, "splicing-unquote") {
				ret = Cons(Atom{"concat"}, Cons(t.cdr.(*ConsCell).car, Cons(ret, Nil)))
				continue
			}
		default:
		}
		ret = Cons(Atom{"cons"}, Cons(syntaxQuote(elt), Cons(ret, Nil)))
	}
	return ret, nil
}

func syntaxQuote(arg Sexpr) Sexpr {
	switch t := arg.(type) {
	case Number, Atom:
		return Cons(Atom{"quote"}, Cons(arg, Nil))
	case *ConsCell:
		if listStartsWith(t, "unquote") {
			return t.cdr.(*ConsCell).car
		}
		qql, err := transformSyntaxQuoteList(t)
		if err != nil {
			return nil
		}
		return qql
	default:
		return t
	}
}

func eval(exprArg Sexpr, e *env) (Sexpr, error) {
	expr := exprArg
	var err error
top:
	if isMacroCall(expr, e) {
		expr, err = macroexpand(expr, e)
		if err != nil {
			return nil, err
		}
	}
	switch t := expr.(type) {
	case Atom:
		return evAtom(t, e)
	case Number:
		return expr, nil
	case *ConsCell:
		if t == Nil {
			return Nil, nil
		}
		// special forms:
		if carAtom, ok := t.car.(Atom); ok {
			switch {
			case carAtom.s == "quote":
				return t.cdr.(*ConsCell).car, nil
			case carAtom.s == "syntax-quote":
				expr = syntaxQuote(t.cdr.(*ConsCell).car)
				goto top
			case carAtom.s == "cond":
				pairList := t.cdr.(*ConsCell)
				if pairList == Nil {
					return Nil, nil
				}
				for {
					if pairList == Nil {
						return Nil, nil
					}
					pair, ok := pairList.car.(*ConsCell)
					if !ok || pair == Nil {
						return nil, fmt.Errorf("cond requires a list of pairs")
					}
					ev, err := eval(pair.car, e)
					if err != nil {
						return nil, err
					}
					if ev == Nil {
						pairList = pairList.cdr.(*ConsCell)
						continue
					}
					// TAIL CALL!!!
					cdrCons, ok := pair.cdr.(*ConsCell)
					if !ok || cdrCons == Nil {
						return nil, fmt.Errorf("cond requires a list of pairs")
					}
					expr = cdrCons.car
					goto top
				}
			// FIXME: Do as a macro:
			case carAtom.s == "and":
				pairList, ok := t.cdr.(*ConsCell)
				if !ok {
					return nil, fmt.Errorf("and requires a list of expressions")
				}
				for {
					if pairList == Nil {
						return True, nil
					}
					ev, err := eval(pairList.car, e)
					if err != nil {
						return nil, err
					}
					if ev == Nil {
						return Nil, nil
					}
					pairList, ok = pairList.cdr.(*ConsCell)
					if !ok {
						return nil, fmt.Errorf("and requires a list of expressions")
					}
				}
			// FIXME: Do as a macro:
			case carAtom.s == "or":
				pairList, ok := t.cdr.(*ConsCell)
				if !ok {
					return nil, fmt.Errorf("or requires a list of expressions")
				}
				for {
					if pairList == Nil {
						return Nil, nil
					}
					ev, err := eval(pairList.car, e)
					if err != nil {
						return nil, err
					}
					if ev != Nil {
						return ev, nil
					}
					pairList, ok = pairList.cdr.(*ConsCell)
					if !ok {
						return nil, fmt.Errorf("or requires a list of expressions")
					}
				}
			case carAtom.s == "loop":
				body, ok := t.cdr.(*ConsCell)
				if !ok {
					return nil, fmt.Errorf("loop requires a body")
				}
				for {
					start := body
				bodyLoop:
					for {
						if start == Nil {
							break bodyLoop
						}
						_, err := eval(start.car, e)
						if err != nil {
							return nil, err
						}
						start = start.cdr.(*ConsCell)
					}
				}
			case carAtom.s == "def":
				return evDef(t.cdr.(*ConsCell), e)
			case carAtom.s == "defn":
				return evDefn(t.cdr.(*ConsCell), false, e)
			case carAtom.s == "defmacro":
				return evDefn(t.cdr.(*ConsCell), true, e)
			case carAtom.s == "errors":
				return evErrors(t.cdr.(*ConsCell), e)
			case carAtom.s == "let":
				args := t.cdr.(*ConsCell)
				if args == Nil {
					return nil, fmt.Errorf("let requires a binding list")
				}
				bindings, ok := args.car.(*ConsCell)
				if !ok {
					return nil, fmt.Errorf("let bindings must be a list")
				}
				body := args.cdr.(*ConsCell)
				newEnv := mkEnv(e)
				for ; bindings != Nil; bindings = bindings.cdr.(*ConsCell) {
					binding, ok := bindings.car.(*ConsCell)
					if !ok {
						return nil, fmt.Errorf("a let binding must be a list")
					}
					name := binding.car.(Atom).s
					val, err := eval(binding.cdr.(*ConsCell).car, e)
					if err != nil {
						return nil, err
					}
					err = newEnv.Set(name, val)
					if err != nil {
						return nil, err
					}
				}

				var ret Sexpr = Nil
				for {
					var err error
					if body == Nil {
						return ret, nil
					}
					// Implement TCO for `let`:
					if body.cdr == Nil {
						expr = body.car
						e = &newEnv
						goto top
					}
					ret, err = eval(body.car, &newEnv)
					if err != nil {
						return nil, err
					}
					body = body.cdr.(*ConsCell)
				}
			case carAtom.s == "lambda":
				return mkLambda(t.cdr.(*ConsCell), false, e)
			}
		}
		// Functions / normal order of evaluation.  Get function to use first:
		evalCar, err := eval(t.car, e)
		if err != nil {
			return nil, err
		}
		// In normal function application, evaluate the arguments executing the
		// function:
		evaledList := []Sexpr{}
		start := t.cdr.(*ConsCell)
		for {
			if start == Nil {
				break
			}
			ee, err := eval(start.car, e)
			if err != nil {
				return nil, err
			}
			evaledList = append(evaledList, ee)
			start = start.cdr.(*ConsCell)
		}
		// User-defined functions:
		lambda, ok := evalCar.(*lambdaFn)
		if ok {
			var err error
			newEnv := mkEnv(lambda.env)
			err = setLambdaArgsInEnv(&newEnv, lambda, evaledList)
			if err != nil {
				return nil, err
			}
			var ret Sexpr = Nil
			for {
				if lambda.body == Nil {
					return ret, nil
				}
				// TCO:
				if lambda.body.cdr == Nil {
					expr = lambda.body.car
					e = &newEnv
					goto top
				}
				ret, err = eval(lambda.body.car, &newEnv)
				if err != nil {
					return nil, err
				}
				lambda.body = lambda.body.cdr.(*ConsCell)
			}
		}
		// Built-in functions:
		builtin, ok := evalCar.(*Builtin)
		if !ok {
			return nil, fmt.Errorf("%s is not a function", evalCar)
		}
		biResult, err := builtin.Fn(evaledList, e)
		if err != nil {
			return nil, err
		}
		return biResult, nil
	default:
		panic(fmt.Sprintf("unknown type to eval: %T", t))
	}
}
