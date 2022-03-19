package main

import "fmt"

// Atom is the primitive symbolic type.
type Atom struct {
	s string
}

func (a Atom) String() string {
	return a.s
}

// Eval for atom returns the atom if it's the truth value; otherwise, it looks
// up the value in the environment.
func (a Atom) Eval(e *env) (Sexpr, error) {
	if a.s == "t" {
		return a, nil
	}
	ret, ok := e.Lookup(a.s)
	if ok {
		return ret, nil
	}
	ret, ok = builtins[a.s]
	if ok {
		return ret, nil
	}
	return nil, fmt.Errorf("unknown symbol: %s", a.s)
}

// Equal returns true if the receiver and the arg are both atoms and have the
// same name
func (a Atom) Equal(b Sexpr) bool {
	if b, ok := b.(Atom); ok {
		return a.s == b.s
	}
	return false
}
