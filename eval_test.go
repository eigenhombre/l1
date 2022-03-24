package main

import (
	"os"
	"strings"
	"testing"
)

func TestEval(t *testing.T) {
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
		{ECases(S("t", "t", OK))},
		{Cases(S("t ;", "t", OK))},
		{ECases(S("()  ;; Nil by any other name, would still smell as sweet...", "()", OK))},
		{ECases(S("(cons t ())", "(t)", OK))},
		{ECases(S("(cons (quote hello) (quote (world)))", "(hello world)", OK))},
		{ECases(S("(quote foo)", "foo", OK))},
		{Cases(S("(quote 3)", "3", OK))},
		{Cases(S("(quote (1 2 3))", "(1 2 3)", OK))},
		{Cases(S("(quote ())", "()", OK))},
		{Cases(S("(quote (((1 2 3))))", "(((1 2 3)))", OK))},
		{ECases(S("(quote (the (ten (laws (of (greenspun))))))", "(the (ten (laws (of (greenspun)))))", OK))},
		{ECases(S("(cdr (quote (is not common lisp)))", "(not common lisp)", OK))},
		{ECases(S("(car (quote (is not common lisp)))", "is", OK))},
		{ECases(S("(len (quote (1 2 3)))", "3", OK))},
		{Cases(S("(len ())", "0", OK))},
		{Cases(S("(len 3)", "", "is not a list"))},
		{Cases(S("+", "<builtin: +>", OK))},
		{Cases(S("1", "1", OK))},
		{Cases(S("-5", "-5", OK))},
		{ECases(S("(+ 1 1)", "2", OK))},
		{ECases(S("(+ 1 2)", "3", OK))},
		{Cases(S("(+ 1 1 2 3)", "7", OK))},
		{Cases(S("(+)", "0", OK))},
		{ECases(S("(* 12349807213490872130987 12349807213490872130987)", "152517738210391179737088822267441718485594169", OK))},
		{Cases(S("(eq 0 0)", "t", OK))},
		{ECases(S("(eq (quote foo) (quote foo))", "t", OK))},
		{ECases(S("(eq (quote foo) (quote bar))", "()", OK))},
		{ECases(S("(eq (quote foo) (quote (foo bar)))", "()", OK))},
		// P.G.'s interpretation of McCarthy says this is (), but
		// it's simpler to have just one equality operator for now,
		// which works for numbers, lists and atoms:
		// {Cases(S("(eq (quote (foo bar)) (quote (foo bar)))", "t", OK))},
		{Cases(S("(eq 2 (+ 1 1))", "t", OK))},
		{Cases(S("(eq 2 (+ 1 1) (- 3 1))", "t", OK))},
		{Cases(S("(atom (quote foo))", "t", OK))},
		{ECases(S("(atom (quote (foo bar)))", "()", OK))},
		{ECases(S("(atom (quote atom))", "t", OK))},
		{Cases(S("(atom atom)", "()", OK))},
		{Cases(S("(+ 1)", "1", OK))},
		{Cases(S("(+ -1)", "-1", OK))},
		{Cases(S("(+ 0)", "0", OK))},
		{Cases(S("(+ 1 2 3 4 5 6 7 8 9 10)", "55", OK))},
		{Cases(S("(+ 999999999999999 1)", "1000000000000000", OK))},
		{Cases(S("(+ 1 999999999999999)", "1000000000000000", OK))},
		{Cases(S("(+ (+ 1))", "1", OK))},
		{Cases(S("(+ (+ 1 2 3) 4 5 6)", "21", OK))},
		{Cases(S("(- 1)", "-1", OK))},
		{Cases(S("(- 1 1)", "0", OK))},
		{Cases(S("31489071430987532109487513094875031984750983147", "31489071430987532109487513094875031984750983147", OK))},
		{Cases(S("(- 12349807213490872130987 12349807213490872130987)", "0", OK))},
		{Cases(S("(- (+ 1 2 3) 4 5 6)", "-9", OK))},
		{Cases(S("(*)", "1", OK))},
		{Cases(S("(* 1 1)", "1", OK))},
		{Cases(S("(* 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20)", "2432902008176640000", OK))},
		{Cases(S("(/ 1 1)", "1", OK))},
		{ECases(S("(cond (() 1) (2 3))", "3", OK))},
		{Cases(S("(", "", "unbalanced parens"))},
		{Cases(S("(1", "", "unbalanced parens"))},
		{Cases(S("((1", "", "unbalanced parens"))},
		{Cases(S("((1)", "", "unbalanced parens"))},
		{Cases(S("((1))(", "", "unbalanced parens"))},
		{Cases(S(")", "", "unexpected right paren"))},
		{Cases(S("a", "", "unknown symbol"))},
		{Cases(S("(/ 4 2)", "2", OK))},
		{Cases(S("(/ 1 2)", "0", OK))},
		{Cases(S("(-)", "", "missing argument"))},
		{Cases(S("(cons)", "", "missing argument"))},
		{Cases(S("(atom)", "", "missing argument"))},
		{Cases(S("(eq)", "", "missing argument"))},
		{Cases(S("(/)", "", "missing argument"))},
		{Cases(S("(/ 1 0)", "", "division by zero"))},
		{Cases(S("(+ 0 (cond))", "", "expected number"))},
		{Cases(S("(+ t 1)", "", "expected number"))},
		{Cases(S("(- t)", "", "expected number"))},
		{Cases(S("(- t 1)", "", "expected number"))},
		{Cases(S("(- 1 t)", "", "expected number"))},
		{Cases(S("(* t)", "", "expected number"))},
		{Cases(S("(/ t)", "", "expected number"))},
		{Cases(S("(/ 1)", "1", OK))},
		{Cases(S("(/ 1 t)", "", "expected number"))},
		{Cases(S("(car)", "", "missing argument"))},
		{Cases(S("(car t)", "", "not a list"))},
		{Cases(S("(cdr)", "", "missing argument"))},
		{Cases(S("(cdr t)", "", "not a list"))},
		{Cases(S("(* 1 1 1 (*) (*) (*))", "1", OK))},
		{Cases(S("(+ 1 1 1 (+) (+) (+))", "3", OK))},
		{ECases(S("(car (quote (1 2 3)))", "1", OK))},
		{ECases(S("(cdr (quote (1 2 3)))", "(2 3)", OK))},
		{ECases(S("(cons 1 (quote (2 3 4)))", "(1 2 3 4)", OK))},
		{Cases(S("(cons (cons 1 (cons 2 ())) (quote (3 4)))", "((1 2) 3 4)", OK))},
		{Cases(S("(quote ((1)))", "((1))", OK))},
		{Cases(S("(quote (()))", "(())", OK))},
		{Cases(S("(quote (() ()))", "(() ())", OK))},
		{Cases(S("(cons () ())", "(())", OK))},
		{Cases(S("(cond)", "()", OK))},
		{Cases(S("(cond (() 3))", "()", OK))},
		{Cases(S("(cond (3 3))", "3", OK))},
		{Cases(S("(cond (() 3) (() 4))", "()", OK))},
		{Cases(S("(cond (t 3) (t 4))", "3", OK))},
		{Cases(S("(cond ((-) t))", "", "missing argument"))},
		{Cases(S("(cond (t (-)))", "", "missing argument"))},
		{Cases(S("(cond (t t) ((-) t))", "t", OK))},
		{Cases(S("(cond (() t) ((-) t))", "", "missing argument"))},
		{Cases(S("(cond (() t) (t (-)))", "", "missing argument"))},
		// `split` function
		{Cases(S("(split)", "", "expects a single argument"))},
		{Cases(S("(split 1)", "(1)", OK))},
		{Cases(S("(split (quote a))", "(a)", OK))},
		{Cases(S("(split (quote (a b c)))", "", "expects an atom or a number"))},
		{ECases(S("(split (quote greenspun))", "(g r e e n s p u n)", OK))},
		{ECases(S("(split (* 12345 67890))", "(8 3 8 1 0 2 0 5 0)", OK))},
		{ECases(S("(len (split (* 99999 99999 99999)))", "15", OK))},
		// `fuse`
		{Cases(S("(fuse)", "", "expects a single argument"))},
		{Cases(S("(fuse ())", "()", OK))},
		{Cases(S("(fuse (quote (a)))", "a", OK))},
		{Cases(S("(fuse (quote (aa)))", "aa", OK))},
		{ECases(S("(fuse (quote (a b)))", "ab", OK))},
		{Cases(S("(fuse (quote (1)))", "1", OK))},
		{Cases(S("(fuse (quote (1 2)))", "12", OK))},
		{ECases(S("(+ 2 (fuse (quote (1 2 3))))", "125", OK))},
		{ECases(S("(fuse (split 1295807125987))", "1295807125987", OK))},
		// `randigits`
		{Cases(S("(randigits)", "", "expects a single argument"))},
		{ECases(S("(len (randigits 10))", "10", OK))},
		// Apply
		{ECases(S("(apply + (quote (1 2 3)))", "6", OK))},
		{Cases(S("(apply)", "", "expects exactly two arguments"))},
		{Cases(S("(apply +)", "", "expects exactly two arguments"))},
		{Cases(S("(apply + ())", "0", OK))},
		{Cases(S("(apply + (quote (1)))", "1", OK))},
		{Cases(S("(apply * (quote (1 2 3)))", "6", OK))},
		{Cases(S("(apply * (split 123))", "6", OK))},
		{ECases(S("(apply * (split 123456789))", "362880", OK))},
		{ECases(S("(apply / (split 1111))", "1", OK))},
		{ECases(S("(apply eq (split (quote ooo)))", "t", OK))},
		{ECases(S("(apply eq (split (quote foo)))", "()", OK))},
		{Cases(S("(apply (lambda (x y) (eq x y)) (split 11))", "t", OK))},
		{Cases(S("(apply (lambda (x y) (eq x y)) (split 123))", "", "wrong number of args"))},
		{ECases(S("(apply (lambda (x y z) (eq x y z)) (split 121))", "()", OK))},
		// Higher order functions:
		{ECases(S("((cond (t +)))", "0", OK))},
		{ECases(S("((car (cons + ())) 1 2 3)", "6", OK))},
		{Cases(
			S("(def a +)", "<builtin: +>", OK),
			S("(a 1 1)", "2", OK))},
		// Whitespace:
		{Cases(S(" t ", "t", OK))},
		{Cases(S("t\n", "t", OK))},
		{Cases(S("\n\nt\n", "t", OK))},
		// Multiple statements:
		{Cases(
			S("1", "1", OK),
			S("2", "2", OK),
			S("3", "3", OK))},
		// Global scope:
		{Cases(
			S("(def a 3)", "3", OK),
			S("a", "3", OK))},
		{Cases(
			S("(def a 3)", "3", OK),
			S("(def a 4)", "4", OK),
			S("a", "4", OK))},
		{ECases(
			S("(def a 6)", "6", OK),
			S("(def b 7)", "7", OK),
			S("(+ a b)", "13", OK))},
		{Cases(
			S("(def l (quote (1 2 3)))", "(1 2 3)", OK),
			S("l", "(1 2 3)", OK))},
		{Cases(
			S("(def a (+ 1 1))", "2", OK),
			S("(def b (+ a a))", "4", OK),
			S("b", "4", OK))},
		// Print:
		{Cases(S("(print 1)", "()", OK))},
		{Cases(S("(print 1 2)", "()", OK))},
		{Cases(S("(print)", "()", OK))},
		// Simple Functions:
		{Cases(S("(lambda ())", "<lambda()>", OK))},
		{Cases(S("(lambda (x))", "<lambda(x)>", OK))},
		{Cases(S("(lambda (a b zz))", "<lambda(a b zz)>", OK))},
		{ECases(S("((lambda ()))", "()", OK))},
		{ECases(S("((lambda (x) (+ 1 x)) 1)", "2", OK))},
		{Cases(S("((lambda () 333))", "333", OK))},
		{Cases(S("((lambda () 1))", "1", OK))},
		// Environment & scope:
		{Cases(
			S("(def x 0)", "0", OK),
			S("(cond ((eq x 0) 0) (t x))", "0", OK),
			S("(def x 1)", "1", OK),
			S("(cond ((eq x 0) 0) (t x))", "1", OK))},
		{Cases(
			S("(def a 1)", "1", OK),
			S("(def b 2)", "2", OK),
			S("((lambda (x) (+ x a b)) 3)", "6", OK))},
		{Cases(
			S("(def f (lambda (x) (+ 1 x)))", "<lambda(x)>", OK),
			S("(f 2)", "3", OK))},
		{Cases(
			S("(def x 0)", "0", OK),
			S("(eq x 0)", "t", OK))},
		{Cases(
			S("(def f (lambda (x) (cond (x 3) (t 4))))", "<lambda(x)>", OK),
			S("(f t)", "3", OK),
			S("(f 1)", "3", OK),
			S("(f (quote (1 2 3)))", "3", OK),
			S("(f ())", "4", OK))},
		{Cases(
			S("(def f (lambda (x) (cond ((eq x 3) 3) (t 4))))", "<lambda(x)>", OK),
			S("(f 3)", "3", OK),
			S("(f (quote (1 2 3)))", "4", OK),
			S("(f ())", "4", OK))},
		{Cases(
			S("(def f (lambda (x) (cond ((eq x 3) 1) (t (+ 1 (f 3))))))", "<lambda(x)>", OK),
			S("(f 3)", "1", OK),
			S("(f 4)", "2", OK))},
		{Cases(
			S("(def f (lambda (x) (+ x (g (- x 1)))))", "<lambda(x)>", OK),
			S("(def g (lambda (x) 0))", "<lambda(x)>", OK),
			S("(f 1)", "1", OK),
		)},
		// This case helped me (JJ) find a subtle math/pointer bug:
		{Cases(
			S(`(def f (lambda (x)
			            (cond ((eq x 0) 0)
						      (t (+ x (f (- x 1)))))))`,
				"<lambda(x)>", OK),
			S("(f 0)", "0", OK),
			S("(f 1)", "1", OK))},
		// Factorial:
		{ECases(
			S("(def fact (lambda (n) (cond ((eq 0 n) 1) (t (* n (fact (- n 1)))))))", "<lambda(n)>", OK),
			S("(fact 50)", "30414093201713378043612608166064768844377641568960512000000000000", OK),
			S("(len (split (fact 1000)))", "2568", OK))},
		// Fibonacci (slow!):
		{ECases(
			S("(def fib (lambda (n) (cond ((eq 0 n) 0) ((eq 1 n) 1) (t (+ (fib (- n 1)) (fib (- n 2)))))))", "<lambda(n)>", OK),
			S("(fib 0)", "0", OK),
			S("(fib 1)", "1", OK),
			S("(fib 7)", "13", OK),
			S("(fib 10)", "55", OK),
			S("(fib 20)", "6765", OK),
		)},
		// Variable shadowing:
		{Cases(
			S("(def a 1)", "1", OK),
			S("(def f (lambda (a) (+ a a)))", "<lambda(a)>", OK),
			S("(f 1)", "2", OK),
			S("(def a 999)", "999", OK),
			S("(f 2)", "4", OK))},
		{ECases(
			S("(def a 1)", "1", OK),
			S("(def f (lambda () (def a 2) a))", "<lambda()>", OK),
			S("(f)", "2", OK),
			S("a", "1", OK),
		)},
		// Lexical closures:
		{Cases(
			S(`(def ffer (lambda ()
			               ((lambda (a)
						      (lambda (x)
							    (+ a x)))
							3)))`,
				"<lambda()>", OK),
			S("((ffer) 4)", "7", OK),
		)},
		{ECases(
			S("(def incrementer (lambda (n) (lambda (x) (+ x n))))", "<lambda(n)>", OK),
			S("(def inc (incrementer 1))", "<lambda(x)>", OK),
			S("(inc 5)", "6", OK),
			S("(def add2 (incrementer 2))", "<lambda(x)>", OK),
			S("(add2 5)", "7", OK),
		)},
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
	examples := []string{"    $ l1"}
	for _, test := range tests {
		globals := mkEnv(nil)
		for _, testCase := range test.evalCases {
			got, err := lexAndParse(testCase.in)
			if isError(err, testCase) {
				continue
			}
			if len(got) != 1 {
				t.Errorf("\n\n%s: got %d results, want 1!!!!!!!!\n\n", testCase.in, len(got))
				continue
			}
			ev, err := got[0].Eval(&globals)
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
	examples = append(examples, "> ^D\n", "$\n")
	outstr := strings.Join(examples, "\n    ")
	bs := []byte(outstr)
	err := os.WriteFile("examples.txt", bs, 0644)
	if err != nil {
		t.Errorf("write file: %v", err)
	}
}
