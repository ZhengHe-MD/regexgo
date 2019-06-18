package regexgo

import (
	"bytes"
)

func insertExplicitConcatOperator(exp string) string {
	bs := bytes.NewBuffer([]byte{})

	for i := 0; i < len(exp); i++ {
		tok := exp[i]

		bs.WriteByte(tok)

		if tok == LP || tok == OR {
			continue
		}

		if i < len(exp)-1 {
			next := exp[i+1]

			if next == STAR || next == RP || next == OR {
				continue
			}

			bs.WriteByte(CC)
		}
	}

	return bs.String()
}

var precedence = map[byte]int{
	OR:   0,
	CC:   1,
	STAR: 2,
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
		for len(st) > 0 && peekStack() != LP && precedence[peekStack()] >= precedence[tok] {
			bs.WriteByte(popStack())
		}
		st = append(st, tok)
	}

	for i := 0; i < len(exp); i++ {
		tok := exp[i]

		switch tok {
		case CC:
			takePrecedence(tok)
		case OR:
			takePrecedence(tok)
		case STAR:
			takePrecedence(tok)
		case LP:
			st = append(st, tok)
		case RP:
			for len(st) > 0 && peekStack() != LP {
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
