package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// Builtin represents a function with a native (Go) implementation.
type Builtin struct {
	Name string
	Fn   func([]Sexpr) (Sexpr, error)
	// Function must take at least this many arguments
	FixedArity int
	// Function can take more arguments
	NAry      bool
	Docstring string
}

func (b Builtin) String() string {
	return fmt.Sprintf("<builtin: %s>", b.Name)
}

// Eval for builtin returns itself.
func (b Builtin) Eval(e *env) (Sexpr, error) {
	return b, nil
}

// Equal returns true if the receiver and the arg are both builtins and have the
// same name.
func (b Builtin) Equal(o Sexpr) bool {
	if o, ok := o.(Builtin); ok {
		return b.Name == o.Name
	}
	return false
}

func doHelp(out io.Writer) {
	fmt.Fprintln(out, "Builtins and Special Forms:")
	fmt.Fprintln(out, "      Name  Arity    Description")
	type fnDoc struct {
		name      string
		farity    int
		ismulti   bool
		doc       string
		isSpecial bool
	}
	forms := []fnDoc{
		{"def", 2, false, "Set a value", true},
		{"lambda", 1, true, "Create a function", true},
		{"cond", 0, true, "Conditional branching", true},
		{"errors", 1, true, "Error checking (for tests)", true},
		{"quote", 1, false, "Quote an expression", true}}
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
}

func compareMultipleNums(cmp func(a, b Number) bool, args []Sexpr) (Sexpr, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("missing argument")
	}
	first, ok := args[0].(Number)
	if !ok {
		return nil, fmt.Errorf("'%s' is not a number", args[0])
	}
	last := first
	for i := 1; i < len(args); i++ {
		num, ok := args[i].(Number)
		if !ok {
			return nil, fmt.Errorf("'%s' is not a number", args[i])
		}
		if !cmp(num, last) {
			return Nil, nil
		}
		last = num
	}
	return Atom{"t"}, nil
}

// moving `builtins` into `init` avoids initialization loop for doHelp:
var builtins map[string]*Builtin

