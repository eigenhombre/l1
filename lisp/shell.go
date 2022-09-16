package lisp

import (
	"bytes"
	"os/exec"
	"strings"
	"syscall"
)

func splitOnSpace(c rune) bool {
	return c == ' ' || c == '\t'
}

func shapeOutput(s string) *ConsCell {
	ret := []Sexpr{}
	lines := strings.Split(strings.Trim(s, "\n"), "\n")
	for _, l := range lines {
		ret = append(ret, stringsToList(strings.FieldsFunc(l, splitOnSpace)...))
	}
	return list(ret...)
}

func doShell(arg Sexpr) (Sexpr, error) {
	cmdCons, ok := arg.(*ConsCell)
	if !ok {
		return nil, baseErrorf("shell argument must be a nonempty list of strings")
	}
	cmdStrings := []string{}
	for cmdCons != Nil {
		switch t := cmdCons.car.(type) {
		case Atom:
			cmdStrings = append(cmdStrings, t.s)
		case Number:
			cmdStrings = append(cmdStrings, t.bi.String())
		default:
			return nil, baseErrorf("shell argument must be a nonempty list of strings")
		}
		cmdCons, ok = cmdCons.cdr.(*ConsCell)
		if !ok {
			return nil, baseErrorf("shell argument must be a nonempty list of strings")
		}
	}
	if len(cmdStrings) == 0 {
		return nil, baseErrorf("shell argument must be a nonempty list of strings")
	}
	cmdStdout := &bytes.Buffer{}
	cmdStderr := &bytes.Buffer{}
	cmd := exec.Command(cmdStrings[0], cmdStrings[1:]...)
	cmd.Stdout = cmdStdout
	cmd.Stderr = cmdStderr
	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
		} else {
			return nil, baseErrorf("error running shell command: %s", err)
		}
	} else {
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
	}
	return list(
		shapeOutput(cmdStdout.String()),
		shapeOutput(cmdStderr.String()),
		Num(waitStatus.ExitStatus()),
	), nil
}
