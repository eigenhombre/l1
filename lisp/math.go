package lisp

import (
	"fmt"
	"math/big"
)

// Number wraps a big.Int (for now)
type Number struct {
	bi big.Int
}

// String returns the string representation of the number.
func (n Number) String() string {
	return n.bi.Text(10)
}

// Add returns the sum of the two numbers.
func (n Number) Add(o Number) Number {
	var ni big.Int = n.bi
	result := big.NewInt(0)
	result.Add(&ni, &o.bi)
	return Number{*result}
}

// Sub returns the difference of the two numbers.
func (n Number) Sub(o Number) Number {
	var ni big.Int = n.bi
	result := big.NewInt(0)
	result.Sub(&ni, &o.bi)
	return Number{*result}
}

// Mul returns the product of the two numbers.
func (n Number) Mul(o Number) Number {
	var ni big.Int = n.bi
	result := big.NewInt(0)
	result.Mul(&ni, &o.bi)
	return Number{*result}
}

// Div returns the (integer) quotient of the two numbers.
func (n Number) Div(o Number) Number {
	var ni big.Int = n.bi
	result := big.NewInt(0)
	result.Div(&ni, &o.bi)
	return Number{*result}
}

// Rem returns the remainder of the division of two numbers.
func (n Number) Rem(o Number) Number {
	var ni big.Int = n.bi
	result := big.NewInt(0)
	result.Rem(&ni, &o.bi)
	return Number{*result}
}

// Equal returns true if the two numbers are equal.
func (n Number) Equal(o Sexpr) bool {
	if o, ok := o.(Number); ok {
		return n.bi.Cmp(&o.bi) == 0
	}
	return false
}

// Less returns true if the first number is less than the second.
func (n Number) Less(o Number) bool {
	return n.bi.Cmp(&o.bi) < 0
}

// LessEqual returns true if the first number is <= the second.
func (n Number) LessEqual(o Number) bool {
	return n.bi.Cmp(&o.bi) <= 0
}

// Greater returns true if the first number is greater than the second.
func (n Number) Greater(o Number) bool {
	return n.bi.Cmp(&o.bi) > 0
}

// GreaterEqual returns true if the first number is >= the second.
func (n Number) GreaterEqual(o Number) bool {
	return n.bi.Cmp(&o.bi) >= 0
}

// Neg returns the negative of the number.
func (n Number) Neg() Number {
	var ni big.Int = n.bi
	result := big.NewInt(0)
	result.Neg(&ni)
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
