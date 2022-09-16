package lisp

import (
	"reflect"
	"testing"
)

func TestNumberStrings(T *testing.T) {
	var tests = []struct {
		input  interface{}
		output string
	}{
		{1, "1"},
		{"1", "1"},
		{666, "666"},
		{"12948712498129877", "12948712498129877"},
	}
	for _, test := range tests {
		if Num(test.input).String() != test.output {
			T.Errorf("Num(%q).String() != %q", test.input, test.output)
		} else {
			T.Logf("Num(%q).String() == %q", test.input, test.output)
		}
	}
}

func TestSexprStrings(T *testing.T) {
	var tests = []struct {
		input Sexpr
		want  string
	}{
		{Nil, "()"},
		{Num(1), "1"},
		{Num("2"), "2"},
		{Cons(Num(1), Cons(Num("2"), Nil)), "(1 2)"},
		{Cons(Num(1), Cons(Num("2"), Cons(Num(3), Nil))), "(1 2 3)"},
		{Cons(Num(1), Num(2)), "(1 . 2)"},
		{Cons(
			Cons(
				Num(3),
				Cons(
					Num("1309875618907812098"),
					Nil)),
			Cons(Num(5), Cons(Num("6"), Nil))), "((3 1309875618907812098) 5 6)"},
	}
	for _, test := range tests {
		if test.input.String() != test.want {
			T.Errorf("%T-valued %v.String() != %s",
				test.input, test.input, test.want)
		} else {
			T.Logf("%T-valued %v.String() == %s",
				test.input, test.input, test.want)
		}
	}
}

func TestStrToSexprs(T *testing.T) {
	S := func(xs ...Sexpr) []Sexpr {
		return xs
	}
	L := func(xs ...Sexpr) Sexpr {
		cons := mkListAsConsWithCdr(xs, Nil)
		return cons
	}
	A := func(s string) Atom {
		return Atom{s}
	}
	var happyPathTests = []struct {
		input string
		want  []Sexpr
	}{
		{"a", S(A("a"))},
		{"b", S(A("b"))},
		{"b c", S(A("b"), A("c"))},
		{"+", S(A("+"))},
		{"foo", S(A("foo"))},
		{"foo-", S(A("foo-"))},
		{"-foo", S(A("-foo"))},
		// Regression test for Issue #25:
		{"foo?", S(A("foo?"))},
		{"'foo", S(L(A("quote"), A("foo")))},
		{"'123", S(L(A("quote"), Num(123)))},
		{"'(1 2 3)", S(L(A("quote"), L(Num(1), Num(2), Num(3))))},
		{"1", S(Num(1))},
		{"a 3", S(A("a"), Num(3))},
		{"()", S(L())},
		{"(a)", S(L(A("a")))},
		{"a ()", S(A("a"), Nil)},
		{"(())", S(L(L()))},
		{"((1))", S(L(L(Num(1))))},
		{"(a)", S(L(A("a")))},
		{"((a))", S(L(L(A("a"))))},
		{"(a b)", S(L(A("a"), A("b")))},
		{"(a (b))", S(L(A("a"), L(A("b"))))},
		{"((a) b)", S(L(L(A("a")), A("b")))},
		{"(1)", S(L(Num(1)))},
		{"(1 2)", S(L(Num(1), Num(2)))},
		{"(1 2 3)", S(L(Num(1), Num(2), Num(3)))},
		{"(1 2 3 4)", S(L(Num(1), Num(2), Num(3), Num(4)))},
		{"((1) (2))", S(L(L(Num(1)), L(Num(2))))},
		{"((1 2) (3 4))", S(L(L(Num(1), Num(2)), L(Num(3), Num(4))))},
		{"((1 2) (3 4) (5 6))", S(L(L(Num(1), Num(2)), L(Num(3), Num(4)), L(Num(5), Num(6))))},
		{"(((1 2) (3 4)) (5 6))", S(L(L(L(Num(1), Num(2)), L(Num(3), Num(4))), L(Num(5), Num(6))))},
	}
	for _, test := range happyPathTests {
		parsed, err := lexAndParse([]string{test.input})
		if err != nil {
			T.Errorf("lexAndParse(%q) failed: %v", test.input, err)
		}
		if !reflect.DeepEqual(parsed, test.want) {
			T.Errorf("%v != %v", parsed, test.want)
		}
	}
	var sadPathTests = []struct {
		input string
	}{
		{"("},
		{"(a"},
		{"((a b"},
		{")"},
		{"))"},
		{")())"},
	}
	for _, test := range sadPathTests {
		_, err := lexAndParse([]string{test.input})
		if err == nil {
			T.Errorf("lexAndParse(%q) should have failed", test.input)
		} else {
			T.Logf("lexAndParse(%q) failed as desired: %v ... OK", test.input, err)
		}
	}
}
