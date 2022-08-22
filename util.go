package main

import (
	"unicode"
	"unicode/utf8"
)

func mkListAsConsWithCdr(xs []Sexpr, cdr Sexpr) Sexpr {
	if len(xs) == 0 {
		return cdr
	}
	return Cons(xs[0], mkListAsConsWithCdr(xs[1:], cdr))
}

func consToExprs(argList Sexpr) []Sexpr {
	args := []Sexpr{}
	for argList != Nil {
		cons := argList.(*ConsCell)
		args = append(args, cons.car)
		argList = cons.cdr
	}
	return args
}

func capitalize(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}
