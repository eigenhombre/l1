package main

import "fmt"

func mkListAsConsWithCdr(xs []Sexpr, cdr Sexpr) Sexpr {
	if len(xs) == 0 {
		return cdr
	}
	return Cons(xs[0], mkListAsConsWithCdr(xs[1:], cdr))
}

func consToExprs(argList Sexpr) ([]Sexpr, error) {
	args := []Sexpr{}
	for argList != Nil {
		cons, ok := argList.(*ConsCell)
		if !ok {
			return nil, fmt.Errorf("'%s' is not a list", argList)
		}
		args = append(args, cons.car)
		argList = cons.cdr
	}
	return args, nil
}
