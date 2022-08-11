package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		input string
		want  Sexpr
	}{
		{"()", Nil},
		{"a", Atom{"a"}},
		{"(1)", Cons(Num(1), Nil)},
		{"(a b)", Cons(Atom{"a"}, Cons(Atom{"b"}, Nil))},
		//{"(a . b)", Cons(Atom{"a"}, Atom{"b"})},
	}
	for _, test := range tests {
		got, err := lexAndParse(test.input)
		if err != nil {
			t.Errorf("lexAndParse(%q) failed: %v", test.input, err)
		}
		if !reflect.DeepEqual(got[0], test.want) {
			t.Errorf("lexAndParse(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}
