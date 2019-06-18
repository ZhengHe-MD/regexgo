package regexgo

import (
	"testing"
)

func TestInsertExplicitConcatOperator(t *testing.T) {
	cases := []struct{
		exp string
		expected string
	} {
		{"abc", "a.b.c"},
		{"(a|b)c", "(a|b).c"},
		{"a*bc", "a*.b.c"},
		{"a*(b|c)d", "a*.(b|c).d"},
		{"a*b|c", "a*.b|c"},
		{"(a|b)*c", "(a|b)*.c"},
	}

	for i, c := range cases {
		out := insertExplicitConcatOperator(c.exp)
		if out != c.expected {
			t.Errorf("Test:%d input:%v got:%v expect:%v",i, c.exp, out, c.expected)
		}
	}
}

func TestToPostfix(t *testing.T) {
	cases := []struct{
		exp string
		expected string
	} {
		{"a.b.c", "ab.c."},
		{"(a|b).c", "ab|c."},
		{"a*.b.c", "a*b.c."},
		{"a*.(b|c).d", "a*bc|.d."},
		{"a*.b|c", "a*b.c|"},
		{"(a|b)*.c", "ab|*c."},
	}

	for i, c := range cases {
		out := toPostfix(c.exp)
		if out != c.expected {
			t.Errorf("Test:%d input:%v got:%v expect:%v", i, c.exp, out, c.expected)
		}
	}
}