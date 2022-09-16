package lisp

import (
	"os"
	"strings"
	"testing"
)

// See also: tests.l1
func TestEval(t *testing.T) {
	globals := InitGlobals()
	err := LexParseEval(RawCore, &globals)
	if err != nil {
		t.Fatal(err)
	}
	type evalCase struct {
		in        string
		out       string
		err       string
		exemplary bool
	}
	S := func(in, out, err string) evalCase {
		return evalCase{in, out, err, false}
	}
	Cases := func(cases ...evalCase) []evalCase {
		return cases
	}
	ECases := func(cases ...evalCase) []evalCase {
		for i := range cases {
			cases[i].exemplary = true
		}
		return cases
	}
	OK := "" // No error reported
	var tests = []struct {
		evalCases []evalCase
	}{
		{Cases(S("t ;", "t", OK))},
		{Cases(S("()  ;; Nil by any other name, would still smell as sweet...", "()", OK))},
		{Cases(S("3 ;; colons (:) are OK too (regression for #25)", "3", OK))},
		{ECases(S("(quote foo)", "foo", OK))},
		{ECases(S("'foo", "foo", OK))},
		{ECases(S("(len (split '水果))", "2", OK))},
		{Cases(S("'123", "123", OK))},
		{Cases(S("'bar", "bar", OK))},
		{Cases(S("(quote (((1 2 3))))", "(((1 2 3)))", OK))},
		{ECases(S("(quote (the (ten (laws (of (greenspun))))))", "(the (ten (laws (of (greenspun)))))", OK))},
		{Cases(S("+", "<builtin: +>", OK))},
		{Cases(S("1", "1", OK))},
		{Cases(S("-5", "-5", OK))},
		// regression tests for Issue #25 (infinite loop):
		{Cases(S("(quote questionable?)", "questionable?", OK))},
		{Cases(S("(quote moneybag$)", "moneybag$", OK))},
		{Cases(S("(quote (a @ b))", "", "unexpected character '@' in input"))},
		{Cases(S("(cond (() 1) (2 3))", "3", OK))},
		{Cases(S("(", "", "unbalanced parens"))},
		{Cases(S("(1", "", "unbalanced parens"))},
		{Cases(S("((1", "", "unbalanced parens"))},
		{Cases(S("((1)", "", "unbalanced parens"))},
		{Cases(S("((1))(", "", "unbalanced parens"))},
		{Cases(S(")", "", "unexpected right paren"))},
		{Cases(S("a", "", "unknown symbol"))},
		{Cases(S("(quote ((1)))", "((1))", OK))},
		{Cases(S("(quote (()))", "(())", OK))},
		{Cases(S("(quote (() ()))", "(() ())", OK))},
		{ECases(S("(cadaaaaaaaaaar '(((((((((((hello world))))))))))))", "world", OK))},
		// Whitespace:
		{Cases(S(" t ", "t", OK))},
		{Cases(S("t\n", "t", OK))},
		{Cases(S("\n\nt\n", "t", OK))},
		// Print:
		{Cases(S("(print 1)", "()", OK))},
		{Cases(S("(print 1 2)", "()", OK))},
		{Cases(S("(print)", "()", OK))},
		// Function representation:
		{Cases(S("(lambda ())", "<lambda()>", OK))},
		{Cases(S("(lambda (x))", "<lambda(x)>", OK))},
		{Cases(S("@", "", "unexpected character '@'"))},
		{ECases(S("((lambda (x . xs) (list x xs)) 1 2 3 4)", "(1 (2 3 4))", OK))},
		{Cases(S("(lambda (x . y))", "<lambda(x . y)>", OK))},
		{Cases(S("(lambda (a b zz))", "<lambda(a b zz)>", OK))},
		// Handling error cases, and `test` blocks:
		{Cases(S("(errors)", "", "no error spec"))},
		{Cases(S("(errors '(no error) t)", "", "error not found"))},
		{Cases(S("(errors t t)", "", "error signature must be a list"))},
		{Cases(S("(errors (+ 1 1) t)", "", "error signature must be a list"))},
		{Cases(S("(errors '(no error) 1 2 3)", "", "error not found"))},
		{Cases(S("(errors '(division by zero) (/ 1 0))", "()", OK))},
		{Cases(S("(errors (cons 'division '(by zero)) (/ 1 0))", "()", OK))},
		{Cases(S("(errors '(one) (/ 1 0))", "", "division by zero"))},
		{Cases(S("(errors '(zero) (/ 1 1))", "", "error not found"))},
		{Cases(S("(test)", "()", OK))},
		{Cases(S("(test 1)", "1", OK))},
		{Cases(S("(test 1 2)", "2", OK))},
		{Cases(S("(test '(divide by zero) (errors '(zero) (/ 1 0)))", "()", OK))},
	}

	isError := func(err error, testCase evalCase) bool {
		if err != nil {
			if testCase.out != "" {
				t.Errorf("%s: expected real output %q, got error %q", testCase.in, testCase.out, err)
			}
			if !strings.Contains(err.Error(), testCase.err) {
				t.Errorf("%s: got error %q, want %q", testCase.in, err, testCase.err)
				// uncomment for more verbose output:
				// } else {
				// t.Logf("%s -> error %q (matches '%q')", testCase.in, err, testCase.err)
			}
			return true
		}
		return false
	}
	examples := []string{"$ l1"}
	for _, test := range tests {
		for _, testCase := range test.evalCases {
			got, err := lexAndParse(strings.Split(testCase.in, "\n"))
			if isError(err, testCase) {
				continue
			}
			if len(got) != 1 {
				t.Errorf("\n\n%s: got %d results ('%q'), want 1!!!!!!!!\n\n",
					testCase.in, len(got), got)
				continue
			}
			ev, err := eval(got[0], &globals)
			if isError(err, testCase) {
				continue
			}
			if testCase.err != "" {
				t.Errorf("\n\n%s: expected error %q, got none!!!!!!!!\n\n", testCase.in, testCase.err)
				continue
			}
			result := ev.String()
			if result != testCase.out {
				t.Errorf("\n\n%s: got %q, want %q!!!!!!!!\n\n", testCase.in, result, testCase.out)
				continue
			}
			// uncomment for more verbose output:
			// t.Logf("%s -> %q", testCase.in, result)
			if testCase.exemplary {
				examples = append(examples, "> "+testCase.in)
				examples = append(examples, result)
			}
		}
	}
	examples = append(examples, "> (help)\n")
	outstr := strings.Join(examples, "\n")
	outstr += ShortDocStr(&globals)
	bs := []byte(outstr)
	bs = append(bs, "\n> ^D\n$\n"...)
	err = os.WriteFile("examples.txt", bs, 0644)
	if err != nil {
		t.Errorf("write file: %v", err)
	}
}
