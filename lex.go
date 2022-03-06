package main

import (
	"fmt"
	"strings"
)

// Use with lexutil.go (which should eventually be its own package).

// Lexemes:
const (
	itemNumber itemType = iota
	itemAtom
	itemLeftParen
	itemRightParen
	itemError
)

// Human-readable versions of above:
var typeMap = map[itemType]string{
	itemNumber:     "NUM",
	itemAtom:       "ATOM",
	itemLeftParen:  "LP",
	itemRightParen: "RP",
}

func (i item) String() string {
	switch i.typ {
	case itemNumber:
		return fmt.Sprintf("%s(%s)", typeMap[i.typ], i.val)
	case itemAtom:
		return fmt.Sprintf("%s(%s)", typeMap[i.typ], i.val)
	case itemLeftParen:
		return "LP"
	case itemRightParen:
		return "RP"
	default:
		panic("bad item type")
	}
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isSpace(r rune) bool {
	return strings.ContainsRune(" \t\n\r", r)
}

func lexBetween(l *lexer) stateFn {
	for {
		switch r := l.next(); {
		case isSpace(r):
			l.ignore()
		case r == eof:
			return nil
		case isDigit(r) || r == '-' || r == '+':
			l.backup()
			return lexInt
		case r == '(':
			l.emit(itemLeftParen)
		case r == ')':
			l.emit(itemRightParen)
		default:
			l.backup()
			return lexAtom
		}
	}
}

func lexAtom(l *lexer) stateFn {
	var validAtomChars = ("0123456789abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"+*/-_!=<>")
	l.acceptRun(validAtomChars)
	l.emit(itemAtom)
	return lexBetween
}

func lexInt(l *lexer) stateFn {
	l.accept("-+")
	nextRune := l.peek()
	if isDigit(nextRune) {
		l.acceptRun("0123456789")
		l.emit(itemNumber)
		return lexBetween
	}
	return lexAtom
}

func lexItems(s string) []item {
	_, ch := lex("main", s, lexBetween)
	items := []item{}
	for tok := range ch {
		items = append(items, tok)
	}
	return items
}
