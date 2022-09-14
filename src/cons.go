package main

// ConsCell is a cons cell.  Use Cons to create one.
type ConsCell struct {
	car Sexpr
	cdr Sexpr
}

// A cons (list) can be used as an error, and consed
// to, to make a stacktrace:
func (c *ConsCell) Error() string {
	return c.String()
}

// Nil is the empty list / cons cell.  Cons with Nil to create a list
// of one item.
var Nil *ConsCell = nil

func (c *ConsCell) String() string {
	ret := "("
	car := c
	for {
		if car == Nil {
			break
		}
		ret += car.car.String()
		cdr, ok := car.cdr.(*ConsCell)
		if !ok {
			return ret + " . " + car.cdr.String() + ")"
		}
		if cdr != Nil {
			ret += " "
		}
		car = cdr
	}
	return ret + ")"
}

// Cons creates a cons cell.
func Cons(i Sexpr, cdr Sexpr) *ConsCell {
	return &ConsCell{i, cdr}
}

// Equal returns true iff the two S-expressions are equal cons-wise
func (c *ConsCell) Equal(o Sexpr) bool {
	_, ok := o.(*ConsCell)
	if !ok {
		return false
	}
	if c == Nil {
		return o == Nil
	}
	if o == Nil {
		return c == Nil
	}
	return c.car.Equal(o.(*ConsCell).car) && c.cdr.Equal(o.(*ConsCell).cdr)
}
