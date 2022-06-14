package search

import (
	"fmt"
	"testing"
)

func Test_Find_Kth_Largest_Qsort(t *testing.T) {
	var ts = []struct {
		input    []int
		k        int
		expected int
	}{{[]int{3, 2, 1, 5, 6, 4}, 2, 5}, {[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := FindKthLargest(ts[i].input, ts[i].k)
			if result != ts[i].expected {
				t.Errorf("input %v, exp %v, got %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_Find_Kth_Largest_Qselect(t *testing.T) {
	var ts = []struct {
		input    []int
		k        int
		expected int
	}{{[]int{3, 2, 1, 5, 6, 4}, 2, 5}, {[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := FindKthLargestQselect(ts[i].input, ts[i].k)
			if result != ts[i].expected {
				t.Errorf("input %v, exp %v, got %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_SearchRange(t *testing.T) {
	var ts = []struct {
		input    []int
		k        int
		expected [2]int
	}{{[]int{}, 0, [2]int{-1, -1}},
		{[]int{5, 7, 7, 8, 8, 10}, 8, [2]int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, [2]int{-1, -1}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			var result [2]int
			copy(result[:], SearchRange(ts[i].input, ts[i].k))
			if result != ts[i].expected {
				t.Errorf("input %v; %v, exp %v, got %v", ts[i].input, ts[i].k, ts[i].expected, result)
			}
		})
	}
}
