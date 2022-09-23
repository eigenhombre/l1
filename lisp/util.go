package lisp

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// FIXME: this should return a cons!
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
			return nil, baseErrorf("expected list, got %q", argList)
		}
		args = append(args, cons.car)
		argList = cons.cdr
	}
	return args, nil
}

func consLength(l *ConsCell) (int, error) {
	ret := 0
	x := l
	var ok bool
	for x != Nil {
		x, ok = x.cdr.(*ConsCell)
		if !ok {
			return 0, baseErrorf("consLength: expected list, got %q", l)
		}
		ret += 1
	}
	return ret, nil
}

func convertStringToDoc(s string) *ConsCell {
	return list(stringsToList(strings.Split(s, " ")...))
}

func stringsToList(listElems ...string) *ConsCell {
	xs := make([]Sexpr, len(listElems))
	for i, s := range listElems {
		xs[i] = Atom{s}
	}
	return list(xs...)
}

func list(listElems ...Sexpr) *ConsCell {
	// FIXME: don't type assert here after mkListAsConsWithCdr returns a Cons:
	return mkListAsConsWithCdr(listElems, Nil).(*ConsCell)
}

func capitalize(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

func unwrapList(arg *ConsCell) string {
	s := arg.String()
	return s[1 : len(s)-1]
}
