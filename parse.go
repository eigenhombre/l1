package main

import (
	"fmt"

	"github.com/eigenhombre/lexutil"
)

func handleQuoteItem(tokens []lexutil.LexItem, i int, operatorName string) (Sexpr, int, error) {
	if i >= len(tokens) {
		return nil, 0, fmt.Errorf("unexpected end of input; index=%d, tokens=%v", i, tokens)
	}
	nextParsed, incr, err := parseNext(tokens, i)
	if err != nil {
		return nil, 0, err
	}
	item := Cons(Atom{operatorName}, Cons(nextParsed, Nil))
	return item, incr, nil
}

func parseNext(tokens []lexutil.LexItem, i int) (Sexpr, int, error) {
	if i >= len(tokens) {
		return nil, 0, fmt.Errorf("unexpected end of input; index=%d, tokens=%v", i, tokens)
	}
	token := tokens[i]
	switch token.Typ {
	case itemNumber:
		return Num(token.Val), 1, nil
	case itemAtom:
		return Atom{token.Val}, 1, nil
	case itemForwardQuote:
		item, incr, err := handleQuoteItem(tokens, i+1, "quote")
		if err != nil {
			return nil, 0, err
		}
		return item, incr + 1, nil
	case itemSyntaxQuote:
		item, incr, err := handleQuoteItem(tokens, i+1, "syntax-quote")
		if err != nil {
			return nil, 0, err
		}
		return item, incr + 1, nil
	case itemUnquote:
		item, incr, err := handleQuoteItem(tokens, i+1, "unquote")
		if err != nil {
			return nil, 0, err
		}
		return item, incr + 1, nil
	case itemSplicingUnquote:
		item, incr, err := handleQuoteItem(tokens, i+1, "splicing-unquote")
		if err != nil {
			return nil, 0, err
		}
		return item, incr + 1, nil
	case itemCommentNext:
		item, incr, err := handleQuoteItem(tokens, i+1, "comment")
		if err != nil {
			return nil, 0, err
		}
		return item, incr + 1, nil
	case itemLeftParen:
		item, incr, err := parseList(tokens[i:])
		if err != nil {
			return nil, 0, err
		}
		return item, incr, nil
	case itemRightParen:
		return nil, 0, fmt.Errorf("unexpected right paren")
	default:
		return nil, 0, fmt.Errorf("unexpected lexeme '%s'", token.Val)
	}
}

func parse(tokens []lexutil.LexItem) ([]Sexpr, error) {
	ret := []Sexpr{}
	i := 0
	// Look for shebang, only at beginning of file:
	if len(tokens) > 0 && tokens[0].Typ == itemShebang {
		i++
	}
	for {
		if i >= len(tokens) {
			break
		}
		item, incr, err := parseNext(tokens, i)
		if err != nil {
			return nil, err
		}
		ret = append(ret, item)
		i += incr
	}
	return ret, nil
}

// parseList is used when a list has been detected in a slice of tokens.
func parseList(tokens []lexutil.LexItem) (Sexpr, int, error) {
	chunkEnd, endTok, err := listChunk(tokens)
	if err != nil {
		return nil, 0, err
	}
	if endTok.Typ == itemDot {
		carList, err := parse(tokens[1:chunkEnd])
		if err != nil {
			return nil, 0, err
		}
		chunk2End, err := dotChunk(tokens[chunkEnd:])
		if err != nil {
			return nil, 0, err
		}
		cdrList, err := parse(tokens[chunkEnd+1 : chunkEnd+chunk2End])
		if err != nil {
			return nil, 0, err
		}
		return mkListAsConsWithCdr(carList, cdrList[0]), chunkEnd + chunk2End + 1, nil
	}
	contents, err := parse(tokens[1:chunkEnd])
	if err != nil {
		return nil, 0, err
	}
	return mkListAsConsWithCdr(contents, Nil), chunkEnd + 1, nil
}

func lexAndParse(s string) ([]Sexpr, error) {
	return parse(lexItems(s))
}

func listChunk(tokens []lexutil.LexItem) (int, lexutil.LexItem, error) {
	level := 0
	for i, token := range tokens {
		switch token.Typ {
		case itemLeftParen:
			level++
		case itemRightParen:
			level--
			if level == 0 {
				return i, token, nil
			}
		case itemDot:
			if level == 1 {
				return i, token, nil
			}
		}
	}
	return 0, lexutil.LexItem{}, fmt.Errorf("unbalanced parens")
}

func dotChunk(tokens []lexutil.LexItem) (int, error) {
	level := 1
	for i, token := range tokens {
		switch token.Typ {
		case itemLeftParen:
			level++
		case itemRightParen:
			level--
			if level == 0 {
				return i, nil
			}
		}
	}
	return 0, fmt.Errorf("unbalanced parens")
}
