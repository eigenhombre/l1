package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"strings"
)

func readLine() (string, error) {
	bio := bufio.NewReader(os.Stdin)
	// FIXME: don't discard hasMoreInLine
	line, _, err := bio.ReadLine()
	switch err {
	case nil:
		return string(line), nil
	default:
		return "", err
	}
}

func evalExprs(exprs []Sexpr, e *env, doPrint bool) error {
	for _, g := range exprs {
		res, err := eval(g, e)
		if err != nil {
			if doPrint {
				fmt.Printf("ERROR:\n%v\n", err)
			}
			return err
		}
		if doPrint {
			fmt.Printf("%v\n", res)
		}
	}
	return nil
}

func lexParseEval(s string, e *env, doPrint bool) error {
	got, err := lexAndParse(strings.Split(s, "\n"))
	if err != nil {
		return err
	}
	return evalExprs(got, e, false)
}

func repl(e *env) {
top:
	for {
		fmt.Print("> ")
		tokens := []token{}
	Inner:
		for {
			s, err := readLine()
			switch err {
			case nil:
				these := lexItems([]string{s})
				tokens = append(tokens, these...)
				bal, err := isBalanced(tokens)
				if err != nil {
					fmt.Printf("ERROR:\n%v\n", err)
					goto top
				}
				if bal {
					break Inner
				}
			case io.EOF:
				fmt.Println()
				return
			default:
				panic(err)
			}
		}
		exprs, err := parse(tokens)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		evalExprs(exprs, e, true)
	}
}

func initGlobals() env {
	globals := mkEnv(nil)
	globals.Set("SPACE", Atom{" "})
	globals.Set("NEWLINE", Atom{"\n"})
	globals.Set("TAB", Atom{"\t"})
	globals.Set("BANG", Atom{"!"})
	globals.Set("QMARK", Atom{"?"})
	globals.Set("PERIOD", Atom{"."})
	globals.Set("COMMA", Atom{","})
	globals.Set("COLON", Atom{":"})
	globals.Set("HASH", Atom{"#"})
	globals.Set("ATSIGN", Atom{"@"})
	globals.Set("CHECK", Atom{"✓"})
	return globals
}

func main() {
	var versionFlag, docFlag, longDocFlag bool
	var cpuProfile, evalExpr string
	flag.BoolVar(&versionFlag, "v", false, "Get l1 version")
	flag.StringVar(&cpuProfile, "p", "", "Write CPU profile to file")
	flag.StringVar(&evalExpr, "e", "", "Evaluate expression")
	flag.BoolVar(&docFlag, "doc", false, "Print documentation")
	flag.BoolVar(&longDocFlag, "longdoc", false, "Print documentation")

	flag.Parse()

	if versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	if cpuProfile != "" {
		f, err := os.Create(cpuProfile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	globals := initGlobals()

	err := lexParseEval(rawCore, &globals, false)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to load l1 core library!")
		os.Exit(1)
	}

	if docFlag {
		fmt.Println(shortDocStr(&globals))
		os.Exit(0)
	}
	if longDocFlag {
		fmt.Println(longDocStr(&globals))
		os.Exit(0)
	}
	if evalExpr != "" {
		err = lexParseEval(evalExpr, &globals, true)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	files := flag.Args()
	if len(files) > 0 {
		for _, file := range files {
			err := loadFile(&globals, file)
			if err != nil {
				fmt.Printf("ERROR:\n%v\n", err)
				os.Exit(1)
			}
		}
		return
	}
	repl(&globals)
}
