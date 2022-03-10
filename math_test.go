package main

import (
	"testing"
)

func TestNumConstruction(t *testing.T) {
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
			t.Errorf("Num(%q).String() != %q", test.input, test.output)
		} else {
			t.Logf("Num(%q).String() == %q", test.input, test.output)
		}
	}
}

func TestNumMath(t *testing.T) {
	var tests = []struct {
		a    Number
		op   string
		b    Number
		want Number
	}{
		{Num(1), "+", Num(2), Num(3)},
		{Num("1359871035987103978"), "+", Num(666), Num("1359871035987104644")},
		{Num(1), "-", Num(2), Num(-1)},
		{Num(9999), "*", Num(666), Num(666 * 9999)},
		{Num(2), "/", Num(2), Num(1)},
	}
	for _, test := range tests {
		f := map[string](func(Number, Number) Number){
			"+": func(a, b Number) Number { return a.Add(b) },
			"-": func(a, b Number) Number { return a.Sub(b) },
			"*": func(a, b Number) Number { return a.Mul(b) },
			"/": func(a, b Number) Number { return a.Div(b) },
		}
		if got := f[test.op](test.a, test.b); !got.Equals(test.want) {
			t.Errorf("%v %s %v = %v, want %v", test.a, test.op, test.b, got, test.want)
		} else {
			t.Logf("%v %s %v = %v", test.a, test.op, test.b, got)
		}
	}
}
