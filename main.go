package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/eigenhombre/lexutil"
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
			fmt.Printf("%v\n", err)
			return false
		}
		if doPrint {
			fmt.Printf("%v\n", res)
		}
	}
	return true
}

func lexParseEval(s string, e env, doPrint bool) bool {
	got, err := lexAndParse(s)
	if err != nil {
		fmt.Printf("%v\n", err)
		return false
	}
	return evalExprs(got, e, false)
}

func repl(e env) {
	for {
		fmt.Print("> ")
		tokens := []lexutil.LexItem{}
	Inner:
		for {
			s, err := readLine()
			switch err {
			case nil:
				these := lexItems(s)
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

func main() {
	globals := mkEnv(nil)
	if len(os.Args) > 1 {
		bytes, err := os.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		if !lexParseEval(string(bytes), globals, false) {
			os.Exit(1)
		}
		return
	}
	repl(globals)
}
