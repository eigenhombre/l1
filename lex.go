package main

import (
	"fmt"
	"strings"

	"github.com/eigenhombre/lexutil"
)

// Use with lexutil.go (which should eventually be its own package).

// Lexemes:
const (
	itemNumber lexutil.ItemType = iota
	itemAtom
	itemLeftParen
	itemRightParen
	itemError
)

// Human-readable versions of above:
var typeMap = map[lexutil.ItemType]string{
	itemNumber:     "NUM",
	itemAtom:       "ATOM",
	itemLeftParen:  "LP",
	itemRightParen: "RP",
}

// LexRepr returns a string representation of a known lexeme.
func LexRepr(i lexutil.LexItem) string {
	switch i.Typ {
	case itemNumber:
		return fmt.Sprintf("%s(%s)", typeMap[i.Typ], i.Val)
	case itemAtom:
		return fmt.Sprintf("%s(%s)", typeMap[i.Typ], i.Val)
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

func ignoreComment(l *lexutil.Lexer) {
	for {
		if r := l.Next(); r == '\n' || r == lexutil.EOF {
			return
		}
	}
}

func lexBetween(l *lexutil.Lexer) lexutil.StateFn {
	for {
		switch r := l.Next(); {
		case isSpace(r):
			l.Ignore()
		case r == ';':
			ignoreComment(l)
		case r == lexutil.EOF:
			return nil
		case isDigit(r) || r == '-' || r == '+':
			l.Backup()
			return lexInt
		case r == '(':
			l.Emit(itemLeftParen)
		case r == ')':
			l.Emit(itemRightParen)
		default:
			l.Backup()
			return lexAtom
		}
	}
}

func lexAtom(l *lexutil.Lexer) lexutil.StateFn {
	var validAtomChars = ("0123456789abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"+*/-_!=<>?$^[]{}&")
	l.AcceptRun(validAtomChars)
	l.Emit(itemAtom)
	return lexBetween
}

func lexInt(l *lexutil.Lexer) lexutil.StateFn {
	l.Accept("-+")
	nextRune := l.Peek()
	if isDigit(nextRune) {
		l.AcceptRun("0123456789")
		l.Emit(itemNumber)
		return lexBetween
	}
	return lexAtom
}

func lexItems(s string) []lexutil.LexItem {
	l := lexutil.Lex("main", s, lexBetween)
	ret := []lexutil.LexItem{}
	for tok := range l.Items {
		ret = append(ret, tok)
	}
	return ret
}

func isBalanced(tokens []lexutil.LexItem) bool {
	level := 0
	for _, token := range tokens {
		switch token.Typ {
		case itemLeftParen:
			level++
		case itemRightParen:
			level--
		}
	}
	return level == 0
}
