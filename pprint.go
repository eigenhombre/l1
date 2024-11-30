package main

import (
	"fmt"
	"strings"
)

func pprint(s Sexpr) string {
	c, ok := s.(*ConsCell)
	if !ok {
		return s.String()
	}
	return pprintCons(c, 0)
}

func pprintCons(c *ConsCell, indent int) string {
	fmt.Println("indent:", indent)
	firstTry := c.String()
	if len(firstTry)+indent < 80 {
		return c.String()
	}
	indentStr := strings.Repeat(" ", indent+1)
	head := c.car.String()
	ret := "("
	for c != Nil {
		if c.cdr == Nil {
			ret += head
			break
		}
		ret += head + "\n" + indentStr
		c = c.cdr.(*ConsCell)
		head = c.car.String()
	}
	return ret + ")"
}
