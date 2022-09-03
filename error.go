package main

import (
	"fmt"
	"strings"
)

func extendStacktrace(carList *ConsCell, err error) error {
	ret, ok := err.(*ConsCell)
	if !ok {
		return err
	}
	return Cons(carList, ret)
}

func extendWithSplitString(msg string, err error) error {
	return extendStacktrace(stringsToList(strings.Split(msg, " ")...), err)
}

func startStacktrace(carList *ConsCell) error {
	return Cons(carList, Nil)
}

func baseError(msg string) error {
	return startStacktrace(stringsToList(strings.Split(msg, " ")...))
}

func baseErrorf(format string, a ...interface{}) error {
	return startStacktrace(stringsToList(strings.Split(fmt.Sprintf(format, a...), " ")...))
}
