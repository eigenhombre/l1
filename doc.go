package main

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

type fnDoc struct {
	name      string
	farity    int
	ismulti   bool
	doc       string
	isSpecial bool
}

// When you add a special form to eval, you should add it here as well.s
var specialForms = []fnDoc{
	{"and", 0, true, "Boolean and", true},
	{"cond", 0, true, "Conditional branching", true},
	{"def", 2, false, "Set a value", true},
	{"defn", 2, true, "Create and name a function", true},
	{"defmacro", 2, true, "Create and name a macro", true},
	{"errors", 1, true, "Error checking (for tests)", true},
	{"lambda", 1, true, "Create a function", true},
	{"let", 1, true, "Create a local scope", true},
	{"loop", 1, true, "Loop forever", true},
	{"or", 0, true, "Boolean or", true},
	{"quote", 1, false, "Quote an expression", true},
	{"syntax-quote", 1, false, "Syntax-quote an expression", true},
}

const formatStr = "%14s %2s %5s  %s"

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
	return fmt.Sprintf(formatStr,
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

func doHelp(out io.Writer, e *env) {
	type namedDoc struct {
		name string
		doc  string
	}
	outStrs := []string{}
	sortedStrs := []namedDoc{}
	outStrs = append(outStrs,
		"l1 - a Lisp interpreter.\n",
		fmt.Sprintf(formatStr, "Name", "Type", "Arity", "Description"),
		fmt.Sprintf(formatStr, "----", "---", "----", "-----------"),
		"                S - special form",
		"                M - macro",
		"                N - native (Go) function",
		"                F - Lisp function\n",
	)

	// Special forms:
	for _, fn := range specialForms {
		sortedStrs = append(sortedStrs, namedDoc{fn.name,
			formatFunctionInfo(fn.name,
				fn.doc,
				fn.farity,
				fn.ismulti,
				fn.isSpecial,
				fn.isSpecial,
				false)})
	}
	// Builtins:
	for _, builtin := range builtins {
		sortedStrs = append(sortedStrs, namedDoc{builtin.Name,
			formatFunctionInfo(builtin.Name,
				builtin.Docstring,
				builtin.FixedArity,
				builtin.NAry,
				false,
				false,
				true)})
	}
	// User-defined / internal l1 functions:
	for _, lambdaName := range EnvKeys(e) {
		expr, _ := e.Lookup(lambdaName)
		l, ok := expr.(*lambdaFn)
		if ok {
			sortedStrs = append(sortedStrs, namedDoc{lambdaName,
				formatFunctionInfo(lambdaName,
					functionDescriptionFromDoc(*l),
					len(l.args),
					l.restArg != "",
					false,
					l.isMacro,
					false)})
		}
	}
	sort.Slice(sortedStrs, func(i, j int) bool {
		return sortedStrs[i].name < sortedStrs[j].name
	})
	for _, doc := range sortedStrs {
		outStrs = append(outStrs, doc.doc)
	}
	fmt.Fprintln(out, strings.Join(outStrs, "\n"))
}

func helpStr(e *env) string {
	helpBuf := bytes.NewBufferString("")
	doHelp(helpBuf, e)
	return helpBuf.String()
}
