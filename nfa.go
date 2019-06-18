package regex_explained

// Thompson NFA Construction and Search

type state struct {
	isEnd              bool
	transition         map[byte]*state
	epsilonTransitions []*state
}

type nfa struct {
	start *state
	end   *state
}

func createState(isEnd bool) *state {
	return &state{
		isEnd:              isEnd,
		transition:         make(map[byte]*state),
		epsilonTransitions: nil,
	}
}

func addEpsilonTransition(from *state, to *state) {
	from.epsilonTransitions = append(from.epsilonTransitions, to)
}

func addTransition(from, to *state, symbol byte) {
	from.transition[symbol] = to
}

// Construct an NFA that recognizes only the empty string
func fromEpsilon() *nfa {
	start, end := createState(false), createState(true)
	addEpsilonTransition(start, end)
	return &nfa{start, end}
}

// Construct an NFA that recognizes only a single character string
func fromSymbol(symbol byte) *nfa {
	start, end := createState(false), createState(true)
	addTransition(start, end, symbol)
	return &nfa{start, end}
}

// concatenates two NFAs
func concat(first, second *nfa) *nfa {
	addEpsilonTransition(first.end, second.start)
	first.end.isEnd = false
	return &nfa{first.start, second.end}
}

// unions two NFAs
func union(first, second *nfa) *nfa {
	start, end := createState(false), createState(true)

	addEpsilonTransition(start, first.start)
	addEpsilonTransition(start, second.start)

	addEpsilonTransition(first.end, end)
	addEpsilonTransition(second.end, end)
	first.end.isEnd = false
	second.end.isEnd = false

	return &nfa{start, end}
}

func closure(n *nfa) *nfa {
	start, end := createState(false), createState(true)

	addEpsilonTransition(start, end)
	addEpsilonTransition(start, n.start)

	addEpsilonTransition(n.end, end)
	addEpsilonTransition(n.end, n.start)
	n.end.isEnd = false

	return &nfa{start, end}
}

func toNFA(postfixExp string) *nfa {
	if postfixExp == "" {
		return fromEpsilon()
	}

	var st []*nfa

	popStack := func() (top *nfa) {
		top, st = st[len(st)-1], st[:len(st)-1]
		return top
	}

	for i := 0; i < len(postfixExp); i++ {
		tok := postfixExp[i]

		var next *nfa
		switch tok {
		case STAR:
			next = closure(popStack())
		case OR:
			right := popStack()
			left := popStack()
			next = union(left, right)
		case CC:
			right := popStack()
			left := popStack()
			next = concat(left, right)
		default:
			next = fromSymbol(tok)
		}
		st = append(st, next)
	}

	return popStack()
}

func addNextState(s *state, ns []*state, visited map[*state]interface{}) (ret []*state) {
	if len(s.epsilonTransitions) > 0 {
		for _, es := range s.epsilonTransitions {
			if _, ok := visited[es]; !ok {
				visited[es] = struct{}{}
				ret = append(ret, addNextState(es, ns, visited)...)
			}
		}
	} else {
		ret = append(ret, s)
	}
	return
}

// breadth-first search
func bfs(n *nfa, word string) bool {
	var currStates []*state

	currStates = append(currStates, addNextState(n.start, currStates, make(map[*state]interface{}))...)

	for i := 0; i < len(word); i++ {
		symbol := word[i]
		var nextStates []*state

		for _, cs := range currStates {
			if nextState := cs.transition[symbol]; nextState != nil {
				nextStates = append(
					nextStates,
					addNextState(nextState, nextStates, make(map[*state]interface{}))...)
			}
		}

		currStates = nextStates
	}

	for _, cs := range currStates {
		if cs.isEnd {
			return true
		}
	}
	return false
}

// depth-first search
func dfs(s *state, visited map[*state]interface{}, exp string, pos int) bool {
	if _, ok := visited[s]; ok {
		return false
	}

	visited[s] = struct{}{}

	if pos == len(exp) {
		if s.isEnd {
			return true
		}

		for _, t := range s.epsilonTransitions {
			if dfs(t, visited, exp, pos) {
				return true
			}
		}
	} else {
		ns := s.transition[exp[pos]]

		if ns != nil {
			if dfs(ns, make(map[*state]interface{}), exp, pos+1) {
				return true
			}
		} else {
			for _, t := range s.epsilonTransitions {
				if dfs(t, visited, exp, pos) {
					return true
				}
			}
		}
	}
	return false
}

const (
	DFS = iota
	BFS
)

type MatchOptions struct {
	Method int
}

func MatchString(n *nfa, word string, options *MatchOptions) bool {
	if options.Method == DFS {
		return dfs(n.start, make(map[*state]interface{}), word, 0)
	} else {
		return bfs(n, word)
	}
}

func Compile(exp string) *nfa {
	return toNFA(toPostfix(insertExplicitConcatOperator(exp)))
}