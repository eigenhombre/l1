package main

import (
	"fmt"
	"io"
	"sort"
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
	{"errors", 1, true, "Error checking (for tests)", true},
	{"lambda", 1, true, "Create a function", true},
	{"let", 1, true, "Create a local scope", true},
	{"loop", 1, true, "Loop forever", true},
	{"or", 0, true, "Boolean or", true},
	{"quote", 1, false, "Quote an expression", true},
}

func doHelp(out io.Writer, e *env) {
	fmt.Fprintln(out, "Builtins and Special Forms:")
	fmt.Fprintln(out, "      Name  Arity    Description")
	forms := specialForms
	for _, builtin := range builtins {
		forms = append(
			forms,
			fnDoc{
				builtin.Name,
				builtin.FixedArity,
				builtin.NAry,
				builtin.Docstring,
				false})
	}
	// sort by name
	sort.Slice(forms, func(i, j int) bool {
		return forms[i].name < forms[j].name
	})
	for _, form := range forms {
		special := ""
		if form.isSpecial {
			special = "SPECIAL FORM: "
		}
		isMultiArity := " "
		if form.ismulti {
			isMultiArity = "+"
		}
		argstr := fmt.Sprintf("%d%s", form.farity, isMultiArity)
		fmt.Fprintf(
			out,
			"%10s %5s     %s%s\n",
			form.name,
			argstr,
			special,
			form.doc)
	}
	lambdaNames := []string{}
	for _, name := range EnvKeys(e) {
		expr, _ := e.Lookup(name)
		if _, ok := expr.(*lambdaFn); ok {
			lambdaNames = append(lambdaNames, name)
		}
	}
	sort.Slice(lambdaNames, func(i, j int) bool {
		return lambdaNames[i] < lambdaNames[j]
	})

	fmt.Fprint(out, "\n\nOther available functions:\n\n")
	for _, lambdaName := range lambdaNames {
		fmt.Fprint(out, lambdaDocString(lambdaName, e))
	}
}
