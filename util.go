package main

import "fmt"

func mkListAsCons(xs []Sexpr) *ConsCell {
	if len(xs) == 0 {
		return Nil
	}
	return Cons(xs[0], mkListAsCons(xs[1:]))
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
