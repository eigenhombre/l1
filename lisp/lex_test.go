package lisp

import (
	"reflect"
	"testing"

	"github.com/eigenhombre/lexutil"
)

func TestLex(t *testing.T) {
	abbrev := func(typ lexutil.ItemType) func(string, int) Token {
		return func(input string, line int) Token {
			return Token{lexutil.LexItem{Typ: typ, Val: input}, line}
		}
	}
	S := func(input ...string) []string {
		return input
	}
	N := abbrev(itemNumber)
	LP := abbrev(itemLeftParen)
	RP := abbrev(itemRightParen)
	A := abbrev(itemAtom)
	DOT := abbrev(itemDot)
	QUOTE := abbrev(itemForwardQuote)
	BACKQUOTE := abbrev(itemSyntaxQuote)
	UNQUOTE := abbrev(itemUnquote)
	SPLICINGUNQUOTE := abbrev(itemSplicingUnquote)
	COMMENTNEXT := abbrev(itemCommentNext)
	SHEBANG := abbrev(itemShebang)
	Err := abbrev(itemError)
	toks := func(items ...Token) []Token {
		if len(items) == 0 {
			return []Token{}
		}
		return items
	}
	var tests = []struct {
		input  []string
		output []Token
	}{
		{S(""), toks()},
		{S(" "), toks()},
		{S("\t"), toks()},
		{S("\n\t\r   \r\n"), toks()},
		{S(";"), toks()},
		{S(";; Ignored until end of line"), toks()},
		{S("1"), toks(N("1", 1))},
		{S("1 ;; is such a lonely number"), toks(N("1", 1))},
		{S("12"), toks(N("12", 1))},
		{S("123 "), toks(N("123", 1))},
		{S(" 312"), toks(N("312", 1))},
		{S("111 222"), toks(N("111", 1), N("222", 1))},
		{S(" 111", "222 "), toks(N("111", 1), N("222", 2))},
		{S("-1"), toks(N("-1", 1))},
		{S("+0"), toks(N("+0", 1))},
		{S("+3 -5 "), toks(N("+3", 1), N("-5", 1))},
		{S("("), toks(LP("(", 1))},
		{S("( "), toks(LP("(", 1))},
		{S(" ("), toks(LP("(", 1))},
		{S("1 ("), toks(N("1", 1), LP("(", 1))},
		{S("(1"), toks(LP("(", 1), N("1", 1))},
		{S(")"), toks(RP(")", 1))},
		{S("(3)"), toks(LP("(", 1), N("3", 1), RP(")", 1))},
		{S("Z"), toks(A("Z", 1))},
		{S("Z "), toks(A("Z", 1))},
		{S("ZZ"), toks(A("ZZ", 1))},
		{S("(EQUAL ", "(TIMES ", "3 4", ") 12", ")"), toks(
			LP("(", 1),
			A("EQUAL", 1),
			LP("(", 2),
			A("TIMES", 2),
			N("3", 3),
			N("4", 3),
			RP(")", 4),
			N("12", 4),
			RP(")", 5),
		)},
		{S("+"), toks(A("+", 1))},
		{S("(+ +1 -2)"), toks(LP("(", 1), A("+", 1), N("+1", 1), N("-2", 1), RP(")", 1))},
		{S("/"), toks(A("/", 1))},
		{S("(/ 1 2)"), toks(LP("(", 1), A("/", 1), N("1", 1), N("2", 1), RP(")", 1))},
		{S("(QUOTE (LAMBDA (X) (PLUS X X)))"), toks(
			LP("(", 1), A("QUOTE", 1), LP("(", 1), A("LAMBDA", 1), LP("(", 1), A("X", 1), RP(")", 1),
			LP("(", 1), A("PLUS", 1), A("X", 1), A("X", 1), RP(")", 1), RP(")", 1), RP(")", 1))},
		// Error handling:
		{S("(atom1 . atom2)"), toks(
			LP("(", 1),
			A("atom1", 1),
			DOT(".", 1),
			A("atom2", 1),
			RP(")", 1))},
		{S("'quoted-atom"), toks(QUOTE("'", 1), A("quoted-atom", 1))},
		{S("~atom"), toks(UNQUOTE("~", 1), A("atom", 1))},
		{S("~@atom"), toks(SPLICINGUNQUOTE("~@", 1), A("atom", 1))},
		{S("`(a ~b ~@c)"), toks(BACKQUOTE("`", 1), LP("(", 1),
			A("a", 1),
			UNQUOTE("~", 1), A("b", 1),
			SPLICINGUNQUOTE("~@", 1), A("c", 1), RP(")", 1))},
		{S("((a . b))"), toks(LP("(", 1), LP("(", 1), A("a", 1), DOT(".", 1), A("b", 1), RP(")", 1), RP(")", 1))},
		{S("((a) . b)"), toks(LP("(", 1), LP("(", 1), A("a", 1), RP(")", 1), DOT(".", 1), A("b", 1), RP(")", 1))},
		{S("@"), toks(Err("unexpected character '@' in input", 1))},
		{S("(", "1 ", "2 @ ", "3)"), toks(
			LP("(", 1),
			N("1", 2),
			N("2", 3),
			Err("unexpected character '@' in input", 3),
			N("3", 4),
			RP(")", 4))},
		{S("#_1"), toks(COMMENTNEXT("#_", 1), N("1", 1))},
		{S("#_(1 2 3)"), toks(COMMENTNEXT("#_", 1), LP("(", 1), N("1", 1), N("2", 1), N("3", 1), RP(")", 1))},
		{S("#!/bin/bash\n1(+)\n"), toks(SHEBANG("#!/bin/bash", 1),
			N("1", 1), LP("(", 1), A("+", 1), RP(")", 1))},
	}

	for _, test := range tests {
		items := LexItems(test.input)
		if !reflect.DeepEqual(items, test.output) {
			t.Errorf("%q: expected %v, got %v ... ERROR", test.input, test.output, items)
		} else {
			t.Logf("%q -> %v ... OK", test.input, items)
		}
	}
}
