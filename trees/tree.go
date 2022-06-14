package trees

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
