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
		{"(a . b)", Cons(Atom{"a"}, Atom{"b"})},
		{"((a . b))", Cons(Cons(Atom{"a"}, Atom{"b"}), Nil)},
		{"(quote a)", Cons(Atom{"quote"}, Cons(Atom{"a"}, Nil))},
		{"'(a b)", Cons(Atom{"quote"}, Cons(Cons(Atom{"a"}, Cons(Atom{"b"}, Nil)), Nil))},
		{"'(a . b)", Cons(Atom{"quote"}, Cons(Cons(Atom{"a"}, Atom{"b"}), Nil))},
		{"'((a b) . c)", Cons(Atom{"quote"}, Cons(Cons(Cons(Atom{"a"}, Cons(Atom{"b"}, Nil)), Atom{"c"}), Nil))},
		{"(a b . c)", Cons(Atom{"a"}, Cons(Atom{"b"}, Atom{"c"}))},
		{"(1 2 3 . 4)", Cons(Num(1), Cons(Num(2), Cons(Num(3), Num(4))))},
		{"((1) 2 3 . 4)", Cons(Cons(Num(1), Nil), Cons(Num(2), Cons(Num(3), Num(4))))},
		{"(1 (2 . 3) 4 . 5)", Cons(Num(1), Cons(Cons(Num(2), Num(3)), Cons(Num(4), Num(5))))},
		{"(1 2 . (3 4))", Cons(Num(1), Cons(Num(2), Cons(Num(3), Cons(Num(4), Nil))))},
		{"(1 2 . (3 . 4))", Cons(Num(1), Cons(Num(2), Cons(Num(3), Num(4))))},
		{"'((a) . b)", Cons(Atom{"quote"}, Cons(Cons(Cons(Atom{"a"}, Nil), Atom{"b"}), Nil))},
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
