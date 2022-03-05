package main

import (
	"fmt"
	"math/big"
)

// sexpr is a general-purpose data structure for representing
// S-expressions; for now, it has String only, but it may have
// evaluable or other methods added later.
type sexpr interface {
	String() string
}

// consCell is a cons cell.  Use Cons to create one.
type consCell struct {
	car sexpr
	cdr sexpr
}

// Nil is the empty list / cons cell.  Cons with Nil to create a list
// of one item.
var Nil *consCell = nil

func (j *consCell) String() string {
	ret := "("
	for car := j; car != Nil; car = car.cdr.(*consCell) {
		if car.car == Nil {
			return ret + ")"
		}
		ret += car.car.String()
		if car.cdr != Nil {
			ret += " "
		}
	}
	return ret + ")"
}

// number is a big.Int, but narrower in its string representation.
type number struct {
	big.Int
}

// String returns the string representation of the number.
func (n number) String() string {
	return n.Text(10)
}

func Num(ob interface{}) number {
	var n number
	switch s := ob.(type) {
	case string:
		n.SetString(s, 10)
	case int:
		n.SetInt64(int64(s))
	default:
		panic(fmt.Sprintf("Num: unknown type %T\n", ob))
	}
	return n
}

func Cons(i sexpr, cdr *consCell) *consCell {
	return &consCell{i, cdr}
}
