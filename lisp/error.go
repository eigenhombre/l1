package lisp

import (
	"fmt"
	"strings"
)

func extendWithList(carList *ConsCell, err error) error {
	ret, ok := err.(*ConsCell)
	if !ok {
		return err
	}
	return Cons(carList, ret)
}

func extendError(msg string, err error) error {
	return extendWithList(stringsToList(strings.Split(msg, " ")...), err)
}

func startStacktrace(carList *ConsCell) error {
	return Cons(carList, Nil)
}

func baseError(msg string) error {
	return startStacktrace(stringsToList(strings.Split(msg, " ")...))
}

func baseErrorf(format string, a ...interface{}) error {
	return startStacktrace(stringsToList(
		strings.Split(fmt.Sprintf(format, a...), " ")...))
}
