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
		{"(+ 1 1)", "2"},
		{"(+ 1 1 2 3)", "7"},
		{"(+ 1)", "1"},
		{"(+ -1)", "-1"},
		{"(+ 0)", "0"},
		{"(+ 1 2 3 4 5 6 7 8 9 10)", "55"},
		{"(+ 999999999999999 1)", "1000000000000000"},
		{"(+ 1 999999999999999)", "1000000000000000"},
	}
	for _, test := range tests {
		got, err := lexAndParse(test.in)
		if err != nil {
			t.Errorf("lexAndParse(%q) failed: %v", test.in, err)
		}
		res := eval(got[0], env{"a": Num(1)})
		if res.String() != test.out {
			t.Errorf("eval(%q) = %q, want %q", test.in, res, test.out)
		} else {
			t.Logf("eval(%q) = %q", test.in, res)
		}
	}
}

func TestNum(t *testing.T) {
	var tests = []struct {
		in  interface{}
		out string
	}{
		{"1", "1"},
		{1, "1"},
		// {int8(1), "1"},  // other int types not supported yet
		{"1089710983751098757", "1089710983751098757"},
	}
	for _, test := range tests {
		got := Num(test.in).String()
		if got != test.out {
			t.Errorf("Num(%q) = %q, want %q", test.in, got, test.out)
		} else {
			t.Logf("Num(%q) = %q", test.in, got)
		}
	}
}
