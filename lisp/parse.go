package lisp

func handleQuoteItem(tokens []Token, i int, operatorName string) (Sexpr, int, error) {
	if i >= len(tokens) {
		return nil, 0, baseErrorf("unexpected end of input; index=%d, tokens=%v", i, tokens)
	}
	nextParsed, incr, err := parseNext(tokens, i)
	if err != nil {
		return nil, 0, extendError("handleQuoteItem parseNext", err)
	}
	item := Cons(Atom{operatorName}, Cons(nextParsed, Nil))
	return item, incr, nil
}

func parseNext(tokens []Token, i int) (Sexpr, int, error) {
	if i >= len(tokens) {
		return nil, 0, baseErrorf("unexpected end of input; index=%d, tokens=%v", i, tokens)
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
			return nil, 0, extendError("parseNext itemForwardQuote handleQuoteItem", err)
		}
		return item, incr + 1, nil
	case itemSyntaxQuote:
		item, incr, err := handleQuoteItem(tokens, i+1, "syntax-quote")
		if err != nil {
			return nil, 0, extendError("parseNext itemSyntaxQuote handleQuoteItem", err)
		}
		return item, incr + 1, nil
	case itemUnquote:
		item, incr, err := handleQuoteItem(tokens, i+1, "unquote")
		if err != nil {
			return nil, 0, extendError("parseNext itemUnquote handleQuoteItem", err)
		}
		return item, incr + 1, nil
	case itemSplicingUnquote:
		item, incr, err := handleQuoteItem(tokens, i+1, "splicing-unquote")
		if err != nil {
			return nil, 0, extendError("parseNext itemSplicingUnquote handleQuoteItem", err)
		}
		return item, incr + 1, nil
	case itemCommentNext:
		item, incr, err := handleQuoteItem(tokens, i+1, "comment")
		if err != nil {
			return nil, 0, extendError("parseNext itemCommentNext handleQuoteItem", err)
		}
		return item, incr + 1, nil
	case itemLeftParen:
		item, incr, err := parseList(tokens[i:])
		if err != nil {
			return nil, 0, extendError("parseNext itemLeftParen parseList", err)
		}
		return item, incr, nil
	case itemRightParen:
		return nil, 0, baseErrorf("unexpected right paren on line %d", token.line)
	default:
		return nil, 0, baseErrorf("unexpected lexeme '%s' on line %d", token.lexeme.Val, token.line)
	}
}

// Parse takes a slice of tokens and returns a slice of Sexprs.
func Parse(tokens []Token) ([]Sexpr, error) {
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
			return nil, extendError("parse parseNext", err)
		}
		ret = append(ret, item)
		i += incr
	}
	return ret, nil
}

// parseList is used when a list has been detected in a slice of tokens.
func parseList(tokens []Token) (Sexpr, int, error) {
	chunkEnd, endTok, err := listChunk(tokens)
	if err != nil {
		return nil, 0, err
	}
	if endTok.lexeme.Typ == itemDot {
		carList, err := Parse(tokens[1:chunkEnd])
		if err != nil {
			return nil, 0, err
		}
		chunk2End, err := dotChunk(tokens[chunkEnd:])
		if err != nil {
			return nil, 0, err
		}
		cdrList, err := Parse(tokens[chunkEnd+1 : chunkEnd+chunk2End])
		if err != nil {
			return nil, 0, err
		}
		return mkListAsConsWithCdr(carList, cdrList[0]), chunkEnd + chunk2End + 1, nil
	}
	contents, err := Parse(tokens[1:chunkEnd])
	if err != nil {
		return nil, 0, err
	}
	return mkListAsConsWithCdr(contents, Nil), chunkEnd + 1, nil
}

func lexAndParse(ss []string) ([]Sexpr, error) {
	return Parse(LexItems(ss))
}

func listChunk(tokens []Token) (int, Token, error) {
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
	return 0, Token{}, baseError("unbalanced parens")
}

func dotChunk(tokens []Token) (int, error) {
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
	return 0, baseError("unbalanced parens")
}
