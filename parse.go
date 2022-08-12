package main

import (
	"fmt"

	"github.com/eigenhombre/lexutil"
)

// parse returns a list of sexprs parsed from a list of tokens.
func parse(tokens []lexutil.LexItem) ([]Sexpr, error) {
	ret := []Sexpr{}
	i := 0
	for {
		if i >= len(tokens) {
			break
		}
		token := tokens[i]
		switch token.Typ {
		case itemNumber:
			ret = append(ret, Num(token.Val))
			i++
		case itemAtom:
			ret = append(ret, Atom{token.Val})
			i++
		case itemForwardQuote:
			i++ // skip ' token
			if i >= len(tokens) {
				return nil, fmt.Errorf("unexpected end of input")
			}
			if tokens[i].Typ != itemLeftParen {
				inner, err := parse(tokens[i : i+1])
				if err != nil {
					return nil, err
				}
				ret = append(ret, Cons(Atom{"quote"}, Cons(inner[0], Nil)))
				i++
			} else {
				inner, incr, err := parseList(tokens[i:])
				if err != nil {
					return nil, err
				}
				ret = append(ret, Cons(Atom{"quote"}, Cons(inner, Nil)))
				i += incr
			}
		case itemLeftParen:
			item, incr, err := parseList(tokens[i:])
			if err != nil {
				return nil, err
			}
			ret = append(ret, item)
			i += incr
		case itemRightParen:
			return nil, fmt.Errorf("unexpected right paren")
		default:
			return nil, fmt.Errorf("unexpected lexeme '%s'", token.Val)
		}
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
