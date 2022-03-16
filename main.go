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

func lexParseEvalPrint(s string, e env) {
	got, err := lexAndParse(s)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	for _, g := range got {
		_, err := g.Eval(&e)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
	}
}

func repl(e env) {
	for {
		fmt.Print("> ")
		s, err := readLine()
		switch err {
		case nil:
			lexParseEvalPrint(s, e)
		case io.EOF:
			fmt.Println()
			return
		default:
			panic(err)
		}
	}
}

func main() {
	globals := env{}
	if len(os.Args) > 1 {
		bytes, err := os.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		lexParseEvalPrint(string(bytes), globals)
		return
	}
	repl(globals)
}
