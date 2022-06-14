package trees

import (
	"fmt"
	"reflect"
	"testing"
)

func buildTree(v []int) *TreeNode {
	if len(v) == 0 {
		return nil
	}
	root := &TreeNode{Val: v[0]}

	queue := append([]*TreeNode(nil), root)
	for i := 1; i < len(v) && len(queue) > 0; {
		var tn *TreeNode
		tn, queue = queue[0], queue[1:] // pop&shift
		for edgeId := range []byte{0, 1} {
			nodePtr := &tn.Left
			if edgeId == 1 {
				nodePtr = &tn.Right
			}
			if v[i] != -1 {
				*nodePtr = &TreeNode{Val: v[i]}
			}
			if i++; i >= len(v) {
				return root
			}
			if *nodePtr != nil {
				queue = append(queue, *nodePtr)
			}
		}
	}
	return root
}

func Test_MaxDepth(t *testing.T) {
	var ts = []struct {
		input    []int
		expected int
	}{{[]int{3, 9, 20, -1, -1, 15, 7}, 3}, {[]int{1, -1, 2}, 2}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			tree := buildTree(ts[i].input)
			result := maxDepth(tree)
			if result != ts[i].expected {
				t.Errorf("input %v, exp %v, got %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_LevelOrder(t *testing.T) {
	var ts = []struct {
		input    []int
		expected [][]int
	}{{[]int{3, 9, 20, -1, -1, 15, 7}, [][]int{{3}, {9, 20}, {15, 7}}},
		{[]int{1}, [][]int{{1}}}, {[]int{}, [][]int{{}}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			tree := buildTree(ts[i].input)
			result := levelOrder(tree)
			if tree != nil && !reflect.DeepEqual(result, ts[i].expected) {
				t.Errorf("input %v, exp %v, got %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}
