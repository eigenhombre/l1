package main

import "testing"

func TestEval(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{"1", "1"},
		{"1089710983751098757", "1089710983751098757"},
		{"t", "t"},
		{"()", "()"},
		{"a", "1"},
		{"(quote 3)", "3"},
		{"(quote foo)", "foo"},
		{"(quote (1 2 3))", "(1 2 3)"},
		{"(quote ())", "()"},
		{"(quote (((1 2 3))))", "(((1 2 3)))"},
	}
	for _, test := range tests {
		res := eval(lexAndParse(test.in)[0], env{"a": Num(1)})
		if res.String() != test.out {
			t.Errorf("eval(%q) = %q, want %q", test.in, res, test.out)
		} else {
			t.Logf("eval(%q) = %q", test.in, res)
		}
	}
}
