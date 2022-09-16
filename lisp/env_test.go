package lisp

import (
	"reflect"
	"testing"
)

func TestEnv(t *testing.T) {
	assertVal := func(e *Env, sym string, val Sexpr) {
		lookupVal, found := e.Lookup(sym)
		if !found {
			t.Errorf("expected to find %s in %v", sym, e)
		}
		if !reflect.DeepEqual(val, lookupVal) {
			t.Errorf("expected %s to be %v, got %v", sym, val, lookupVal)
		}
	}
	top := mkEnv(nil)
	top.Set("a", Num(1))
	assertVal(&top, "a", Num(1))
	child := mkEnv(&top)
	assertVal(&child, "a", Num(1))
	child.Set("b", Num(2))
	assertVal(&child, "b", Num(2))

	err := child.Set("t", Num(3))
	if err == nil {
		t.Errorf("expected error setting t")
	}
}
