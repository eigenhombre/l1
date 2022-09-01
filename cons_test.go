package main

import (
	"testing"
)

func TestConsAsError(t *testing.T) {
	inner := func() error {
		return Cons(Atom{"anError"}, Nil)
	}
	err := inner()
	if err == nil {
		t.Error("expected error")
	}
	if err.Error() != "(anError)" {
		t.Error("wrong error message")
	}
}

func StringsToList(listElems ...string) *ConsCell {
	xs := make([]Sexpr, len(listElems))
	for i, s := range listElems {
		xs[i] = Atom{s}
	}
	return List(xs...)
}

func List(listElems ...Sexpr) *ConsCell {
	// FIXME: don't type assert here after mkListAsConsWithCdr returns a Cons:
	return mkListAsConsWithCdr(listElems, Nil).(*ConsCell)
}

func extendStacktrace(carList *ConsCell, err error) error {
	ret, ok := err.(*ConsCell)
	if !ok {
		return err
	}
	return Cons(carList, ret)
}

func TestConsStackTrace(t *testing.T) {
	inner1 := func() error {
		return List(StringsToList("innerError"))
	}
	inner2 := func() error {
		return extendStacktrace(List(Atom{"middleError"},
			StringsToList("with", "some", "extra", "info")), inner1())
	}
	inner3 := func() error {
		return extendStacktrace(List(Atom{"outerError"}), inner2())
	}
	err := inner3().(*ConsCell)
	if err == nil {
		t.Error("expected error")
	}
	if err.Error() != "((outerError) (middleError (with some extra info)) (innerError))" {
		t.Error("wrong error message:", err.Error())
	}
	// Ensure we can pick apart the stacktrace
	if !err.car.Equal(List(Atom{"outerError"})) {
		t.Error("incorrect car for error message:", err.Error())
	}
}
