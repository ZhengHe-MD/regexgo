package regexgo

import (
	"bytes"
	"regex-explained/token"
)

func insertExplicitConcatOperator(exp string) string {
	bs := bytes.NewBuffer([]byte{})

	for i := 0; i < len(exp); i++ {
		tok := exp[i]

		bs.WriteByte(tok)

		if tok == token.LP || tok == token.OR {
			continue
		}

		if i < len(exp)-1 {
			next := exp[i+1]

			if next == token.OR ||
				next == token.RP ||
				next == token.ZeroOrMore ||
				next == token.ZeroOrOne ||
				next == token.OneOrMore {
				continue
			}

			bs.WriteByte(token.CC)
		}
	}

	return bs.String()
}

// NOTE: The operator precedence, from weakest to strongest binding, is
// alternation -> concatenation -> repetition
var precedence = map[byte]int{
	token.OR:         0,
	token.CC:         1,
	token.ZeroOrMore: 2,
	token.ZeroOrOne:  3,
	token.OneOrMore:  4,
}

func toPostfix(exp string) string {
	bs := bytes.NewBuffer([]byte{})

	var st []byte

	peekStack := func() byte {
		return st[len(st)-1]
	}

	popStack := func() (top byte) {
		top, st = st[len(st)-1], st[:len(st)-1]
		return
	}

	takePrecedence := func(tok byte) {
		for len(st) > 0 && peekStack() != token.LP && precedence[peekStack()] >= precedence[tok] {
			bs.WriteByte(popStack())
		}
		st = append(st, tok)
	}

	for i := 0; i < len(exp); i++ {
		tok := exp[i]

		switch tok {
		case token.CC:
			takePrecedence(tok)
		case token.OR:
			takePrecedence(tok)
		case token.ZeroOrMore:
			takePrecedence(tok)
		case token.ZeroOrOne:
			takePrecedence(tok)
		case token.OneOrMore:
			takePrecedence(tok)
		case token.LP:
			st = append(st, tok)
		case token.RP:
			for len(st) > 0 && peekStack() != token.LP {
				bs.WriteByte(popStack())
			}
			popStack()
		default:
			bs.WriteByte(tok)
		}

	}

	for len(st) > 0 {
		bs.WriteByte(popStack())
	}

	return bs.String()
}
