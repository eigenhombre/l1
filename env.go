package main

import (
	"fmt"
)

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

// EnvKeys returns the keys of an environment, including any parents' keys.
func EnvKeys(m *env) []string {
	ret := []string{}
	for k := range m.syms {
		ret = append(ret, k)
	}
	if m.parent != nil {
		ret = append(ret, EnvKeys(m.parent)...)
	}
	return ret
}

// Lookup returns the value of a symbol in an environment or its parent(s).
func (e *env) Lookup(s string) (Sexpr, bool) {
	if v, ok := e.syms[s]; ok {
		return v, true
	}
	if e.parent != nil {
		return e.parent.Lookup(s)
	}
	return nil, false
}

// Set sets the value of a symbol in an environment.
func (e *env) Set(s string, v Sexpr) error {
	if s == "t" {
		return baseError("cannot bind or set t")
	}
	e.syms[s] = v
	return nil
}

func (e *env) Update(s string, v Sexpr) error {
	if s == "t" {
		return baseError("cannot bind or set t")
	}
	if _, ok := e.syms[s]; ok {
		e.syms[s] = v
		return nil
	}
	if e.parent != nil {
		return e.parent.Update(s, v)
	}
	return baseErrorf("%s is not bound in any environment", s)
}

func (e *env) String() string {
	ret := ""
	for k, v := range e.syms {
		ret += fmt.Sprintf("%s=%s\n", k, v)
	}
	ret += "\n"
	if e.parent != nil {
		ret += fmt.Sprintf("PARENT: %s\n", e.parent.String())
	}
	return ret
}
