package lisp

import (
	"fmt"
	"sort"
	"strings"
)

const (
	special  = "special form"
	macro    = "macro"
	native   = "native function"
	function = "function"
)

type formRec struct {
	name      string
	isSpecial bool
	isMacro   bool
	isNative  bool
	farity    int
	ismulti   bool
	doc       string
	ftype     string
	args      *ConsCell
	examples  string
}

func a(s string) Sexpr { return Atom{s} }

// When you add a special form to eval, you should add it here as well:
var specialForms = []formRec{
	{
		name:      "and",
		farity:    0,
		isSpecial: true,
		ismulti:   true,
		doc:       "Boolean and",
		ftype:     special,
		args:      Cons(Nil, a("xs")),
		examples: `(and)
;;=>
true
> (and t t)
;;=>
true
> (and t t ())
;;=>
()
> (and () (/ 1 0))
;;=>
()
`,
	},
	{
		name:      "cond",
		farity:    0,
		isSpecial: true,
		ismulti:   true,
		doc:       "Fundamental branching construct",
		ftype:     special,
		args:      Cons(Nil, a("pairs")),
		examples: `> (cond)
;;=> ()
> (cond (t 1) (t 2) (t 3))
;;=> 1
> (cond (() 1) (t 2))
;;=> 2
`,
	},
	{
		name:      "def",
		farity:    2,
		isSpecial: true,
		ismulti:   false,
		doc:       "Set a value",
		ftype:     special,
		args:      list(a("name"), a("value")),
		examples: `> (def a 1)
;;=>
1
> a
;;=>
1
`,
	},
	{
		name:      "defn",
		farity:    2,
		isSpecial: true,
		ismulti:   true,
		doc:       "Create and name a function",
		ftype:     special,
		args:      Cons(a("name"), Cons(a("args"), a("body"))),
		examples: `> (defn add (x y) (+ x y))
;;=>
()
> (add 1 2)
;;=>
3
> (defn add (x y)
    (doc (add two numbers)
         (examples
           (add 1 2)))
    (+ x y))
;;=>
()
> (doc add)
;;=>
((add two numbers) (examples (add 1 2)))
`,
	},
	{
		name:      "defmacro",
		farity:    2,
		isSpecial: true,
		ismulti:   true,
		doc:       "Create and name a macro",
		ftype:     special,
		args:      Cons(a("name"), Cons(a("args"), a("body"))),
		examples: `> (defmacro ignore-car (l)
    (doc (ignore first element of list,
                 treat rest as normal expression)
         (examples
           (ignore-car (adorable + 1 2 3))
           (ignore-car (deplorable - 4 4))))
    (cdr l))
;;=>
()
> (ignore-car (hilarious * 2 3 4))
;;=>
24
	`,
	},
	{
		name:      "error",
		farity:    1,
		isSpecial: true,
		ismulti:   false,
		doc:       "Raise an error",
		ftype:     special,
		args:      list(a("l")),
		examples: `> (defn ensure-list (x)
    (when-not (list? x)
      (error '(ensure-list argument not a list!))))
;;=>
()
> (ensure-list 3)
;;=>
ERROR in '(ensure-list 3)':
(ensure-list argument not a list!)
`,
	},
	{
		name:      "errors",
		farity:    1,
		isSpecial: true,
		ismulti:   true,
		doc:       "Error checking (for tests)",
		args:      Cons(a("expected"), a("body")),
		ftype:     special,
		examples: `> (errors '(is not a function)
    (1))
;;=>
()
> (errors '(is not a function)
    (+))
;;=>
ERROR in '(errors (quote (is not a function)) (+))':
error not found in ((quote (is not a function)) (+))
`,
	},
	{
		name:      "lambda",
		farity:    1,
		isSpecial: true,
		ismulti:   true,
		doc:       "Create a function",
		ftype:     special,
		args:      Cons(a("args"), a("more")),
		examples: `> ((lambda () t))
;;=>
t
> ((lambda (x) (+ 5 x)) 5)
;;=>
10
> ((lambda my-length (x)
     (if-not x
       0
       (+ 1 (my-length (cdr x)))))
    (range 20))
;;=>
20
`,
	},
	{
		name:      "let",
		farity:    1,
		isSpecial: true,
		ismulti:   true,
		doc:       "Create a local scope with bindings",
		ftype:     special,
		args:      Cons(a("binding-pairs"), a("body")),

		examples: `> (let ((a 1)
        (b 2))
    (+ a b))
;;=>
3
`,
	},
	{
		name:      "loop",
		farity:    1,
		isSpecial: true,
		ismulti:   true,
		doc:       "Loop forever",
		ftype:     special,
		args:      Nil,
		examples: `> (loop
    (printl '(Help me, I am looping forever!))
    (sleep 1000))
;; Prints =>
Help me, I am looping forever!
Help me, I am looping forever!
Help me, I am looping forever!
...
`,
	},
	{
		name:      "or",
		farity:    0,
		isSpecial: true,
		ismulti:   true,
		doc:       "Boolean or",
		ftype:     special,
		args:      Cons(Nil, a("xs")),
		examples: `> (or)
;; => false
> (or t t)
;; => true
> (or t t ())
;; => t`,
	},
	{
		name:      "quote",
		farity:    1,
		isSpecial: true,
		ismulti:   false,
		doc:       "Quote an expression",
		ftype:     special,
		args:      list(a("x")),
		examples: `> (quote foo)
foo
> (quote (1 2 3))
(1 2 3)
> '(1 2 3)
(1 2 3)
`,
	},
	{
		name:      "set!",
		farity:    2,
		isSpecial: true,
		ismulti:   false,
		doc:       "Update a value in an existing binding",
		ftype:     special,
		args:      list(a("name"), a("value")),
		examples: `> (def a 1)
;;=>
1
> a
;;=>
1
> (set! a 2)
;;=>
2
> a
;;=>
2
`,
	},
	{
		name:      "swallow",
		farity:    0,
		isSpecial: true,
		ismulti:   true,
		doc:       "Swallow errors thrown in body, return t if any occur",
		ftype:     special,
		args:      Cons(Nil, a("body")),
		examples: `> (swallow
	(error '(boom)))
;;=>
t
> (swallow 1 2 3)
;;=>
()
`,
	},
	{
		name:      "syntax-quote",
		farity:    1,
		isSpecial: true,
		ismulti:   false,
		doc:       "Syntax-quote an expression",
		ftype:     special,
		args:      list(a("x")),
		examples: `> (syntax-quote foo)
foo
> (syntax-quote (1 2 3 4))
(1 2 3 4)
> (syntax-quote (1 (unquote (+ 1 1)) (splicing-unquote (list 3 4))))
(1 2 3 4)
` + "> `(1 ~(+ 1 1) ~@(list 3 4))" + `
(1 2 3 4)
`,
	},
	{
		name:      "try",
		farity:    0,
		isSpecial: true,
		ismulti:   true,
		doc:       "Try to evaluate body, catch errors and handle them",
		ftype:     special,
		args:      Cons(Nil, a("body")),
		examples: `> (try (error '(boom)))
;;=>
ERROR:
((boom))
> (try
    (error '(boom))
    (catch e
      (printl e)))
;;=>
(boom)
> (try (/ 1 0) (catch e (len e)))
2
>
`,
	},
}

