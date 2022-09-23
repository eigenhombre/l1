package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/pprof"

	"github.com/eigenhombre/l1/lisp"
)

func repl(e *lisp.Env) {
top:
	for {
		fmt.Print("> ")
		tokens := []lisp.Token{}
	Inner:
		for {
			s, err := lisp.ReadLine()
			switch err {
			case nil:
				these := lisp.LexItems([]string{s})
				tokens = append(tokens, these...)
				bal, err := lisp.IsBalanced(tokens)
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
		exprs, err := lisp.Parse(tokens)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		lisp.EvalExprs(exprs, e, true)
	}
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
		fmt.Println(lisp.Version)
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

	globals := lisp.InitGlobals()

	err := lisp.LexParseEval(lisp.RawCore, &globals)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to load l1 core library!")
		os.Exit(1)
	}

	if docFlag {
		fmt.Println(lisp.ShortDocStr(&globals))
		os.Exit(0)
	}
	if longDocFlag {
		ld, err := lisp.LongDocStr(&globals)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(ld)
		os.Exit(0)
	}
	if evalExpr != "" {
		err = lisp.LexParseEval(evalExpr, &globals)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	files := flag.Args()
	if len(files) > 0 {
		for _, file := range files {
			err := lisp.LoadFile(&globals, file)
			if err != nil {
				fmt.Printf("ERROR:\n%v\n", err)
				os.Exit(1)
			}
		}
		return
	}
	repl(&globals)
}
