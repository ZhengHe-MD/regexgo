package regexgo

import (
	"testing"
)

func TestMatchString(t *testing.T) {
	cases := []struct {
		exp          string
		words        []string
		expectedRets []bool
	}{
		{
			"(a|b)*c",
			[]string{"ac", "abc", "aabababbc", "aaaab"},
			[]bool{true, true, true, false},
		},
		{
			"ab",
			[]string{"a", "b", "ab", "abc"},
			[]bool{false, false, true, false},
		},
		{
			"a*b",
			[]string{"b", "ab", "aab", "aba", "aaaaaaab"},
			[]bool{true, true, true, false, true},
		},
		{
			"a*|b*",
			[]string{"a", "b", "aaaa", "bbbb", "abab", "aaabbb"},
			[]bool{true, true, true, true, false, false},
		},
	}

	for i, c := range cases {
		n := Compile(c.exp)
		for j := 0; j < len(c.words); j++ {
			if MatchString(n, c.words[j], &MatchOptions{DFS}) != c.expectedRets[j] {
				t.Errorf("Test:%d regexp:%s word:%s method:%s ret:%v", i, c.exp, c.words[j], "DFS", c.expectedRets[j])
			}

			if MatchString(n, c.words[j], &MatchOptions{BFS}) != c.expectedRets[j] {
				t.Errorf("Test:%d regexp:%s word:%s method:%s ret:%v", i, c.exp, c.words[j], "BFS", c.expectedRets[j])
			}
		}
	}
}
