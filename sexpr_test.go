package main

import (
	"fmt"
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
