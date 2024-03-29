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

func Test_RightSideViewBfs(t *testing.T) {
	var ts = []struct {
		input    []int
		expected []int
	}{{[]int{1, 2, 3, -1, 5, -1, 4}, []int{1, 3, 4}},
		{[]int{1, -1, 3}, []int{1, 3}},
		{[]int{}, []int{}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			tree := buildTree(ts[i].input)
			result := rightSideViewBfs(tree)
			if tree != nil && !reflect.DeepEqual(result, ts[i].expected) {
				t.Errorf("input %v, exp %v, got %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_RightSideViewDfs(t *testing.T) {
	var ts = []struct {
		input    []int
		expected []int
	}{{[]int{1, 2, 3, -1, 5, -1, 4}, []int{1, 3, 4}},
		{[]int{1, -1, 3}, []int{1, 3}},
		{[]int{}, []int{}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			tree := buildTree(ts[i].input)
			result := rightSideViewDfs(tree)
			if tree != nil && !reflect.DeepEqual(result, ts[i].expected) {
				t.Errorf("input %v, exp %v, got %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Benchmark_RightSideViewDfs(b *testing.B) {
	var ts = []struct {
		input    []int
		expected []int
	}{{[]int{1, 2, 3, -1, 5, -1, 4}, []int{1, 3, 4}}}

	tree := buildTree(ts[0].input)
	for n := 0; n < b.N; n++ {
		rightSideViewDfs(tree)
	}
}

func Test_CountNodesRecursive(t *testing.T) {
	var ts = []struct {
		input    []int
		expected int
	}{{[]int{1, 2, 3, 4, 5, 6}, 6}, {[]int{1}, 1}, {[]int{}, 0}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			tree := buildTree(ts[i].input)
			result := countNodesRecursive(tree)
			if result != ts[i].expected {
				t.Errorf("input %v, exp %v, got %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_CountNodesOpt(t *testing.T) {
	var ts = []struct {
		input    []int
		expected int
	}{{[]int{1, 2, 3}, 3}, {[]int{1, 2, 3, 4, 5, 6}, 6}, {[]int{}, 0}, {[]int{1}, 1}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			tree := buildTree(ts[i].input)
			result := countNodesOpt(tree)
			if result != ts[i].expected {
				t.Errorf("input %v, exp %v, got %v\n", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_IsValidBST(t *testing.T) {
	var ts = []struct {
		input    []int
		expected bool
	}{{[]int{2, 1, 3}, true}, {[]int{5, 1, 4, -1, -1, 3, 6}, false}, {[]int{2147483647}, true}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			tree := buildTree(ts[i].input)
			result := isValidBST(tree)
			if result != ts[i].expected {
				t.Errorf("input %v, exp %v, got %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_InOrderSuccessor(t *testing.T) {
	var ts = []struct {
		input     []int
		start     int
		successor int
	}{{[]int{20, 8, 22, 4, 12, 10, 14}, 8, 10},
		{[]int{20, 8, 22, 4, 12, 10, 14}, 10, 12},
		{[]int{20, 8, 22, 4, 12, 10, 14}, 14, 20}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			tree := new(BinaryTree)
			for _, n := range ts[i].input {
				tree.Insert(n)
			}
			result := tree.InOrderSuccessor(ts[i].start)
			if result != ts[i].successor {
				t.Errorf("input %v, start %d, exp %v, got %v", ts[i].input, ts[i].start,
					ts[i].successor, result)
			}
		})
	}
}

// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/description/
func Test_DistanceK(t *testing.T) {
	var ts = []struct {
		input     []int
		target, k int
		expected  []int
	}{{[]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, 5, 2, []int{7, 4, 1}},
		{[]int{1}, 1, 3, []int{}},
	}
	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			root := CreateTreeFromSlice(ts[i].input)
			target := FindNodeInTree(root, ts[i].target)

			result := distanceK(root, target, ts[i].k)
			if !reflect.DeepEqual(result, ts[i].expected) {
				t.Errorf("input %v, target %d, k %d, expected %v, got %v",
					ts[i].input, ts[i].target, ts[i].k, ts[i].expected, result)
			}
		})
	}

}
