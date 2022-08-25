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
	longDoc   string
	columnDoc string
	ftype     string
	argsStr   string
}

// When you add a special form to eval, you should add it here as well.s
var specialForms = []formRec{
	{"and", 0, true, "Boolean and", "", "", special,
		"(() . xs)"},
	{"cond", 0, true, "Conditional branching", "", "", special,
		"(() . pairs)"},
	{"def", 2, false, "Set a value", "", "", special,
		"(x val)"},
	{"defn", 2, true, "Create and name a function", "", "", special,
		"(f args . body)"},
	{"defmacro", 2, true, "Create and name a macro", "", "", special,
		"(f args . body)"},
	{"errors", 1, true, "Error checking (for tests)", "", "", special,
		"(x . xs)"},
	{"lambda", 1, true, "Create a function", "", "", special,
		"(args . body) or (name args . body)"},
	{"let", 1, true, "Create a local scope", "", "", special,
		"(bindings . body)"},
	{"loop", 1, true, "Loop forever", "", "", special,
		"(() . body)"},
	{"or", 0, true, "Boolean or", "", "", special,
		"(() . xs)"},
	{"quote", 1, false, "Quote an expression", "", "", special,
		"(x)"},
	{"syntax-quote", 1, false, "Syntax-quote an expression", "", "", special,
		"(x)"},
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

func formatLambdaArgs(args []string, restArg string) string {
	if restArg == "" {
		return fmt.Sprintf("(%s)", strings.Join(args, " "))
	}
	if len(args) == 0 {
		return fmt.Sprintf("(() . %s)", restArg)
	}
	return fmt.Sprintf("(%s . %s)", strings.Join(args, " "), restArg)
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
			name:    builtin.Name,
			farity:  builtin.FixedArity,
			ismulti: builtin.NAry,
			doc:     builtin.Docstring,
			ftype:   native,
			argsStr: builtin.ArgString,
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
		if ok {
			out = append(out, formRec{
				name:    lambdaName,
				farity:  len(l.args),
				ismulti: l.restArg != "",
				doc:     functionDescriptionFromDoc(*l),
				ftype:   ftype,
				argsStr: formatLambdaArgs(l.args, l.restArg),
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
		summary += fmt.Sprintf("\n[`%s`](#%s)", form.name, form.name)
	}
	summary += "\n# Operators\n"
	outStrs := []string{summary}
	for _, doc := range sortedForms {
		isMulti := " "
		if doc.ismulti {
			isMulti = "+"
		}
		outStrs = append(outStrs, fmt.Sprintf(`
## %s

%s

Type: %s

Arity: %d%s

Args: %s

-----------------------------------------------------
		`,
			codeQuote(doc.name),
			capitalize(doc.doc),
			doc.ftype,
			doc.farity,
			isMulti,
			fmt.Sprintf("`%s`", doc.argsStr)))
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