const columnsFormat = "%14s %2s %5s  %s"

func formatFunctionInfo(name, shortDesc string,
	arity int,
	isMultiArity, isSpecial, isMacro, isNativeFn bool) string {

	isMultiArityStr := " "
	if isMultiArity {
		isMultiArityStr = "+"
	}
	formType := "F"
	if isSpecial {
		formType = "S"
	} else if isMacro {
		formType = "M"
	} else if isNativeFn {
		formType = "N"
	}
	argstr := fmt.Sprintf("%d%s", arity, isMultiArityStr)
	return fmt.Sprintf(columnsFormat,
		name,
		formType,
		argstr,
		capitalize(shortDesc))
}

func functionDescriptionFromDoc(l lambdaFn) string {
	if l.doc == Nil {
		return "UNDOCUMENTED"
	}
	carDoc := l.doc.car.String()
	shortDoc := carDoc[1 : len(carDoc)-1]
	return shortDoc
}

func functionExamplesFromDoc(l lambdaFn) *ConsCell {
	doc := l.doc
	if doc == Nil {
		return Nil
	}
	for {
		if doc == Nil {
			return Nil
		}
		docCons, ok := doc.car.(*ConsCell)
		if !ok || docCons == Nil {
			return Nil
		}
		if docCons.car.Equal(Atom{"examples"}) {
			return doc.car.(*ConsCell).cdr.(*ConsCell)
		}
		doc = doc.cdr.(*ConsCell)
	}
}

func examplesToString(examples *ConsCell, e *Env) string {
	ret := ""
	for {
		if examples == Nil {
			break
		}
		example := examples.car
		if example == Nil {
			break
		}
		output, err := eval(example, e)
		if err != nil {
			ret += fmt.Sprintf("> %s\n;;=>\nERROR: %s\n", example, err)
		} else {
			ret += fmt.Sprintf("> %s\n;;=>\n%s\n", example, output)
		}
		var ok bool
		examples, ok = examples.cdr.(*ConsCell)
		if !ok {
			ret += "ERROR: examples must be lists!"
		}

	}
	return ret
}

