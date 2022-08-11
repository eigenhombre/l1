package main

import (
	"reflect"
	"testing"

	"github.com/eigenhombre/lexutil"
)

func TestLex(t *testing.T) {
	abbrev := func(typ lexutil.ItemType) func(string) lexutil.LexItem {
		return func(input string) lexutil.LexItem {
			return lexutil.LexItem{Typ: typ, Val: input}
		}
	}
	N := abbrev(itemNumber)
	LP := abbrev(itemLeftParen)
	RP := abbrev(itemRightParen)
	A := abbrev(itemAtom)
	DOT := abbrev(itemDot)
	QUOTE := abbrev(itemForwardQuote)
	Err := abbrev(itemError)
	toks := func(items ...lexutil.LexItem) []lexutil.LexItem {
		if len(items) == 0 {
			return []lexutil.LexItem{}
		}
		return items
	}
	var tests = []struct {
		input  string
		output []lexutil.LexItem
	}{
		{"", toks()},
		{" ", toks()},
		{"\t", toks()},
		{"\n\t\r   \r\n", toks()},
		{";", toks()},
		{";; Ignored until end of line", toks()},
		{"1", toks(N("1"))},
		{"1 ;; is such a lonely number", toks(N("1"))},
		{"12", toks(N("12"))},
		{"123 ", toks(N("123"))},
		{" 312", toks(N("312"))},
		{"111 222", toks(N("111"), N("222"))},
		{" 111 \n222 ", toks(N("111"), N("222"))},
		{"-1", toks(N("-1"))},
		{"+0", toks(N("+0"))},
		{"+3 -5 ", toks(N("+3"), N("-5"))},
		{"(", toks(LP("("))},
		{"( ", toks(LP("("))},
		{" (", toks(LP("("))},
		{"1 (", toks(N("1"), LP("("))},
		{"(1", toks(LP("("), N("1"))},
		{")", toks(RP(")"))},
		{"(3)", toks(LP("("), N("3"), RP(")"))},
		{"Z", toks(A("Z"))},
		{"(EQUAL (TIMES 3 4) 12)", toks(
			LP("("),
			A("EQUAL"),
			LP("("),
			A("TIMES"),
			N("3"),
			N("4"),
			RP(")"),
			N("12"),
			RP(")"),
		)},
		{"+", toks(A("+"))},
		{"(+ +1 -2)", toks(LP("("), A("+"), N("+1"), N("-2"), RP(")"))},
		{"/", toks(A("/"))},
		{"(/ 1 2)", toks(LP("("), A("/"), N("1"), N("2"), RP(")"))},
		{"(QUOTE (LAMBDA (X) (PLUS X X)))", toks(
			LP("("), A("QUOTE"), LP("("), A("LAMBDA"), LP("("), A("X"), RP(")"),
			LP("("), A("PLUS"), A("X"), A("X"), RP(")"), RP(")"), RP(")"))},
		// Error handling:
		{"(atom1 . atom2)", toks(
			LP("("),
			A("atom1"),
			DOT("."),
			A("atom2"),
			RP(")"))},
		{"'quoted-atom", toks(QUOTE("'"), A("quoted-atom"))},
		{"&", toks(Err("unexpected character '&' in input"))},
		{"(1 2 & 3)", toks(
			LP("("),
			N("1"),
			N("2"),
			Err("unexpected character '&' in input"),
			N("3"),
			RP(")"))},
	}

	for _, test := range tests {
		items := lexItems(test.input)
		if !reflect.DeepEqual(items, test.output) {
			t.Errorf("!!! %q: expected %v, got %v", test.input, test.output, items)
		} else {
			t.Logf("%q -> %v", test.input, items)
		}
	}
}
