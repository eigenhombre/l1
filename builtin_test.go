package main

import (
	"reflect"
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
		{"-", "", "((unexpected end of input))"},
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

func TestListFromSemver(t *testing.T) {
	var tests = []struct {
		input string
		want  []Sexpr
	}{
		{"v1", []Sexpr{Num(1)}},
		{"1", []Sexpr{Num(1)}},
		{"v1.", []Sexpr{Num(1)}},
		{"v1.2", []Sexpr{Num(1), Num(2)}},
		{"v11.2", []Sexpr{Num(11), Num(2)}},
		{"v11.22", []Sexpr{Num(11), Num(22)}},
		{"v11.22.", []Sexpr{Num(11), Num(22)}},
		{"v11.22.33", []Sexpr{Num(11), Num(22), Num(33)}},
		{"v0.0.0", []Sexpr{Num(0), Num(0), Num(0)}},
		{"v01.02.03", []Sexpr{Num(1), Num(2), Num(3)}},
		{"v1.2.3-dirty", []Sexpr{Num(1), Num(2), Num(3), Atom{"dirty"}}},
	}
	for _, test := range tests {
		got := semverAsExprs(test.input)
		if len(test.want) != len(got) {
			t.Errorf("listFromSemver(%q) = %q, want %q", test.input, got, test.want)
		} else if !reflect.DeepEqual(test.want, got) {
			t.Errorf("listFromSemver(%q) = %q, want %q", test.input, got, test.want)
		} else {
			t.Logf("listFromSemver(%q) = %q", test.input, got)
		}
	}
}