func availableForms(e *Env) ([]formRec, error) {
	// Start with special forms...
	out := specialForms

	// Add builtins...:
	for _, builtin := range builtins {
		out = append(out, formRec{
			name:     builtin.Name,
			farity:   builtin.FixedArity,
			ismulti:  builtin.NAry,
			isNative: true,
			doc:      builtin.Docstring,
			ftype:    native,
			args:     builtin.Args,
			examples: examplesToString(builtin.Examples, e),
		})
	}
	// Add user-defined / internal l1 functions...:
	for _, lambdaName := range EnvKeys(e) {
		expr, _ := e.Lookup(lambdaName)
		l, ok := expr.(*lambdaFn)
		if !ok {
			continue
		}
		ftype := function
		if l.isMacro {
			ftype = macro
		}
		cl, err := consLength(l.args)
		if err != nil {
			return nil, extendError("availableForms", err)
		}
		args := l.args
		if l.restArg != "" {
			args = combineArgs(l.args, Atom{l.restArg})
		}
		if ok && l.doc != Nil {
			examples := examplesToString(functionExamplesFromDoc(*l), e)
			out = append(out, formRec{
				name:     lambdaName,
				farity:   cl,
				isMacro:  l.isMacro,
				ismulti:  l.restArg != "",
				doc:      functionDescriptionFromDoc(*l),
				ftype:    ftype,
				args:     args,
				examples: examples,
			})
		}

	}
	// Order by name:
	sort.Slice(out, func(i, j int) bool {
		return out[i].name < out[j].name
	})
	return out, nil
}

func combineArgs(args *ConsCell, cdr Sexpr) *ConsCell {
	if cdr == Nil {
		return args
	}
	if args == Nil {
		return Cons(Nil, cdr)
	}
	if args.cdr == Nil {
		return Cons(args.car, cdr)
	}
	return Cons(args.car, combineArgs(args.cdr.(*ConsCell), cdr))
}

func codeQuote(s string) string {
	return fmt.Sprintf("`%s`", s)
}

func escapeSpecialChars(s string) string {
	s = strings.ReplaceAll(s, "?", "-QMARK")
	s = strings.ReplaceAll(s, "!", "-BANG")
	s = strings.ReplaceAll(s, "*", "-STAR")
	return s
}

// LongDocStr returns a long, Markdown docstr for a function, macro or
// special form.
func LongDocStr(e *Env) (string, error) {
	sortedForms, err := availableForms(e)
	if err != nil {
		return "", extendError("long-form doc", err)
	}
	summary := fmt.Sprintf("# API Index\n%d forms available:", len(sortedForms))
	for _, form := range sortedForms {
		nameStr := fmt.Sprintf("`%s`", form.name)
		if form.ftype == macro {
			nameStr = fmt.Sprintf("*`%s`*", form.name)
		} else if form.ftype == special {
			nameStr = fmt.Sprintf("**`%s`**", form.name)
		}
		summary += fmt.Sprintf("\n[%s](#%s)", nameStr, escapeSpecialChars(form.name))
	}
	summary += "\n# Operators\n"
	outStrs := []string{summary}
	for _, doc := range sortedForms {
		isMulti := ""
		if doc.ismulti {
			isMulti = "+"
		}
		examples := ""
		if doc.examples != "" {
			examples = fmt.Sprintf("\n### Examples\n\n```\n%s\n```\n", doc.examples)
		}
		outStrs = append(outStrs, fmt.Sprintf(`
<a id="%s"></a>
## %s

%s

Type: %s

Arity: %d%s

Args: %s

%s

[<sub><sup>Back to index</sup></sub>](#api-index)
-----------------------------------------------------
`,
			escapeSpecialChars(doc.name),
			codeQuote(doc.name),
			capitalize(doc.doc),
			doc.ftype,
			doc.farity,
			isMulti,
			fmt.Sprintf("`%s`", doc.args),
			examples))
	}
	return strings.Join(outStrs, "\n"), nil
}

// ShortDocStr returns an abbreviated explanation of all functions,
// macros and special forms.
func ShortDocStr(e *Env) (string, error) {
	outStrs := []string{}
	outStrs = append(outStrs,
		"l1 - a Lisp interpreter.\n",
		fmt.Sprintf(columnsFormat, "", "Type", "", ""),
		fmt.Sprintf(columnsFormat, "", "---", "", ""),
		"                S - special form",
		"                M - macro",
		"                N - native (Go) function",
		"                F - Lisp function\n",
		fmt.Sprintf(columnsFormat, "Name", "Type", "Arity", "Description"),
		fmt.Sprintf(columnsFormat, "----", "---", "----", "-----------"),
	)
	af, err := availableForms(e)
	if err != nil {
		return "", extendError("short-form doc", err)
	}
	for _, doc := range af {
		outStrs = append(outStrs, formatFunctionInfo(doc.name,
			doc.doc,
			doc.farity,
			doc.ismulti,
			doc.isSpecial,
			doc.isMacro,
			doc.isNative))
	}
	return strings.Join(outStrs, "\n"), nil
}

// The (forms) function returns a list of builtins, macros and special
// forms.  Each element of the list is a list of the form:
// (name type arity hasRest)
// FIXME: add doc.
func formsAsSexprList(e *Env) ([]Sexpr, error) {
	out := []Sexpr{}
	af, err := availableForms(e)
	if err != nil {
		return nil, extendError("sexpr forms", err)
	}
	for _, form := range af {
		var multi Sexpr = Nil
		if form.ismulti {
			multi = True
		}
		out = append(out, list(Atom{form.name},
			Atom{strings.Replace(form.ftype, " ", "-", -1)},
			Num(form.farity),
			multi))
	}
	return out, nil
}
