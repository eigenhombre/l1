package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func lexParseEval(s string, e env, doPrint bool) {
	got, err := lexAndParse(s)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	for _, g := range got {
		res, err := g.Eval(&e)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		if doPrint {
			fmt.Printf("%v\n", res)
		}
	}
}

func repl(e env) {
	for {
		fmt.Print("> ")
		s, err := readLine()
		switch err {
		case nil:
			lexParseEval(s, e, true)
		case io.EOF:
			fmt.Println()
			return
		default:
			panic(err)
		}
	}
}

func main() {
	globals := mkEnv(nil)
	if len(os.Args) > 1 {
		bytes, err := os.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		lexParseEval(string(bytes), globals, false)
		return
	}
	repl(globals)
}
