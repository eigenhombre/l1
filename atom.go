package main

// Atom is the primitive symbolic type.
type Atom struct {
	s string
}

func (a Atom) String() string {
	return a.s
}

// Equal returns true if the receiver and the arg are both atoms and have the
// same name
func (a Atom) Equal(b Sexpr) bool {
	if b, ok := b.(Atom); ok {
		return a.s == b.s
	}
	return false
}

// True is the generic truthy item.  Everything but Nil is true.  See also Nil
// in cons.go.
var True Atom = Atom{"t"}
