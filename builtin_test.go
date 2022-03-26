package main

import (
	"testing"
)

func TestListOfNums(t *testing.T) {
	OK := ""
	var tests = []struct {
		input string
		want  string
		err   string
	}{
		{"1", "(1)", OK},
		{"12", "(1 2)", OK},
		{"123", "(1 2 3)", OK},
		{"-1", "(-1)", OK},
		{"-12345", "(-1 2 3 4 5)", OK},
		{"-", "", "unexpected end of input"},
	}
	for _, test := range tests {
		got, err := listOfNums(test.input)
		switch {
		case len(test.err) == 0 && err == nil:
			if test.want != got.String() {
				t.Errorf("listOfNums(%q) = %q, want %q", test.input, got, test.want)
			} else {
				t.Logf("listOfNums(%q) = %q", test.input, got)
			}
		case len(test.err) > 0 && err != nil:
			if test.err != err.Error() {
				t.Errorf("listOfNums(%q) error = %q, want %q", test.input, err, test.err)
			} else {
				t.Logf("listOfNums(%q) error = %q", test.input, err)
			}
		default:
			t.Errorf("listOfNums(%q) error = %q, want %q", test.input, err, test.err)
		}
	}
}
