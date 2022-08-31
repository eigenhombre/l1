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

func evalExprs(exprs []Sexpr, e env, doPrint bool) bool {
	for _, g := range exprs {
		res, err := eval(g, &e)
		if err != nil {
			fmt.Printf("ERROR in '%s':\n%v\n", g, err)
			return false
		}
		if doPrint {
			fmt.Printf("%v\n", res)
		}
	}
	return true
}

func lexParseEval(s string, e env, doPrint bool) bool {
	got, err := lexAndParse(strings.Split(s, "\n"))
	if err != nil {
		fmt.Printf("%v\n", err)
		return false
	}
	return evalExprs(got, e, false)
}

func repl(e env) {
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
				if isBalanced(tokens) {
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
	return globals
}

func main() {
	var versionFlag, docFlag, longDocFlag bool
	var cpuProfile string
	flag.BoolVar(&versionFlag, "v", false, "Get l1 version")
	flag.StringVar(&cpuProfile, "p", "", "Write CPU profile to file")
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

	if !lexParseEval(rawCore, globals, false) {
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

	files := flag.Args()
	if len(files) > 0 {
		for _, file := range files {
			bytes, err := os.ReadFile(file)
			if err != nil {
				panic(err)
			}
			if !lexParseEval(string(bytes), globals, false) {
				os.Exit(1)
			}
		}
		return
	}
	repl(globals)
}
