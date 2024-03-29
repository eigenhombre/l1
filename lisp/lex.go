package lisp

import (
	"fmt"
	"strings"

	"github.com/eigenhombre/lexutil"
)

// Use with lexutil.go (which should eventually be its own package).

// Token is a lexeme with a line number.
type Token struct {
	lexeme lexutil.LexItem
	line   int
}

// Token Types:
const (
	itemNumber lexutil.ItemType = iota
	itemAtom
	itemLeftParen
	itemRightParen
	itemForwardQuote
	itemSyntaxQuote
	itemUnquote
	itemSplicingUnquote
	itemDot
	itemCommentNext
	itemShebang
	itemError
)

// Human-readable versions of above:
var typeMap = map[lexutil.ItemType]string{
	itemNumber:          "NUM",
	itemAtom:            "ATOM",
	itemLeftParen:       "LP",
	itemRightParen:      "RP",
	itemForwardQuote:    "QUOTE",
	itemSyntaxQuote:     "SYNTAXQUOTE",
	itemUnquote:         "UNQUOTE",
	itemSplicingUnquote: "SPLICINGUNQUOTE",
	itemDot:             "DOT",
	itemCommentNext:     "COMMENTNEXT",
	itemShebang:         "SHEBANG",
	itemError:           "ERR",
}

// LexRepr returns a string representation of a known lexeme.
func LexRepr(i Token) string {
	switch i.lexeme.Typ {
	case itemNumber:
		return fmt.Sprintf("%s(%s)", typeMap[i.lexeme.Typ], i.lexeme.Val)
	case itemAtom:
		return fmt.Sprintf("%s(%s)", typeMap[i.lexeme.Typ], i.lexeme.Val)
	case itemLeftParen:
		return "LP"
	case itemRightParen:
		return "RP"
	case itemError:
		return fmt.Sprintf("%s(%s)", typeMap[i.lexeme.Typ], i.lexeme.Val)
	case itemForwardQuote:
		return "QUOTE"
	case itemSyntaxQuote:
		return "BACKQUOTE"
	case itemUnquote:
		return "UNQUOTE"
	case itemSplicingUnquote:
		return "SPLICINGUNQUOTE"
	case itemDot:
		return "DOT"
	case itemCommentNext:
		return "COMMENTNEXT"
	case itemShebang:
		return "SHEBANG"
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

func ignoreToEndOfLine(l *lexutil.Lexer) {
	for {
		if r := l.Next(); r == '\n' || r == lexutil.EOF {
			return
		}
	}
}

func lexStart(l *lexutil.Lexer) lexutil.StateFn {
	for {
		switch r := l.Next(); {
		case isSpace(r):
			l.Ignore()
		case r == ';':
			ignoreToEndOfLine(l)
		case r == lexutil.EOF:
			return nil
		case isDigit(r) || r == '-' || r == '+':
			l.Backup()
			return lexInt
		case r == '(':
			l.Emit(itemLeftParen)
		case r == ')':
			l.Emit(itemRightParen)
		case isAtomStart(r):
			return lexAtom
		case r == '\'':
			l.Emit(itemForwardQuote)
		case r == '`':
			l.Emit(itemSyntaxQuote)
		case r == '~':
			l.Backup()
			return lexUnquote
		case r == '.':
			l.Emit(itemDot)
		case r == '#':
			l.Backup()
			return lexHashSugar
		default:
			l.Errorf("unexpected character %q in input", itemError, r)
		}
	}
}

var disallowedForAtomAfterStart = " \t\n\r()~@#;`'"
var disallowedForAtomStart = "0123456789+-." + disallowedForAtomAfterStart

func isAtomStart(r rune) bool {
	return !(strings.ContainsRune(disallowedForAtomStart, r))
}

func acceptIf(l *lexutil.Lexer, f func(rune) bool) bool {
	for {
		x := l.Peek()
		if x == lexutil.EOF {
			return false
		}
		if f(x) {
			l.Next()
			continue
		}
		return false
	}
}

func lexAtom(l *lexutil.Lexer) lexutil.StateFn {
	acceptIf(l, isAtomStart)
	acceptIf(l, func(r rune) bool {
		return !(strings.ContainsRune(disallowedForAtomAfterStart, r))
	})
	l.Emit(itemAtom)
	return lexStart
}

func lexHashSugar(l *lexutil.Lexer) lexutil.StateFn {
	l.Accept("#")
	nextRune := l.Peek()
	if nextRune == '_' {
		l.Next()
		l.Emit(itemCommentNext)
		return lexStart
	} else if nextRune == '!' {
		l.Next()
		for {
			r := l.Next()
			if r == '\n' || r == lexutil.EOF {
				l.Backup()
				l.Emit(itemShebang)
				return lexStart
			}
		}
	}
	l.Errorf("unexpected character %q in input", itemError, nextRune)
	return lexStart
}

func lexUnquote(l *lexutil.Lexer) lexutil.StateFn {
	l.Accept("~")
	nextRune := l.Peek()
	if nextRune == '@' {
		l.Next()
		l.Emit(itemSplicingUnquote)
	} else {
		l.Emit(itemUnquote)
	}
	return lexStart
}

func lexInt(l *lexutil.Lexer) lexutil.StateFn {
	l.Accept("-+")
	nextRune := l.Peek()
	if isDigit(nextRune) {
		l.AcceptRun("0123456789")
		l.Emit(itemNumber)
		return lexStart
	}
	return lexAtom
}

// LexItems lexes a string into a slice of tokens.
func LexItems(ss []string) []Token {
	ret := []Token{}
	for line, s := range ss {
		l := lexutil.Lex("main", s, lexStart)
		for tok := range l.Items {
			// Programmers may be civilians, counting lines from 1 rather than
			// 0:
			ret = append(ret, Token{tok, line + 1})
		}
	}
	return ret
}

// IsBalanced returns true iff parens are balanced
func IsBalanced(tokens []Token) (bool, error) {
	level := 0
	for _, token := range tokens {
		switch token.lexeme.Typ {
		case itemLeftParen:
			level++
		case itemRightParen:
			level--
		}
	}
	if level < 0 {
		return false, fmt.Errorf("too many right parens")
	}
	return level == 0, nil
}
