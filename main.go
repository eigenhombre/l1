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

func main() {
	globals := env{}
	for {
		fmt.Print("> ")
		s, err := readLine()
		switch err {
		case nil:
			got, err := lexAndParse(s)
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}
			fmt.Println(got[0].Eval(&globals)) // fixme: handle multiple items
		case io.EOF:
			fmt.Println()
			return
		default:
			panic(err)
		}
	}
}
