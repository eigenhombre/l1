package main

import (
	"fmt"
)

func handleQuoteItem(tokens []token, i int, operatorName string) (Sexpr, int, error) {
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

func parseNext(tokens []token, i int) (Sexpr, int, error) {
	if i >= len(tokens) {
		return nil, 0, fmt.Errorf("unexpected end of input; index=%d, tokens=%v", i, tokens)
	}
	token := tokens[i]
	switch token.lexeme.Typ {
	case itemNumber:
		return Num(token.lexeme.Val), 1, nil
	case itemAtom:
		return Atom{token.lexeme.Val}, 1, nil
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
		return nil, 0, fmt.Errorf("unexpected right paren on line %d", token.line)
	default:
		return nil, 0, fmt.Errorf("unexpected lexeme '%s' on line %d", token.lexeme.Val, token.line)
	}
}

func parse(tokens []token) ([]Sexpr, error) {
	ret := []Sexpr{}
	i := 0
	// Look for shebang, only at beginning of file:
	if len(tokens) > 0 && tokens[0].lexeme.Typ == itemShebang {
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
func parseList(tokens []token) (Sexpr, int, error) {
	chunkEnd, endTok, err := listChunk(tokens)
	if err != nil {
		return nil, 0, err
	}
	if endTok.lexeme.Typ == itemDot {
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

func lexAndParse(ss []string) ([]Sexpr, error) {
	return parse(lexItems(ss))
}

func listChunk(tokens []token) (int, token, error) {
	level := 0
	for i, token := range tokens {
		switch token.lexeme.Typ {
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
	return 0, token{}, fmt.Errorf("unbalanced parens")
}

func dotChunk(tokens []token) (int, error) {
	level := 1
	for i, token := range tokens {
		switch token.lexeme.Typ {
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
