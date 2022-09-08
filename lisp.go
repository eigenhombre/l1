package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Sexpr is a general-purpose interface for representing
// S-expressions.
type Sexpr interface {
	String() string
	Equal(Sexpr) bool
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
	return nil, baseErrorf("unknown symbol: %s", a.s)
}

// Do this once, it's pretty expensive:
var cxrRe = regexp.MustCompile(`^c([ad]+)r$`)

func isCxr(a Atom) bool {
	return cxrRe.MatchString(a.s) && a.s != "car" && a.s != "cdr"
}

func extractCxrLambda(t Atom, e *env) (Sexpr, error) {
	args := Nil
	for i := 1; i < len(t.s)-1; i++ {
		// isCxr guarantees that the string is of the form c[ad]+r, so
		// runes are either 'a' or 'd', of length 1:
		args = Cons(Atom{string(t.s[i])}, args)
	}
	newLambda := &lambdaFn{
		args:    []string{"xs"},
		body:    list(list(Atom{"c*r"}, list(Atom{"quote"}, args), Atom{"xs"})),
		isMacro: false,
		env:     e,
	}
	return newLambda, nil
}

func evDef(args *ConsCell, e *env) (Sexpr, error) {
	if args == Nil {
		return nil, baseError("missing argument")
	}
	carAtom, ok := args.car.(Atom)
	if !ok {
		return nil, baseError("first argument must be an atom")
	}
	name := carAtom.s
	args, ok = args.cdr.(*ConsCell)
	if !ok || args == Nil {
		return nil, baseError("missing argument")
	}
	val, err := eval(args.car, e)
	if err != nil {
		return nil, extendError("evaluating def value", err)
	}
	err = e.SetTopLevel(name, val)
	if err != nil {
		return nil, extendError("setting def result", err)
	}
	return val, nil
}

func evSet(args *ConsCell, e *env) (Sexpr, error) {
	if args == Nil {
		return nil, baseError("missing argument")
	}
	carAtom, ok := args.car.(Atom)
	if !ok {
		return nil, baseError("first argument must be an atom")
	}
	name := carAtom.s
	args, ok = args.cdr.(*ConsCell)
	if !ok || args == Nil {
		return nil, baseError("missing argument")
	}
	val, err := eval(args.car, e)
	if err != nil {
		return nil, extendError("evaluating set value", err)
	}
	err = e.Update(name, val)
	if err != nil {
		return nil, extendError("updating set result", err)
	}
	return val, nil
}

func evDefn(args *ConsCell, isMacro bool, e *env) (Sexpr, error) {
	errPreamble := "defn"
	if isMacro {
		errPreamble = "defmacro"
	}

	if args == Nil {
		return nil, baseErrorf("%s requires a function name", errPreamble)
	}
	name, ok := args.car.(Atom)
	if !ok {
		return nil, baseErrorf("%s name must be an atom", errPreamble)
	}
	args = args.cdr.(*ConsCell)
	if args == Nil {
		return nil, baseErrorf("%s requires an argument list", errPreamble)
	}
	fn, err := mkLambda(args, isMacro, e)
	if err != nil {
		return nil, extendError("creating lambda function", err)
	}
	err = e.SetTopLevel(name.s, fn)
	if err != nil {
		return nil, extendError("setting defn result", err)
	}
	return Nil, nil
}

func evErrors(args *ConsCell, e *env) (Sexpr, error) {
	if args == Nil {
		return nil, baseError("no error spec given")
	}
	sigExpr, ok := args.car.(*ConsCell)
	if !ok {
		return nil, baseError("error signature must be a list")
	}
	sigEvaled, err := eval(sigExpr, e)
	if err != nil {
		return nil, extendError("evaluating error signature", err)
	}
	sigList, ok := sigEvaled.(*ConsCell)
	if !ok {
		return nil, baseError("error signature must be a list")
	}
	errorStr := unwrapList(sigList)
	bodyArgs := args.cdr.(*ConsCell)
	for {
		if bodyArgs == Nil {
			return nil, baseErrorf("error not found in %s", args)
		}
		toEval := bodyArgs.car
		_, err := eval(toEval, e)
		if err != nil {
			if strings.Contains(err.Error(), errorStr) {
				return Nil, nil
			}
			return nil, baseErrorf("error '%s' not found in '%s'",
				errorStr, err.Error())
		}
		bodyArgs = bodyArgs.cdr.(*ConsCell)
	}
}

