package main

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		// Canonical examples for README:
		{"t", "t"},
		{"()", "()"},
		{"(quote foo)", "foo"},
		{
			"(quote (the (ten (laws (of (greenspun))))))",
			"(the (ten (laws (of (greenspun)))))",
		},
		{"(cdr (quote (is not common lisp)))", "(not common lisp)"},
		{"(car (quote (is not common lisp)))", "is"},
		{"1", "1"},
		{"-5", "-5"},
		{"(* 12349807213490872130987 12349807213490872130987)",
			"152517738210391179737088822267441718485594169"},
		{"(+)", "0"},
		{"(+ 1 1 2 3)", "7"},
		// {"(-)", "ERROR"},
		{"(+ 1 1)", "2"},
		{"(eq (quote foo) (quote foo))", "t"},
		{"(eq (quote foo) (quote bar))", "()"},
		{"(eq (quote foo) (quote (foo bar)))", "()"},
		{"(atom (quote foo))", "t"},
		{"(atom (quote (foo bar)))", "()"},
		{"(+ 1)", "1"},
		{"(+ -1)", "-1"},
		{"(+ 0)", "0"},
		{"(+ 1 2 3 4 5 6 7 8 9 10)", "55"},
		{"(+ 999999999999999 1)", "1000000000000000"},
		{"(+ 1 999999999999999)", "1000000000000000"},
		{"(+ (+ 1))", "1"},
		{"(+ (+ 1 2 3) 4 5 6)", "21"},
		{"(- 1)", "-1"},
		{"(- 1 1)", "0"},
		{
			"31489071430987532109487513094875031984750983147",
			"31489071430987532109487513094875031984750983147",
		},
		{"(- 12349807213490872130987 12349807213490872130987)", "0"},
		{"(- (+ 1 2 3) 4 5 6)", "-9"},
		{"(*)", "1"},
		{"(* 1 1)", "1"},
		{"(* 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20)", "2432902008176640000"},
		{"(/ 1 1)", "1"},
		// {"(/)", "ERROR"},
		{"(/ 4 2)", "2"},
		{"(/ 1 2)", "0"},
		// `a` is bound in the test environment:
		{"a", "1"},
		{"(* a a a (*) (*) (*))", "1"},
		{"(+ a a a (+) (+) (+))", "3"},
		{"(quote 3)", "3"},
		{"(quote (1 2 3))", "(1 2 3)"},
		{"(quote ())", "()"},
		{"(quote (((1 2 3))))", "(((1 2 3)))"},
		{"(car (quote (1 2 3)))", "1"},
		{"(cdr (quote (1 2 3)))", "(2 3)"},
		{"(eq (quote (foo bar)) (quote (foo bar)))", "()"},
	}
	for i, test := range tests {
		got, err := lexAndParse(test.in)
		if err != nil {
			t.Errorf("lexAndParse(%q) failed: %v", test.in, err)
		}
		res := eval(got[0], env{"a": Num(1)})
		if res.String() != test.out {
			t.Errorf("eval(%q) = %q, want %q", test.in, res, test.out)
		} else {
			if i > 16 {
				t.Logf("eval(%q) = %q", test.in, res)
			} else {
				// Print first few formatted for README
				fmt.Printf("    > %s\n    %s\n", test.in, res)
			}
		}
	}
}
