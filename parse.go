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
			var quoted Sexpr
			var delta int
			// quoted and delta depend on whether the quoted expression is
			// a list or an atom/num:
			if tokens[i].Typ != itemLeftParen {
				inner, err := parse(tokens[i : i+1])
				if err != nil {
					return nil, err
				}
				quoted = inner[0]
				delta = 1
			} else {
				start, end, err := balancedParenPoints(tokens[i:])
				if err != nil {
					return nil, err
				}
				inner, err := parse(tokens[i+start+1 : i+end])
				if err != nil {
					return nil, err
				}
				quoted = mkListAsCons(inner)
				delta = end - start + 1
			}
			i += delta
			quoteList := []Sexpr{Atom{"quote"}}
			quoteList = append(quoteList, quoted)
			ret = append(ret, mkListAsCons(quoteList))
		case itemLeftParen:
			start, end, err := balancedParenPoints(tokens[i:])
			if err != nil {
				return nil, err
			}

			// find '&' lexeme if it exists; ignore it
			for j := start; j < end; j++ {
				if tokens[i+j].Typ == itemDot {
					fmt.Println("found rest sep")
					end = j
					break
				}
			}
			// YAH / WIP
			inner, err := parse(tokens[i+start+1 : i+end])
			if err != nil {
				return nil, err
			}
			ret = append(ret, mkListAsCons(inner))
			i = i + end + 1
		case itemRightParen:
			return nil, fmt.Errorf("unexpected right paren")
		default:
			return nil, fmt.Errorf("unexpected lexeme '%s'", token.Val)
		}
	}
	return ret, nil
}

func lexAndParse(s string) ([]Sexpr, error) {
	return parse(lexItems(s))
}
