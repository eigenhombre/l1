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
	for {
		fmt.Print("> ")
		s, err := readLine()
		switch err {
		case nil:
			fmt.Println(eval(lexAndParse(s)[0], env{})) // fixme: handle multiple items
		case io.EOF:
			fmt.Println()
			return
		default:
			panic(err)
		}
	}
}
