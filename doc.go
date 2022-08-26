package main

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
	farity    int
	ismulti   bool
	doc       string
	columnDoc string
	ftype     string
	argsStr   string
	examples  string
}

// When you add a special form to eval, you should add it here as well:
var specialForms = []formRec{
	{
		name:    "and",
		farity:  0,
		ismulti: true,
		doc:     "Boolean and",
		ftype:   special,
		argsStr: "(() . xs)",
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
		name:    "cond",
		farity:  0,
		ismulti: true,
		doc:     "Fundamental branching construct",
		ftype:   special,
		argsStr: "(() . pairs)",
		examples: `> (cond)
;;=> ()
> (cond (t 1) (t 2) (t 3))
;;=> 1
> (cond (() 1) (t 2))
;;=> 2
`,
	},
	{
		name:    "def",
		farity:  2,
		ismulti: false,
		doc:     "Set a value",
		ftype:   special,
		argsStr: "(name value)",
		examples: `> (def a 1)
;;=>
1
> a
;;=>
1
`,
	},
	{
		name:    "defn",
		farity:  2,
		ismulti: true,
		doc:     "Create and name a function",
		ftype:   special,
		argsStr: "(name args . rest)",
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
		name:    "defmacro",
		farity:  2,
		ismulti: true,
		doc:     "Create and name a macro",
		ftype:   special,
		argsStr: "(name args . body)",
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
		name:    "error",
		farity:  1,
		ismulti: false,
		doc:     "Raise an error",
		ftype:   special,
		argsStr: "(msg-list)",
	},
	{
		name:    "errors",
		farity:  1,
		ismulti: true,
		doc:     "Error checking (for tests)",
		ftype:   special,
		argsStr: "(message-pattern-list . exprs)",
	},
	{
		name:    "lambda",
		farity:  1,
		ismulti: true,
		doc:     "Create a function",
		ftype:   special,
		argsStr: "(args . body) or (name args . body)",
	},
	{
		name:    "let",
		farity:  1,
		ismulti: true,
		doc:     "Create a local scope with bindings",
		ftype:   special,
		argsStr: "(bindings . body)",
	},
	{
		name:    "loop",
		farity:  1,
		ismulti: true,
		doc:     "Loop forever",
		ftype:   special,
		argsStr: "(() . body)",
	},
	{
		name:    "or",
		farity:  0,
		ismulti: true,
		doc:     "Boolean or",
		ftype:   special,
		argsStr: "(() . xs)",
		examples: `> (or)
;; => false
> (or t t)
;; => true
> (or t t ())
;; => t`,
	},
	{
		name:    "quote",
		farity:  1,
		ismulti: false,
		doc:     "Quote an expression",
		ftype:   special,
		argsStr: "(x)",
		examples: `> (quote foo)
foo
> (quote (1 2 3))
(1 2 3)
> '(1 2 3)
(1 2 3)
`,
	},
	{
		name:    "syntax-quote",
		farity:  1,
		ismulti: false,
		doc:     "Syntax-quote an expression",
		ftype:   special,
		argsStr: "(x)",
		examples: `> (syntax-quote foo)
foo
> (syntax-quote (1 2 3 4))
(1 2 3 4)
> (syntax-quote (1 (unquote (+ 1 1)) (splicing-unquote (list 3 4))))
(1 2 3 4)
` + "> `(1 ~(1 1) ~@(list 3 4))" + `
(1 2 3 4)
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
		if doc.car.(*ConsCell).car.Equal(Atom{"examples"}) {
			return doc.car.(*ConsCell).cdr.(*ConsCell)
		}
		doc = doc.cdr.(*ConsCell)
	}
}

func formatLambdaArgs(args []string, restArg string) string {
	if restArg == "" {
		return fmt.Sprintf("(%s)", strings.Join(args, " "))
	}
	if len(args) == 0 {
		return fmt.Sprintf("(() . %s)", restArg)
	}
	return fmt.Sprintf("(%s . %s)", strings.Join(args, " "), restArg)
}

func examplesToString(examples *ConsCell, e *env) string {
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

func availableForms(e *env) []formRec {
	// Special forms - only need to add formatted column description:
	out := []formRec{}
	for _, form := range specialForms {
		form.columnDoc = formatFunctionInfo(form.name, form.doc, form.farity, form.ismulti, true, false, false)
		out = append(out, form)
	}

	// Builtins:
	for _, builtin := range builtins {
		out = append(out, formRec{
			name:     builtin.Name,
			farity:   builtin.FixedArity,
			ismulti:  builtin.NAry,
			doc:      builtin.Docstring,
			ftype:    native,
			argsStr:  builtin.ArgString,
			examples: examplesToString(builtin.Examples, e),
			columnDoc: formatFunctionInfo(builtin.Name,
				builtin.Docstring,
				builtin.FixedArity,
				builtin.NAry,
				false,
				false,
				true),
		})
	}
	// User-defined / internal l1 functions:
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
		examples := examplesToString(functionExamplesFromDoc(*l), e)
		if ok {
			out = append(out, formRec{
				name:     lambdaName,
				farity:   len(l.args),
				ismulti:  l.restArg != "",
				doc:      functionDescriptionFromDoc(*l),
				ftype:    ftype,
				argsStr:  formatLambdaArgs(l.args, l.restArg),
				examples: examples,
				columnDoc: formatFunctionInfo(lambdaName,
					functionDescriptionFromDoc(*l),
					len(l.args),
					l.restArg != "",
					false,
					l.isMacro,
					false),
			})
		}

	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].name < out[j].name
	})
	return out
}

func codeQuote(s string) string {
	return fmt.Sprintf("`%s`", s)
}

func longDocStr(e *env) string {
	sortedForms := availableForms(e)
	summary := fmt.Sprintf("# API Index\n%d forms available:", len(sortedForms))
	for _, form := range sortedForms {
		nameStr := fmt.Sprintf("`%s`", form.name)
		if form.ftype == macro {
			nameStr = fmt.Sprintf("*`%s`*", form.name)
		} else if form.ftype == special {
			nameStr = fmt.Sprintf("**`%s`**", form.name)
		}
		summary += fmt.Sprintf("\n[%s](#%s)", nameStr, form.name)
	}
	summary += "\n# Operators\n"
	outStrs := []string{summary}
	for _, doc := range sortedForms {
		isMulti := " "
		if doc.ismulti {
			isMulti = "+"
		}
		examples := ""
		if doc.examples != "" {
			examples = fmt.Sprintf("\n### Examples\n\n```\n%s\n```\n", doc.examples)
		}
		outStrs = append(outStrs, fmt.Sprintf(`
## %s

%s

Type: %s

Arity: %d%s

Args: %s

%s
-----------------------------------------------------
		`,
			codeQuote(doc.name),
			capitalize(doc.doc),
			doc.ftype,
			doc.farity,
			isMulti,
			fmt.Sprintf("`%s`", doc.argsStr),
			examples))
	}
	return strings.Join(outStrs, "\n")
}

func shortDocStr(e *env) string {
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
	sortedForms := availableForms(e)
	for _, doc := range sortedForms {
		outStrs = append(outStrs, doc.columnDoc)
	}
	return strings.Join(outStrs, "\n")
}

// a map... my kingdom for a map...
func formsAsSexprList(e *env) []Sexpr {
	out := []Sexpr{}
	for _, form := range availableForms(e) {
		out = append(out, Atom{form.name})
	}
	return out
}
