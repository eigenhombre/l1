package main

import "fmt"

// env stores a local environment, possibly pointing to a caller's environment.
type env struct {
	syms   map[string]Sexpr
	parent *env
}

// mkEnv makes a new Env.
func mkEnv(parent *env) env {
	return env{
		syms:   map[string]Sexpr{},
		parent: parent,
	}
}

// Lookup returns the value of a symbol in an environment or its parent(s).
func (e *env) Lookup(s string) (Sexpr, bool) {
	if e.syms[s] != nil {
		return e.syms[s], true
	}
	if e.parent != nil {
		return e.parent.Lookup(s)
	}
	return nil, false
}

// Set sets the value of a symbol in an environment.
func (e *env) Set(s string, v Sexpr) {
	e.syms[s] = v
}

func (e *env) String() string {
	ret := ""
	for k, v := range e.syms {
		ret += fmt.Sprintf("%s=%s;", k, v)
	}
	return ret
}
