package arr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	var ts = []struct {
		input    []int
		target   int
		expected [2]int
	}{{[]int{2, 7, 11, 15}, 9, [2]int{0, 1}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			var result [2]int
			copy(result[:], twoSum(ts[i].input, ts[i].target))
			if result != ts[i].expected {
				t.Errorf("input : %v; exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_maxArea(t *testing.T) {
	var ts = []struct {
		input    []int
		expected int
	}{{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := maxArea(ts[i].input)
			if result != ts[i].expected {
				t.Errorf("input : %v; exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_trap(t *testing.T) {
	var ts = []struct {
		input    []int
		expected int
	}{{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := trap(ts[i].input)
			if result != ts[i].expected {
				t.Errorf("input : %v; exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_fibo(t *testing.T) {
	t.Run("fibo", func(t *testing.T) {
		assert.Equal(t, []int64{1, 2, 3, 5, 8, 13, 21}, fibo_gen(6))
	})
}
