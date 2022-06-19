package heaps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PQ_Peek(t *testing.T) {
	var ts = []struct {
		input []int
		exp   int
	}{{[]int{20, 22, 1, 5}, 22}, {[]int{15, 12, 50, 7, 40, 22}, 50}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			h := NewPriorityQueue()
			for _, v := range ts[i].input {
				h.Push(v)
			}
			v, _ := h.Peek()
			assert.Equal(t, v, ts[i].exp)
		})
	}
}

func Test_PQ_Pop(t *testing.T) {
	var ts = []struct {
		input []int
		exp   []int
	}{{[]int{20, 22, 1, 5}, []int{22, 20, 5, 1}}, {[]int{15, 12, 50, 7, 40, 22}, []int{50, 40, 22, 15, 12, 7}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			h := NewPriorityQueue()
			for _, v := range ts[i].input {
				h.Push(v)
			}
			result := make([]int, 0)
			for v, ok := h.Pop(); ok; v, ok = h.Pop() {
				result = append(result, v)
			}
			assert.Equal(t, result, ts[i].exp)
		})
	}

}
