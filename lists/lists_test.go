package lists

import (
	"fmt"
	"reflect"
	"testing"
)

func makeList(arr []int) *ListNode {
	var head, last *ListNode = nil, nil
	for _, v := range arr {
		var node = ListNode{
			Val:  v,
			Next: nil,
		}
		if head == nil {
			head = &node
			last = head
		} else {
			last.Next = &node
			last = last.Next
		}
	}
	return head
}

func Test_ReverseLinkedList(t *testing.T) {
	var ts = []struct {
		input, expected []int
	}{{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := listToSlice(reverseList(makeList(ts[i].input)))
			if !reflect.DeepEqual(result, ts[i].expected) {
				t.Errorf("input : %v exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_ReverseBetween(t *testing.T) {
	var ts = []struct {
		input    []int
		l, r     int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
		{[]int{3, 5}, 1, 2, []int{5, 3}},
	}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			var result []int
			ls := reverseBetween(makeList(ts[i].input), ts[i].l, ts[i].r)
			copy(result[:], listToSlice(ls))
			if reflect.DeepEqual(result, ts[i].expected) {
				t.Errorf("input : %v exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_IsPalindrome(t *testing.T) {
	var ts = []struct {
		input    []int
		expected bool
	}{{[]int{1, 2, 2, 1}, true}, {[]int{1, 2}, false}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := isPalindrome(makeList(ts[i].input))
			if result != ts[i].expected {
				t.Errorf("input : %v exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_DetectCycle(t *testing.T) {
	var ts = []struct {
		input    []int
		pos      int
		expected int
	}{{[]int{}, -1, -1},
		{[]int{3, 2, 0, -4}, 1, 2},
		{[]int{1}, -1, -1},
		{[]int{-1, -7, 7, -4, 19, 6, -9, -5, -2, -5}, 6, -9}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			ls := makeList(ts[i].input)
			if ts[i].expected != -1 {
				var link_to *ListNode = nil
				last := ls
				for cycle_pos := 0; last.Next != nil; cycle_pos, last = cycle_pos+1, last.Next {
					if cycle_pos == ts[i].pos {
						link_to = last
					}
				}
				last.Next = link_to
			}
			result := detectCycle(ls)
			switch ts[i].expected {
			case -1:
				if result != nil {
					t.Errorf("input : %v exp.: nil, got: %v", ts[i].input, result)
				}
			default:
				if result.Val != ts[i].expected {
					t.Errorf("input : %v exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
				}
			}

		})
	}
}
