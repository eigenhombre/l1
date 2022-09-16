package lisp

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	OK := ""
	var tests = []struct {
		input string
		want  Sexpr
		error string
	}{
		{"()", Nil, OK},
		{"a", Atom{"a"}, OK},
		{"(1)", Cons(Num(1), Nil), OK},
		{"(a b)", Cons(Atom{"a"}, Cons(Atom{"b"}, Nil)), OK},
		{"(a . b)", Cons(Atom{"a"}, Atom{"b"}), OK},
		{"((a . b))", Cons(Cons(Atom{"a"}, Atom{"b"}), Nil), OK},
		{"(quote a)", Cons(Atom{"quote"}, Cons(Atom{"a"}, Nil)), OK},
		{"'(a b)", Cons(Atom{"quote"}, Cons(Cons(Atom{"a"}, Cons(Atom{"b"}, Nil)), Nil)), OK},
		{"'(a . b)", Cons(Atom{"quote"}, Cons(Cons(Atom{"a"}, Atom{"b"}), Nil)), OK},
		{"'((a b) . c)", Cons(Atom{"quote"}, Cons(Cons(Cons(Atom{"a"}, Cons(Atom{"b"}, Nil)), Atom{"c"}), Nil)), OK},
		{"(a b . c)", Cons(Atom{"a"}, Cons(Atom{"b"}, Atom{"c"})), OK},
		{"(1 2 3 . 4)", Cons(Num(1), Cons(Num(2), Cons(Num(3), Num(4)))), OK},
		{"((1) 2 3 . 4)", Cons(Cons(Num(1), Nil), Cons(Num(2), Cons(Num(3), Num(4)))), OK},
		{"(1 (2 . 3) 4 . 5)", Cons(Num(1), Cons(Cons(Num(2), Num(3)), Cons(Num(4), Num(5)))), OK},
		{"(1 2 . (3 4))", Cons(Num(1), Cons(Num(2), Cons(Num(3), Cons(Num(4), Nil)))), OK},
		{"(1 2 . (3 . 4))", Cons(Num(1), Cons(Num(2), Cons(Num(3), Num(4)))), OK},
		{"'((a) . b)", Cons(Atom{"quote"}, Cons(Cons(Cons(Atom{"a"}, Nil), Atom{"b"}), Nil)), OK},
		{"`a", Cons(Atom{"syntax-quote"}, Cons(Atom{"a"}, Nil)), OK},
		{"`(a b)", Cons(Atom{"syntax-quote"}, Cons(Cons(Atom{"a"}, Cons(Atom{"b"}, Nil)), Nil)), OK},
		{"`(a . b)", Cons(Atom{"syntax-quote"}, Cons(Cons(Atom{"a"}, Atom{"b"}), Nil)), OK},
		{"~b", Cons(Atom{"unquote"}, Cons(Atom{"b"}, Nil)), OK},
		{"`~b", Cons(Atom{"syntax-quote"}, Cons(Cons(Atom{"unquote"}, Cons(Atom{"b"}, Nil)), Nil)), OK},
		{"`(~b)", Cons(Atom{"syntax-quote"}, Cons(Cons(Cons(Atom{"unquote"}, Cons(Atom{"b"}, Nil)), Nil), Nil)), OK},
		{"~@c", Cons(Atom{"splicing-unquote"}, Cons(Atom{"c"}, Nil)), OK},
		{"`(a ~b ~@c)", Cons(Atom{"syntax-quote"},
			Cons(Cons(Atom{"a"},
				Cons(Cons(Atom{"unquote"},
					Cons(Atom{"b"}, Nil)),
					Cons(Cons(Atom{"splicing-unquote"},
						Cons(Atom{"c"}, Nil)), Nil))), Nil)), OK},
		{"#_(a b c)", Cons(Atom{"comment"}, Cons(Cons(Atom{"a"}, Cons(Atom{"b"}, Cons(Atom{"c"}, Nil))), Nil)), OK},
		{"#_1", Cons(Atom{"comment"}, Cons(Num(1), Nil)), OK},
		{"#!/bin/bash\n(1 2)\n", Cons(Num(1), Cons(Num(2), Nil)), OK},
		{"\n\n#!/bin/bash\n(1 2)\n", Cons(Num(1), Cons(Num(2), Nil)), OK},
		// Make sure that shebang must come first...
		{"1\n#!/bin/bash", Nil, "unexpected lexeme"},
		// ... and that it reports line number correctly:
		{"1\n#!/bin/bash", Nil, "on line 2"},
		{")", Nil, "unexpected right paren"},
		// line numbers in parse errors:
		{"1\n2\n3\n)", Nil, "unexpected right paren on line 4"},
	}
	for _, test := range tests {
		got, err := lexAndParse(strings.Split(test.input, "\n"))
		if err != nil {
			if test.error == OK {
				t.Errorf("lexAndParse(%q) failed: %v", test.input, err)
			}
			if !strings.Contains(err.Error(), test.error) {
				t.Errorf("lexAndParse(%q) failed with wrong error: want %q, got %q", test.input, test.error, err)
			}
			continue
		}

		if len(got) != 1 {
			t.Errorf("lexAndParse(%q) returned %d values ('%s'), want 1", test.input, len(got), got)
			continue
		}
		if !reflect.DeepEqual(got[0], test.want) {
			t.Errorf("lexAndParse(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}
