package main

import "testing"

func TestEval(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{"1", "1"},
		{"1089710983751098757", "1089710983751098757"},
		{"()", "()"},
	}
	for _, test := range tests {
		res := eval(lexAndParse(test.in)[0])
		if res.String() != test.out {
			t.Errorf("eval(%q) = %q, want %q", test.in, res, test.out)
		}
	}
}
