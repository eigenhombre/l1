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

func TestStackTrace(t *testing.T) {
	inner1 := func() error {
		return baseError("innerError")
	}
	inner2 := func() error {
		return extendStacktrace(list(Atom{"middleError"},
			stringsToList("with", "some", "extra", "info")), inner1())
	}
	inner3 := func() error {
		return extendStacktrace(list(Atom{"outerError"}), inner2())
	}
	err := inner3().(*ConsCell)
	if err == nil {
		t.Error("expected error")
	}
	if err.Error() != "((outerError) (middleError (with some extra info)) (innerError))" {
		t.Error("wrong error message:", err.Error())
	}
	// Ensure we can pick apart the stacktrace
	if !err.car.Equal(list(Atom{"outerError"})) {
		t.Error("incorrect car for error message:", err.Error())
	}
}
