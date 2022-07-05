package strs

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/backspace-string-compare/
//
// Given two strings s and t, return true if they are equal when both are typed into empty text editors.
// '#' means a backspace character.

func Test_BackspaceCompare(t *testing.T) {
	var ts = []struct {
		s1, s2   string
		expected bool
	}{{"ab#c", "ad#c", true}, {"a##c", "#a#c", true}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := backSpaceCompare(ts[i].s1, ts[i].s2)
			if result != ts[i].expected {
				t.Errorf("input : %v %v; exp.: %v, got: %v", ts[i].s1, ts[i].s2, ts[i].expected, result)
			}
		})
	}
}

func Test_BackspaceCompareRev(t *testing.T) {
	var ts = []struct {
		s1, s2   string
		expected bool
	}{{"ab#c", "ad#c", true}, {"xywrrmp", "xywrrmu#p", true}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := backSpaceCompareRev(ts[i].s1, ts[i].s2)
			if result != ts[i].expected {
				t.Errorf("input : %v %v; exp.: %v, got: %v", ts[i].s1, ts[i].s2, ts[i].expected, result)
			}
		})
	}
}

// Reverse string

func Test_Reverse(t *testing.T) {
	var ts = []struct {
		input, expected string
	}{{"123456", "654321"}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := reverse(ts[i].input)
			if result != ts[i].expected {
				t.Errorf("input : %v exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

// https://leetcode.com/problems/valid-palindrome/

func Test_validPalindrome(t *testing.T) {
	var ts = []struct {
		input    string
		expected bool
	}{{"atbbga", false}, {"aba", true}, {"abca", true}, {"abc", false}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := validPalindrome(ts[i].input)
			if result != ts[i].expected {
				t.Errorf("input : %v exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_IsPalindrome(t *testing.T) {
	var ts = []struct {
		input    string
		expected bool
	}{{"A man, a plan, a canal: Panama", true}, {"race a car", false}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := isPalindrome(ts[i].input)
			if result != ts[i].expected {
				t.Errorf("input : %v exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

//
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
//
// Given a string s, find the length of the longest substring without repeating characters.
// Sliding window
//
func Test_LengthOfLongestSubstring(t *testing.T) {
	var ts = []struct {
		input    string
		expected int
	}{{"bbbbb", 1}, {"ab", 2}, {"abcabcbb", 3}, {"pwwkew", 3}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := lengthOfLongestSubstring(ts[i].input)
			if result != ts[i].expected {
				t.Errorf("input : %v exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

// Longest Palindromic Substring
// Given a string, find the longest substring which is a palindrome.

func Test_LongestPalindromicSubstring(t *testing.T) {
	var ts = []struct {
		input    string
		expected string
	}{{"Geeks", "ee1"}, {"forgeeksskeegfor", "geeksskeeg"}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := longestPalindromicSubstring(ts[i].input)
			if result != ts[i].expected {
				t.Errorf("input : %v exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}
