package main

import (
	"fmt"
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
		fmt.Printf("Checking that %T-valued %v -> %q\n",
			test.input, test.input, test.output)
		if Num(test.input).String() != test.output {
			T.Errorf("Num(%q).String() != %q", test.input, test.output)
		}
	}
}

func TestSexprStrings(T *testing.T) {
	var tests = []struct {
		input sexpr
		want  string
	}{
		{Nil, "()"},
		{Num(1), "1"},
		{Num("2"), "2"},
		{Cons(Num(1), Cons(Num("2"), Nil)), "(1 2)"},
		{Cons(Num(1), Cons(Num("2"), Cons(Num(3), Nil))), "(1 2 3)"},
		{Cons(
			Cons(
				Num(3),
				Cons(
					Num("1309875618907812098"),
					Nil)),
			Cons(Num(5), Cons(Num("6"), Nil))), "((3 1309875618907812098) 5 6)"},
	}
	for _, test := range tests {
		fmt.Printf("Checking string representation: %v\n", test.input)
		if test.input.String() != test.want {
			T.Errorf("%T-valued %v.String() != %s",
				test.input, test.input, test.want)
		}
	}
}

func TestAtoms(T *testing.T) {
	var tests = []struct {
		input string
	}{
		{"Z"},
		{"ZY"},
		{"+"},
	}
	for _, test := range tests {
		if Atom(test.input).String() != test.input {
			T.Errorf("Atoms(%q).String() != %q", test.input, test.input)
		}
	}
}

func TestFindMatchingParens(T *testing.T) {
	LP := item{itemLeftParen, "("}
	RP := item{itemRightParen, ")"}
	A := item{itemAtom, "A"}
	N := item{itemNumber, "1"}
	I := func(i ...item) []item { return i }
	var tests = []struct {
		input      []item
		begin, end int
	}{
		{I(LP, RP), 0, 1},
		{I(LP, LP, RP, RP), 0, 3},
		{I(LP, RP, LP, RP), 0, 1},
		{I(LP, A, RP), 0, 2},
		{I(LP, N, RP), 0, 2},
		{I(LP, LP, A, RP, RP), 0, 4},
		{I(LP, LP, A, RP, A, RP), 0, 5},
		{I(LP, A, LP, A, RP, RP), 0, 5},
	}
	for _, test := range tests {
		begin, end := balancedParenPoints(test.input)
		if begin != test.begin || end != test.end {
			T.Errorf("balancedParenPoints(%v) = %d, %d, want %d, %d",
				test.input, begin, end, test.begin, test.end)
		}
	}
}

func TestStrToSexprs(T *testing.T) {
	S := func(xs ...sexpr) []sexpr {
		return xs
	}
	L := func(xs ...sexpr) sexpr {
		cons := mkList(xs)
		return cons
	}
	var tests = []struct {
		input string
		want  []sexpr
	}{
		{"a", S(Atom("a"))},
		{"b", S(Atom("b"))},
		{"b c", S(Atom("b"), Atom("c"))},
		{"+", S(Atom("+"))},
		{"foo", S(Atom("foo"))},
		{"1", S(Num(1))},
		{"a 3", S(Atom("a"), Num(3))},
		{"()", S(L())},
		{"(a)", S(L(Atom("a")))},
		{"a ()", S(Atom("a"), Nil)},
		{"(())", S(L(L()))},
		{"((1))", S(L(L(Num(1))))},
		{"(a)", S(L(Atom("a")))},
		{"((a))", S(L(L(Atom("a"))))},
		{"(a b)", S(L(Atom("a"), Atom("b")))},
		{"(a (b))", S(L(Atom("a"), L(Atom("b"))))},
		{"((a) b)", S(L(L(Atom("a")), Atom("b")))},
		{"(1)", S(L(Num(1)))},
		{"(1 2)", S(L(Num(1), Num(2)))},
		{"(1 2 3)", S(L(Num(1), Num(2), Num(3)))},
		{"(1 2 3 4)", S(L(Num(1), Num(2), Num(3), Num(4)))},
		{"((1) (2))", S(L(L(Num(1)), L(Num(2))))},
	}
	for _, test := range tests {
		if !reflect.DeepEqual(lexAndParse(test.input), test.want) {
			T.Errorf("%v != %v", lexAndParse(test.input), test.want)
		}
	}
}
