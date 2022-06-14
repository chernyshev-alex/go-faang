package lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func listToSlice(ls *ListNode) []int {
	var result = make([]int, 0, 500)
	for ; ls != nil; ls = ls.Next {
		result = append(result, ls.Val)
	}
	return result
}

// https://leetcode.com/problems/reverse-linked-list/
//
// Given the head of a singly linked list, reverse the list, and return the reversed list.
//

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return h
}

// https://leetcode.com/problems/reverse-linked-list-ii/
//
// Given the head of a singly linked list and two integers left and right
// where left <= right, reverse the nodes of the list from position left to position right,
// and return the reversed list.
//
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 2.134.5 -> 2.314.5  -> 2.431.5
	pos := 1
	ls := &ListNode{Next: head}
	for n := ls; n != nil; n = n.Next {
		if pos == left {
			ln := n // anchor
			for cn := ln.Next; pos >= left && pos < right; {
				rn := cn.Next
				_t := ln.Next

				ln.Next = rn
				cn.Next = rn.Next
				rn.Next = _t
				//printList("h", ls.Next)
				pos++
			}
			break
		}
		pos++
	}
	return ls.Next
}

// https://leetcode.com/problems/palindrome-linked-list/
// Given the head of a singly linked list, return true if it is a palindrome.
//
func isPalindrome(head *ListNode) bool {
	sz := 0
	for n := head; n != nil; n, sz = n.Next, sz+1 {
	}

	mid := sz / 2
	if sz%2 != 0 {
		mid = (sz - 1) / 2
	}

	curr := head
	stack := make([]int, mid)
	for m := mid; m > 0; m, curr = m-1, curr.Next {
		stack = append(stack, curr.Val)
	}

	if sz%2 != 0 {
		curr = curr.Next
	}

	for ; curr != nil; curr = curr.Next {
		n := len(stack) - 1
		top := stack[n]
		if top != curr.Val {
			return false
		}
		stack = stack[:n]
	}
	return true
}

// https://leetcode.com/problems/linked-list-cycle-ii/
//Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null
//Floyd's tortoise & hare algorithm
//
func detectCycle(head *ListNode) *ListNode {
	// find cycle
	t, h := head, head
	cycleFound := false
	for !cycleFound && h != nil && h.Next != nil {
		if t, h = t.Next, h.Next.Next; t == h {
			cycleFound = true
		}
	}
	if !cycleFound {
		return nil
	}
	// find a rendezvous node = cycle
	for t = head; t != h; t, h = t.Next, h.Next {
	}
	return t
}