func init() {
	builtins = map[string]*Builtin{
		"+": {
			Name:       "+",
			Docstring:  "Add 0 or more numbers",
			FixedArity: 0,
			NAry:       true,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) == 0 {
					return Num(0), nil
				}
				sum := Num(0)
				for _, arg := range args {
					n, ok := arg.(Number)
					if !ok {
						return nil, fmt.Errorf("expected number, got '%s'", arg)
					}
					sum = sum.Add(n)
				}
				return sum, nil
			},
		},
		"-": {
			Name:       "-",
			Docstring:  "Subtract 0 or more numbers from the first argument",
			FixedArity: 1,
			NAry:       true,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) == 0 {
					return nil, fmt.Errorf("missing argument")
				}
				sum, ok := args[0].(Number)
				if !ok {
					return nil, fmt.Errorf("expected number, got '%s'", args[0])
				}
				if len(args) == 1 {
					return args[0].(Number).Neg(), nil
				}
				for _, arg := range args[1:] {
					n, ok := arg.(Number)
					if !ok {
						return nil, fmt.Errorf("expected number, got '%s'", arg)
					}
					sum = sum.Sub(n)
				}
				return sum, nil
			},
		},
		"*": {
			Name:       "*",
			Docstring:  "Multiply 0 or more numbers",
			FixedArity: 0,
			NAry:       true,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) == 0 {
					return Num(1), nil
				}
				prod := Num(1)
				for _, arg := range args {
					n, ok := arg.(Number)
					if !ok {
						return nil, fmt.Errorf("expected number, got '%s'", arg)
					}
					prod = prod.Mul(n)
				}
				return prod, nil
			},
		},
		"/": {
			Name:       "/",
			Docstring:  "Divide the first argument by the rest",
			FixedArity: 2,
			NAry:       true,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) < 1 {
					return nil, fmt.Errorf("missing argument")
				}
				quot, ok := args[0].(Number)
				if !ok {
					return nil, fmt.Errorf("expected number, got '%s'", args[0])
				}
				for _, arg := range args[1:] {
					if arg.Equal(Num(0)) {
						return nil, fmt.Errorf("division by zero")
					}
					n, ok := arg.(Number)
					if !ok {
						return nil, fmt.Errorf("expected number, got '%s'", arg)
					}
					quot = quot.Div(n)
				}
				return quot, nil
			},
		},
		"=": {
			Name:       "=",
			Docstring:  "Return t if the arguments are equal, () otherwise",
			FixedArity: 1,
			NAry:       true,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) < 1 {
					return nil, fmt.Errorf("missing argument")
				}
				for _, arg := range args[1:] {
					if !args[0].Equal(arg) {
						return Nil, nil
					}
				}
				return Atom{"t"}, nil
			},
		},
		"<": {
			Name:       "<",
			Docstring:  "Return t if the arguments are in strictly increasing order, () otherwise",
			FixedArity: 1,
			NAry:       true,
			Fn: func(args []Sexpr) (Sexpr, error) {
				return compareMultipleNums(func(a, b Number) bool {
					return b.Less(a)
				}, args)
			},
		},
		"<=": {
			Name:       "<=",
			Docstring:  "Return t if the arguments are in increasing (or qual) order, () otherwise",
			FixedArity: 1,
			NAry:       true,
			Fn: func(args []Sexpr) (Sexpr, error) {
				return compareMultipleNums(func(a, b Number) bool {
					return b.LessEqual(a)
				}, args)
			},
		},
		">": {
			Name:       ">",
			Docstring:  "Return t if the arguments are in strictly decreasing order, () otherwise",
			FixedArity: 1,
			NAry:       true,
			Fn: func(args []Sexpr) (Sexpr, error) {
				return compareMultipleNums(func(a, b Number) bool {
					return b.Greater(a)
				}, args)
			},
		},
		">=": {
			Name:       ">=",
			Docstring:  "Return t if the arguments are in decreasing (or equal) order, () otherwise",
			FixedArity: 1,
			NAry:       true,
			Fn: func(args []Sexpr) (Sexpr, error) {
				return compareMultipleNums(func(a, b Number) bool {
					return b.GreaterEqual(a)
				}, args)
			},
		},
		"car": {
			Name:       "car",
			Docstring:  "Return the first element of a list",
			FixedArity: 1,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("missing argument")
				}
				carCons, ok := args[0].(*ConsCell)
				if !ok {
					return nil, fmt.Errorf("'%s' is not a list", args[0])
				}
				if carCons == Nil {
					return Nil, nil
				}
				return carCons.car, nil
			},
		},
		"cdr": {
			Name:       "cdr",
			Docstring:  "Return a list with the first element removed",
			FixedArity: 1,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("missing argument")
				}
				cdrCons, ok := args[0].(*ConsCell)
				if !ok {
					return nil, fmt.Errorf("'%s' is not a list", args[0])
				}
				if cdrCons == Nil {
					return Nil, nil
				}
				return cdrCons.cdr, nil
			},
		},
		"cons": {
			Name:       "cons",
			Docstring:  "Add an element to the front of a (possibly empty) list",
			FixedArity: 2,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 2 {
					return nil, fmt.Errorf("missing argument")
				}
				return Cons(args[0], args[1].(*ConsCell)), nil
			},
		},
		"atom": {
			Name:       "atom",
			Docstring:  "Return true if the argument is an atom, false otherwise",
			FixedArity: 1,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("missing argument")
				}
				_, ok := args[0].(Atom)
				if ok {
					return Atom{"t"}, nil
				}
				return Nil, nil
			},
		},
		"print": {
			Name:       "print",
			Docstring:  "Print the arguments",
			FixedArity: 0,
			NAry:       true,
			Fn: func(args []Sexpr) (Sexpr, error) {
				strArgs := []string{}
				for _, arg := range args {
					strArgs = append(strArgs, arg.String())
				}
				fmt.Println(strings.Join(strArgs, " "))
				return Nil, nil
			},
		},
		"help": {
			Name:       "help",
			Docstring:  "Print this message",
			FixedArity: 0,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				doHelp(os.Stdout)
				return Nil, nil
			},
		},
		"is": {
			Name:       "is",
			Docstring:  "Assert that the argument is truthy (not ())",
			FixedArity: 1,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("missing argument")
				}
				if args[0] == Nil {
					return nil, fmt.Errorf("'is' assertion failed")
				}
				return args[0], nil
			},
		},
		"test": {
			Name:       "test",
			Docstring:  "Establish a testing block (return last expression)",
			FixedArity: 0,
			NAry:       true,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) == 0 {
					return Nil, nil
				}
				return args[len(args)-1], nil
			},
		},
		"len": {
			Name:       "len",
			Docstring:  "Return the length of a list",
			FixedArity: 1,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("len expects a single argument")
				}
				list, ok := args[0].(*ConsCell)
				if !ok {
					return nil, fmt.Errorf("'%s' is not a list", args[0])
				}
				count := 0
				for list != nil {
					count++
					list = list.cdr.(*ConsCell)
				}
				return Num(count), nil
			},
		},
		"not": {
			Name:       "not",
			Docstring:  "Return t if the argument is nil, () otherwise",
			FixedArity: 1,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("not expects a single argument")
				}
				if args[0] == Nil {
					return Atom{"t"}, nil
				}
				return Nil, nil
			},
		},
		"split": {
			Name:      "split",
			Docstring: "Split an atom or number into a list of single-digit numbers or single-character atoms",
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("split expects a single argument")
				}
				switch s := args[0].(type) {
				case Atom:
					return listOfChars(s.String()), nil
				case Number:
					return listOfNums(s.String())
				default:
					return nil, fmt.Errorf("split expects an atom or a number")
				}
			},
		},
		"fuse": {
			Name:       "fuse",
			Docstring:  "Fuse a list of numbers or atoms into a single atom",
			FixedArity: 1,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("fuse expects a single argument")
				}
				if args[0] == Nil {
					return Nil, nil
				}
				switch s := args[0].(type) {
				case *ConsCell:
					cons := s
					var str string
					for cons != nil {
						this := cons.car.String()
						str += this
						if cons.cdr == nil {
							break
						}
						cons = cons.cdr.(*ConsCell)
					}
					// if first rune is a digit, return a Number
					firstRune, _ := utf8.DecodeRuneInString(str)
					if unicode.IsDigit(firstRune) {
						return Num(str), nil
					}
					return Atom{str}, nil
				default:
					return nil, fmt.Errorf("fuse expects a list")
				}
			},
		},
		"neg?": {
			Name:       "neg?",
			Docstring:  "Return true if the (numeric) argument is negative, else ()",
			FixedArity: 1,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("neg? expects a single argument")
				}
				num, ok := args[0].(Number)
				if !ok {
					return nil, fmt.Errorf("'%s' is not a number", args[0])
				}
				if num.Less(Num(0)) {
					return Atom{"t"}, nil
				}
				return Nil, nil
			},
		},
		"pos?": {
			Name:       "pos?",
			Docstring:  "Return true if the (numeric) argument is positive, else ()",
			FixedArity: 1,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("pos? expects a single argument")
				}
				num, ok := args[0].(Number)
				if !ok {
					return nil, fmt.Errorf("'%s' is not a number", args[0])
				}
				if num.Greater(Num(0)) {
					return Atom{"t"}, nil
				}
				return Nil, nil
			},
		},
		"randigits": {
			Name:       "randigits",
			Docstring:  "Return a list of random digits of the given length",
			FixedArity: 1,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("randigits expects a single argument")
				}
				if _, ok := args[0].(Number); !ok {
					return nil, fmt.Errorf("randigits expects a number")
				}
				bigint, ok := args[0].(Number)
				if !ok {
					return nil, fmt.Errorf("randigits expects a number")
				}
				n := bigint.bi.Uint64()
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				digits := ""
				for i := uint64(0); i < n; i++ {
					digits += strconv.Itoa(r.Intn(10))
				}
				lon, err := listOfNums(string(digits))
				if err != nil {
					return nil, err
				}
				return lon, nil
			},
		},
		"apply": {
			Name:       "apply",
			Docstring:  "Apply a function to a list of arguments",
			FixedArity: 2,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 2 {
					return nil, fmt.Errorf("apply expects exactly two arguments")
				}
				fnArgs := []Sexpr{}
				start := args[1]
				for start != Nil {
					cons, ok := start.(*ConsCell)
					if !ok {
						return nil, fmt.Errorf("'%s' is not a list", start)
					}
					fnArgs = append(fnArgs, cons.car)
					start = cons.cdr
				}
				return applyFn(args[0], fnArgs)
			},
		},
		"zero?": {
			Name:       "zero?",
			Docstring:  "Return t if the argument is zero, () otherwise",
			FixedArity: 1,
			NAry:       false,
			Fn: func(args []Sexpr) (Sexpr, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("zero? expects a single argument")
				}
				num, ok := args[0].(Number)
				if !ok {
					return Nil, nil
				}
				if num.Equal(Num("0")) {
					return Atom{"t"}, nil
				}
				return Nil, nil
			},
		},
	}
}

// listOfChars returns a list of single-character atoms from another, presumably
// longer atom; used by `split`
func listOfChars(s string) *ConsCell {
	if len(s) == 0 {
		return nil
	}
	return Cons(Atom{s[0:1]}, listOfChars(s[1:]))
}

// listOfNums returns a list of single-digit numbers from another, presumably
// longer number; used by `split`; if the input represents a negative number,
// the first digit is negative:
func listOfNums(s string) (*ConsCell, error) {
	if len(s) == 0 {
		return nil, nil
	}
	if s[0] == '-' {
		if len(s) < 2 {
			return nil, fmt.Errorf("unexpected end of input")
		}
		lon, err := listOfNums(s[2:])
		if err != nil {
			return nil, err
		}
		return Cons(Num(s[0:2]), lon), nil
	}
	lon, err := listOfNums(s[1:])
	if err != nil {
		return nil, err
	}
	return Cons(Num(s[0:1]), lon), nil
}