// Both eval, apply and macroexpansion use this to bind lambda arguments in the
// supplied environment:
func setLambdaArgsInEnv(newEnv *env, lambda *lambdaFn, evaledList []Sexpr) error {
	var err error
	if lambda.restArg != noRestArg {
		if len(lambda.args) > len(evaledList) {
			return baseError("not enough arguments for function")
		}
		err = newEnv.Set(lambda.restArg,
			mkListAsConsWithCdr(evaledList[len(lambda.args):],
				Nil))
		if err != nil {
			return err
		}
	} else {
		if len(lambda.args) < len(evaledList) {
			return baseError("too many arguments for function")
		} else if len(lambda.args) > len(evaledList) {
			return baseError("not enough arguments for function")
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
		return nil, baseError("macro call must be a list")
	}
	lambda, ok := fn.(*lambdaFn)
	if !ok {
		panic("macro call not a lambda function")
	}
	asCons, err := consToExprs(c.cdr)
	if err != nil {
		return nil, extendError("converting macro call to list", err)
	}
	if err := setLambdaArgsInEnv(e, lambda, asCons); err != nil {
		return nil, extendError("setting macro call arguments", err)
	}
	ast := lambda.body
	var ret Sexpr = Nil
	for {
		if ast == Nil {
			return ret, nil
		}
		toEval := ast.car
		ret, err = eval(toEval, e)
		if err != nil {
			return nil, extendError("evaluating macro expansion", err)
		}
		ast, ok = ast.cdr.(*ConsCell)
		if !ok {
			return nil, baseError("macro body must be a list")
		}
	}
}

func macroexpand(expr Sexpr, e *env) (Sexpr, error) {
	var ret Sexpr = expr
	var err error
	for {
		ret, err = macroexpand1(ret, e)
		if err != nil {
			return nil, extendError("macroexpanson", err)
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

// Adapted from
// https://github.com/kanaka/mal/blob/master/impls/go/src/step7_quote/step7_quote.go#L36,
// but done recursively:
func splicingUnquote(l *ConsCell) (*ConsCell, error) {
	if l == Nil {
		return Nil, nil
	}
	cdr, ok := l.cdr.(*ConsCell)
	if !ok {
		return l, nil
	}
	nxt, err := splicingUnquote(cdr)
	if err != nil {
		return nil, extendError("splicing unquote", err)
	}
	elt := l.car
	switch t := elt.(type) {
	case *ConsCell:
		if listStartsWith(t, "splicing-unquote") {
			return Cons(Atom{"concat2"}, Cons(t.cdr.(*ConsCell).car, Cons(nxt, Nil))), nil
		}
	default:
	}
	return Cons(Atom{"cons"}, Cons(syntaxQuote(elt), Cons(nxt, Nil))), nil
}

func syntaxQuote(arg Sexpr) Sexpr {
	switch t := arg.(type) {
	case Number, Atom:
		return Cons(Atom{"quote"}, Cons(arg, Nil))
	case *ConsCell:
		if listStartsWith(t, "unquote") {
			return t.cdr.(*ConsCell).car
		}
		ret, err := splicingUnquote(t)
		if err != nil {
			return nil
		}
		return ret
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
			return nil, extendError("eval macroexpansion", err)
		}
	}
	switch t := expr.(type) {
	case Atom:
		if isCxr(t) {
			return extractCxrLambda(t, e)
		}
		return evAtom(t, e)
	case Number:
		return expr, nil
	case *ConsCell:
		if t == Nil {
			return Nil, nil
		}
		cdrCons, ok := t.cdr.(*ConsCell)
		if !ok {
			return nil, baseError("malformed list for eval")
		}
		// special forms:
		if carAtom, ok := t.car.(Atom); ok {
			switch {
			case carAtom.s == "quote":
				if cdrCons == Nil {
					return nil, baseError("quote needs an argument")
				}
				return cdrCons.car, nil
			case carAtom.s == "syntax-quote":
				if cdrCons == Nil {
					return nil, baseError("syntax-quote needs an argument")
				}
				expr = syntaxQuote(cdrCons.car)
				goto top
			case carAtom.s == "cond":
				pairList := cdrCons
				if pairList == Nil {
					return Nil, nil
				}
				for {
					if pairList == Nil {
						return Nil, nil
					}
					pair, ok := pairList.car.(*ConsCell)
					if !ok || pair == Nil {
						return nil, baseError("cond requires a list of pairs")
					}
					ev, err := eval(pair.car, e)
					if err != nil {
						return nil, extendError("evaluating cond condition", err)
					}
					if ev == Nil {
						pairList = pairList.cdr.(*ConsCell)
						continue
					}
					// TAIL CALL!!!
					cdrCons, ok := pair.cdr.(*ConsCell)
					if !ok || cdrCons == Nil {
						return nil, baseError("cond requires a list of pairs")
					}
					expr = cdrCons.car
					goto top
				}
			// FIXME: Do as a macro:
			case carAtom.s == "and":
				pairList := cdrCons
				for {
					if pairList == Nil {
						return True, nil
					}
					ev, err := eval(pairList.car, e)
					if err != nil {
						return nil, extendError("and operator", err)
					}
					if ev == Nil {
						return Nil, nil
					}
					pairList, ok = pairList.cdr.(*ConsCell)
					if !ok {
						return nil, baseError("and requires a list of expressions")
					}
				}
			// FIXME: Do as a macro:
			case carAtom.s == "or":
				pairList := cdrCons
				for {
					if pairList == Nil {
						return Nil, nil
					}
					ev, err := eval(pairList.car, e)
					if err != nil {
						return nil, extendError("or operator", err)
					}
					if ev != Nil {
						return ev, nil
					}
					pairList, ok = pairList.cdr.(*ConsCell)
					if !ok {
						return nil, baseError("or requires a list of expressions")
					}
				}
			case carAtom.s == "loop":
				body := cdrCons
				for {
					start := body
				bodyLoop:
					for {
						if start == Nil {
							break bodyLoop
						}
						_, err := eval(start.car, e)
						if err != nil {
							return nil, extendError("loop operator", err)
						}
						start = start.cdr.(*ConsCell)
					}
				}
			case carAtom.s == "swallow":
				start := cdrCons
				for {
					if start == Nil {
						return Nil, nil
					}
					_, err := eval(start.car, e)
					if err != nil {
						return True, nil
					}
					start = start.cdr.(*ConsCell)
				}
			case carAtom.s == "def":
				return evDef(cdrCons, e)
			case carAtom.s == "set!":
				return evSet(cdrCons, e)
			case carAtom.s == "defn":
				return evDefn(cdrCons, false, e)
			case carAtom.s == "defmacro":
				return evDefn(cdrCons, true, e)
			case carAtom.s == "error":
				if cdrCons == Nil {
					return nil, baseError("error requires a non-empty argument list")
				}
				errorExpr, err := eval(cdrCons.car, e)
				if err != nil {
					return nil, extendError("error operator", err)
				}
				return nil, Cons(errorExpr, Nil)
			case carAtom.s == "errors":
				return evErrors(cdrCons, e)
			case carAtom.s == "try":
				var ret Sexpr = Nil
				var err error = nil
				var hadError bool = false
				for {
					if cdrCons == Nil {
						return ret, err
					}
					car, ok := cdrCons.car.(*ConsCell)
					if ok && car != Nil {
						carCarAtom, ok := car.car.(Atom)
						if ok && carCarAtom.s == "catch" {
							if !hadError {
								return ret, err
							}
							cdr, ok := car.cdr.(*ConsCell)
							if !ok {
								return nil, baseError("catch body must be a list with a binding name")
							}
							bindingName := cdr.car
							symStr, ok := bindingName.(Atom)
							if !ok {
								return nil, baseError("catch binding name must be a symbol")
							}
							eInner := mkEnv(e)
							errCons, ok := err.(*ConsCell)
							if !ok {
								return nil, baseError("catch body must be a list with a binding name")
							}
							eInner.Set(symStr.s, errCons)

							cdr, ok = cdr.cdr.(*ConsCell)
							if !ok {
								return nil, baseError("catch body must be a list with a binding name")
							}
							err = nil
							for {
								if cdr == Nil {
									return ret, err
								}
								ret, err = eval(cdr.car, &eInner)
								if err != nil {
									return nil, extendError("catch body", err)
								}
								cdr, ok = cdr.cdr.(*ConsCell)
								if !ok {
									return nil, baseError("catch body must be a list with a binding name")
								}
							}
						}
					}
					var ev Sexpr
					if !hadError {
						ev, err = eval(cdrCons.car, e)
						if err != nil {
							hadError = true
						} else {
							ret = ev
						}
					}
					cdrCons, ok = cdrCons.cdr.(*ConsCell)
					if !ok {
						return nil, baseError("try requires a list of expressions")
					}
				}
			case carAtom.s == "let":
				args := cdrCons
				if args == Nil {
					return nil, baseError("let requires a binding list")
				}
				bindings, ok := args.car.(*ConsCell)
				if !ok {
					return nil, baseError("let bindings must be a list")
				}
				body, ok := args.cdr.(*ConsCell)
				if !ok {
					return nil, baseError("let requires a body")
				}
				newEnv := mkEnv(e)
				for ; bindings != Nil; bindings = bindings.cdr.(*ConsCell) {
					binding, ok := bindings.car.(*ConsCell)
					if !ok || binding == Nil {
						return nil, baseError("a let binding must be a list of binding pairs")
					}
					carAtom, ok := binding.car.(Atom)
					if !ok {
						return nil, baseError("a let binding must be a list of binding pairs")
					}
					asCons, ok := binding.cdr.(*ConsCell)
					if !ok {
						return nil, baseError("a let binding must be a list of binding pairs")
					}
					if asCons == Nil {
						return Nil, nil
					}
					val, err := eval(asCons.car, e)
					if err != nil {
						return nil, extendError("evaluating let bindings", err)
					}
					err = newEnv.Set(carAtom.s, val)
					if err != nil {
						return nil, extendError("setting let bindings", err)
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
						return nil, extendError("evaluating let body", err)
					}
					body = body.cdr.(*ConsCell)
				}
			case carAtom.s == "lambda":
				return mkLambda(cdrCons, false, e)
			}
		}
		// Functions / normal order of evaluation.  Get function to use first:
		evalCar, err := eval(t.car, e)
		if err != nil {
			return nil, extendError("evaluating function object", err)
		}
		// In normal function application, evaluate the arguments executing the
		// function:
		evaledList := []Sexpr{}
		start := cdrCons
		for {
			if start == Nil {
				break
			}
			ee, err := eval(start.car, e)
			if err != nil {
				return nil, extendError("evaluating function arguments", err)
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
				return nil, extendError("lambda env setup", err)
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
					return nil, extendWithList(
						Cons(Atom{"lambda"}, Cons(lambda.body.car, Nil)),
						err)
				}
				lambda.body = lambda.body.cdr.(*ConsCell)
			}
		}
		// Built-in functions:
		builtin, ok := evalCar.(*Builtin)
		if !ok {
			return nil, baseErrorf("%s is not a function", evalCar)
		}
		biResult, err := builtin.Fn(evaledList, e)
		if err != nil {
			return nil, extendError(fmt.Sprintf("builtin function %s",
				builtin.Name), err)
		}
		return biResult, nil
	default:
		return nil, baseErrorf("unknown expression type: %q", t)
	}
}
