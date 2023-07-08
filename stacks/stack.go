package stacks

import (
	"unicode"
)

// https://leetcode.com/problems/valid-parentheses/
// Valid Parentheses
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	pars := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	stack := make([]byte, 0)
	for _, c := range s {
		if v, ok := pars[byte(c)]; ok {
			stack = append(stack, byte(v)) // push
		} else {
			if len(stack) > 0 {
				var x byte
				if x, stack = stack[len(stack)-1], stack[:len(stack)-1]; x != byte(c) { //  pop
					return false
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/
// Minimum Remove to Make Valid Parentheses
func minRemoveToMakeValid(s string) string {
	buf := append([]byte(nil), s...)
	stack := make([]int, 0)
	for i, c := range s {
		switch c {
		case '(':
			stack = append(stack, i)
		case ')':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				buf[i] = ' '
			}
		}
	}

	for _, pos := range stack {
		buf[pos] = ' '
	}
	result := make([]byte, 0)
	for i := range buf {
		if !unicode.IsSpace(rune(buf[i])) {
			result = append(result, buf[i])
		}
	}
	return string(result)
}

// https://leetcode.com/problems/implement-queue-using-stacks/
//
//	Implement Queue using Stacks
type MyQueue struct {
	stack1, stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{stack1: []int{}, stack2: []int{}}
}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	var e int
	if len(this.stack2) == 0 {
		for len(this.stack1) > 0 {
			e, this.stack1 = pop(this.stack1)
			this.stack2 = append(this.stack2, e)
		}
	}
	e, this.stack2 = pop(this.stack2)
	return e
}

func (this *MyQueue) Peek() int {
	if len(this.stack2) == 0 {
		for len(this.stack1) > 0 {
			var e int
			e, this.stack1 = this.stack1[len(this.stack1)-1], this.stack1[:len(this.stack1)-1]
			this.stack2 = append(this.stack2, e)
		}
	}
	return this.stack2[len(this.stack2)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1)+len(this.stack2) == 0
}

func pop(stack []int) (e int, s []int) {
	e, s = stack[len(stack)-1], stack[:len(stack)-1]
	return
}
