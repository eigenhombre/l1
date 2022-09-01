package main

import "strings"

func extendStacktrace(carList *ConsCell, err error) error {
	ret, ok := err.(*ConsCell)
	if !ok {
		return err
	}
	return Cons(carList, ret)
}

func startStacktrace(carList *ConsCell) error {
	return Cons(carList, Nil)
}

func baseError(msg string) error {
	return startStacktrace(stringsToList(strings.Split(msg, " ")...))
}
