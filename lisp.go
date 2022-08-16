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
		panic(err)
	}
	err = e.Set(name, val)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func evDefn(args *ConsCell, e *env) (Sexpr, error) {
	if args == Nil {
		return nil, fmt.Errorf("defn requires a function name")
	}
	name, ok := args.car.(Atom)
	if !ok {
		return nil, fmt.Errorf("defn name must be an atom")
	}
	args = args.cdr.(*ConsCell)
	if args == Nil {
		return nil, fmt.Errorf("defn requires an argument list")
	}
	fn, err := mkLambda(args, e)
	if err != nil {
		return nil, err
	}
	err = e.Set(name.s, fn)
	if err != nil {
		return nil, err
	}
	return Nil, nil
}

func eval(expr Sexpr, e *env) (Sexpr, error) {
top:
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
			case carAtom.s == "def":
				return evDef(t.cdr.(*ConsCell), e)
			case carAtom.s == "defn":
				return evDefn(t.cdr.(*ConsCell), e)
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
				return mkLambda(t.cdr.(*ConsCell), e)
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
			newEnv := mkEnv(lambda.env)
			if lambda.restArg != noRestArg {
				if len(lambda.args) > len(evaledList) {
					return nil, fmt.Errorf("not enough arguments for function")
				}
				err = newEnv.Set(lambda.restArg,
					mkListAsConsWithCdr(evaledList[len(lambda.args):],
						Nil))
				if err != nil {
					return nil, err
				}
			} else {
				if len(lambda.args) < len(evaledList) {
					return nil, fmt.Errorf("too many arguments for function")
				} else if len(lambda.args) > len(evaledList) {
					return nil, fmt.Errorf("not enough arguments for function")
				}
			}
			for i, arg := range lambda.args {
				err := newEnv.Set(arg, evaledList[i])
				if err != nil {
					return nil, err
				}
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
