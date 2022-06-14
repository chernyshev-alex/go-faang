package trees

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
//

func __maxDepth(n *TreeNode, depth int) int {
	if n == nil {
		return depth
	}
	return max(__maxDepth(n.Left, depth+1), __maxDepth(n.Right, depth+1))
}

func maxDepth(root *TreeNode) int {
	return __maxDepth(root, 0)
}

// https://leetcode.com/problems/binary-tree-level-order-traversal/
// Given the root of a binary tree, return the level order traversal of its nodes' values.
// (i.e., from left to right, level by level).
//
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	q := append([]*TreeNode(nil), root)
	for len(q) > 0 {
		level, next_level := make([]int, 0), make([]*TreeNode, 0)
		for len(q) > 0 {
			var n *TreeNode
			n, q = q[0], q[1:]
			level = append(level, n.Val)
			for _, nn := range []*TreeNode{n.Left, n.Right} {
				if nn != nil {
					next_level = append(next_level, nn)
				}
			}
		}
		result = append(result, level)
		q = append(q, next_level...)
	}
	return result
}

// https://leetcode.com/problems/binary-tree-right-side-view/
// Given the root of a binary tree, imagine yourself standing on the right side of it,
// return the values of the nodes you can see ordered from top to bottom.
//
func rightSideViewBfs(root *TreeNode) []int {
	result, q := make([]int, 0), make([]*TreeNode, 0)
	if root != nil {
		q = append(q, root)
	}

	var n *TreeNode
	for len(q) > 0 {
		for cnt, qlen := 0, len(q); cnt < qlen; cnt++ {
			n, q = q[0], q[1:]
			for _, leaf := range []*TreeNode{n.Left, n.Right} {
				if leaf != nil {
					q = append(q, leaf)
				}
			}
		}
		result = append(result, n.Val)
	}
	return result
}

func rightSideDfs(n *TreeNode, level int, v []int) []int {
	if n != nil {
		if level >= len(v) {
			v = append(v, n.Val)
		}
		if n.Right != nil {
			return rightSideDfs(n.Right, level+1, v)
		}
		if n.Left != nil {
			return rightSideDfs(n.Left, level+1, v)
		}
	}
	return v
}

func rightSideViewDfs(root *TreeNode) []int {
	return rightSideDfs(root, 0, make([]int, 0))
}

// https://leetcode.com/problems/count-complete-tree-nodes/
// Given the root of a **complete** binary tree, return the number of the nodes in the tree.

func __countNodeRecursive(n *TreeNode) int {
	if n == nil {
		return 0
	}
	return 1 + __countNodeRecursive(n.Left) + __countNodeRecursive(n.Right)
}

func countNodesRecursive(root *TreeNode) int {
	return __countNodeRecursive(root)
}

// -- O(n) ---------------
//
func tree_height(n *TreeNode) (height int) {
	for height, curr := 0, n; curr != nil; curr, height = curr.Left, height+1 {
	}
	return
}

func node_exists(idx int, h int, node *TreeNode) bool {
	for l, r, count := 0, int(math.Pow(float64(h), 2)-1), 0; count < h; count++ {
		mid := int(math.Ceil(float64((l + r)) / 2.0))
		if node != nil {
			if idx >= mid {
				l, node = mid, node.Right
			} else {
				r, node = mid-1, node.Left
			}
		}
	}
	return node != nil
}

func countNodesOpt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	h := tree_height(root)
	if h == 0 {
		return 1
	}

	var upper_count int = int(math.Pow(float64(h), 2)) - 1
	l, r := 0, upper_count
	for l < r {
		var idx_find int = int(math.Ceil(float64(l+r) / 2))
		if node_exists(idx_find, h, root) {
			l = idx_find
		} else {
			r = idx_find - 1
		}
	}
	return upper_count + l + 1
}

// https://leetcode.com/problems/validate-binary-search-tree/
// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
// A valid BST is defined as follows:
//
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.
//
func dfs(n *TreeNode, min int, max int) bool {
	if n == nil {
		return true
	}
	if n.Val <= min || n.Val >= max {
		return false
	}
	return dfs(n.Left, min, n.Val) && dfs(n.Right, n.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt64, math.MaxInt64)
}
