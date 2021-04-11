package regexgo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchString(t *testing.T) {
	tests := map[string]struct {
		givenExp   string
		givenWords []string
		wantRets   []bool
	}{
		"empty string": {
			givenExp:   "",
			givenWords: []string{"", "a", "ab", "abc"},
			wantRets:   []bool{true, false, false, false},
		},
		"single char": {
			givenExp:   "a",
			givenWords: []string{"bcd", "a", "", "xkcd", "abc", "aaa"},
			wantRets:   []bool{false, true, false, false, false, false},
		},
		"closure": {
			givenExp:   "a*",
			givenWords: []string{"", "a", "aa", "aaa", "aaaa", "b"},
			wantRets:   []bool{true, true, true, true, true, false},
		},
		"concatenation of two chars": {
			givenExp:   "ab",
			givenWords: []string{"a", "b", "ab", "abc"},
			wantRets:   []bool{false, false, true, false},
		},
		"union of two chars": {
			givenExp:   "a|b",
			givenWords: []string{"a", "b", "ab", "bb"},
			wantRets:   []bool{true, true, false, false},
		},
		"mixed case 1": {
			givenExp:   "(a|b)*c",
			givenWords: []string{"ac", "abc", "aabababbc", "aaaab"},
			wantRets:   []bool{true, true, true, false},
		},
		"regex for all binary numbers divisible by 3": {
			givenExp:   "(0|(1(01*(00)*0)*1)*)*",
			givenWords: []string{"", "0", "00", "01", "10", "11", "000", "011", "110", "0000", "0011"},
			wantRets:   []bool{true, true, true, false, false, true, true, true, true, true, true},
		},
	}

	for name, tc := range tests {
		n := Compile(tc.givenExp)
		t.Run(name, func(t *testing.T) {
			for i := 0; i < len(tc.givenWords); i++ {
				assert.Equal(t, tc.wantRets[i], MatchString(n, tc.givenWords[i], &MatchOptions{DFS}),
					fmt.Sprintf("exp: %s word: %s want: %v\n", tc.givenExp, tc.givenWords[i], tc.wantRets[i]))
				assert.Equal(t, tc.wantRets[i], MatchString(n, tc.givenWords[i], &MatchOptions{BFS}),
					fmt.Sprintf("exp: %s word: %s want: %v\n", tc.givenExp, tc.givenWords[i], tc.wantRets[i]))
			}
		})
	}
}
