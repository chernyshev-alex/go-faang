package stacks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsValid_Parentheses(t *testing.T) {
	var ts = []struct {
		input    string
		expected bool
	}{{"()", true}, {"()[]{}", true}, {"(]", false}, {"]", false}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := isValid(ts[i].input)
			if result != ts[i].expected {
				t.Errorf("input %v, exp %v, got %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_MinRemoveToMakeValid(t *testing.T) {
	var ts = []struct {
		input    string
		expected string
	}{{"lee(t(c)o)de)", "lee(t(c)o)de"}, {"a)b(c)d", "ab(c)d"}, {"))((", ""}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := minRemoveToMakeValid(ts[i].input)
			if result != ts[i].expected {
				t.Errorf("input %v, exp %v, got %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}
func Test_ImplQueue(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)

	assert.Equal(t, 1, q.Peek())
	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, false, q.Empty())
}
