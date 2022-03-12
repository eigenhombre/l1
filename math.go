package main

import (
	"fmt"
	"math/big"
)

// Number wraps a big.Int (for now)
type Number struct {
	bi big.Int
}

// Eval of a number is just the number itself.
func (n Number) Eval(e *env) Sexpr {
	return n
}

// String returns the string representation of the number.
func (n Number) String() string {
	return n.bi.Text(10)
}

// Add returns the sum of the two numbers.
func (n Number) Add(o Number) Number {
	var ni big.Int = n.bi
	result := ni.Add(&ni, &o.bi)
	return Number{*result}
}

// Sub returns the difference of the two numbers.
func (n Number) Sub(o Number) Number {
	var ni big.Int = n.bi
	result := ni.Sub(&ni, &o.bi)
	return Number{*result}
}

// Mul returns the product of the two numbers.
func (n Number) Mul(o Number) Number {
	var ni big.Int = n.bi
	result := ni.Mul(&ni, &o.bi)
	return Number{*result}
}

// Div returns the (integer) quotient of the two numbers.
func (n Number) Div(o Number) Number {
	var ni big.Int = n.bi
	result := ni.Div(&ni, &o.bi)
	return Number{*result}
}

// Equals returns true if the two numbers are equal.
func (n Number) Equals(o Number) bool {
	return n.bi.Cmp(&o.bi) == 0
}

// Neg returns the negative of the number.
func (n Number) Neg() Number {
	var ni big.Int = n.bi
	result := ni.Neg(&ni)
	return Number{*result}
}

// Num is a `num` constructor, which can take a string or a
// ("normal") number.
func Num(ob interface{}) Number {
	var n Number
	switch s := ob.(type) {
	case string:
		n.bi.SetString(s, 10)
	case int:
		n.bi.SetInt64(int64(s))
	default:
		panic(fmt.Sprintf("Num: unknown type %T\n", ob))
	}
	return n
}
