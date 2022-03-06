package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLine() string {
	bio := bufio.NewReader(os.Stdin)
	// FIXME: don't discard hasMoreInLine
	line, _, err := bio.ReadLine()
	if err != nil {
		panic(err)
	}
	return string(line)
}

func main() {
	for {
		fmt.Print("> ")
		s := readLine()
		fmt.Println("\nLEXEMES______________________")
		for _, item := range lexItems(s) {
			fmt.Printf("%s('%s') ", typeMap[item.typ], item.val)
			fmt.Println()
		}
		fmt.Println("\nPARSED ITEMS________________")
		fmt.Println(lexAndParse(s))
	}
}
