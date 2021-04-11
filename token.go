package regexgo

const (
	// operators
	OR   = '|' // alternation
	CC   = '.' // concatenation
	STAR = '*' // repetition
	QM   = '?' // repetition
	LP   = '('
	RP   = ')'
)

var operatorSet = map[byte]bool{
	OR:   true,
	CC:   true,
	STAR: true,
	QM:   true,
	LP:   true,
	RP:   true,
}

func isReserved(c byte) bool {
	return operatorSet[c]
}
