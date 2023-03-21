package strs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
//
// Given a string s, find the length of the longest substring without repeating characters.
// Sliding window
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
	}{{"Geeks", "ee"}} // , {"forgeeksskeegfor", "geeksskeeg"}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := longestPalindromicSubstring(ts[i].input)
			if result != ts[i].expected {
				t.Errorf("input : %v exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

// find all indexes of subseq in arr
// beware of case: 'arr="aaaaa", subseq="aaa"' -> Se	q(0, 1, 2)
func Test_SubseqIndexes(t *testing.T) {
	var ts = []struct {
		str, substr string
		expected    []int
	}{{"000AAAAA1111", "AA", []int{3, 4, 5, 6}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			assert.Equal(t, ts[i].expected, subseqIndexesIter(ts[i].str, ts[i].substr))

		})
	}
}

func Test_SubseqIndexesReq(t *testing.T) {
	var ts = []struct {
		str, substr string
		expected    []int
	}{{"012AAAAA3456", "AA", []int{3, 4, 5, 6}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			assert.Equal(t, ts[i].expected, subseqIndexesReq(ts[i].str, ts[i].substr))

		})
	}
}

// non overlapping
func Test_SubseqIndexesRegexp(t *testing.T) {
	var ts = []struct {
		str, substr string
		expected    [][]int
	}{{"012AAAAA3456", "AA", [][]int{{3, 5}, {5, 7}}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			assert.Equal(t, ts[i].expected, subseqIndexesRegExp(ts[i].str, ts[i].substr))

		})
	}
}
func Test_AllSubsets(t *testing.T) {
	var ts = []struct {
		input    []byte
		expected [][]byte
	}{{[]byte{'a', 'b', 'c'}, [][]byte{{}, {'a'}, {'b'}, {'a', 'b'}, {'c'}, {'a', 'c'}, {'b', 'c'}, {'a', 'b', 'c'}}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			assert.Equal(t, ts[i].expected, allSubsets(ts[i].input))
		})
	}
}

func Test_PermutationsRec(t *testing.T) {
	var ts = []struct {
		input    string
		expected []string
	}{{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			output := make([]string, 0)
			allPermutationsRec(string(ts[i].input), "", &output)
			assert.Equal(t, ts[i].expected, output)

		})
	}
}

func Test_PermutationsIter(t *testing.T) {
	var ts = []struct {
		input    string
		expected []string
	}{{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := allPermutationsIter([]byte(ts[i].input))
			assert.Equal(t, ts[i].expected, result)

		})
	}
}

// write compressor: 'AATTAAAARYYY' -> '2A2T4AR3Y'
func Test_StringCompressorRec(t *testing.T) {
	var ts = []struct {
		input    string
		expected string
	}{{"AATTAAAARYYY", "2A2T4AR3Y"}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := stringCompressorRec(ts[i].input)
			assert.Equal(t, ts[i].expected, result)

		})
	}
}

func Test_StringCompressorIter(t *testing.T) {
	var ts = []struct {
		input    string
		expected string
	}{{"AATTAAAARYYY", "2A2T4AR3Y"}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := stringCompressorIter(ts[i].input)
			assert.Equal(t, ts[i].expected, result)

		})
	}
}

// Reverse individual words
// Given a string with "words" (something which is divided by one or more whitespaces).
// Return a string with each "word" reversed  retaining the same amount of whitespaces between them.
// Example:
// input:  "abc def   lll"
// output: "cba fed   lll"

func Test_ReverseStrWithSpaces(t *testing.T) {
	var ts = []struct {
		input    string
		expected string
	}{{"abc def   lll", "cba fed   lll"}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := reverseStrWithSpaces(ts[i].input)
			assert.Equal(t, ts[i].expected, result)

		})
	}
}

func Test_lexicographicalMaxStringt(t *testing.T) {
	var ts = []struct {
		input    string
		expected string
	}{{"ababaa", "babaa"}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			assert.Equal(t, ts[i].expected, lexicographicalMaxString(ts[i].input))

		})
	}
}

func Test_LongestCommonPrefixHorizontal(t *testing.T) {
	var ts = []struct {
		input    []string
		expected string
	}{{[]string{"flower", "flow", "flight"}, "fl"}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			assert.Equal(t, ts[i].expected, longestCommonPrefixHorz(ts[i].input))

		})
	}
}
func Test_LongestCommonPrefixVertical(t *testing.T) {
	var ts = []struct {
		input    []string
		expected string
	}{{[]string{"flower", "flow", "flight"}, "fl"}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			assert.Equal(t, ts[i].expected, longestCommonPrefixVert(ts[i].input))

		})
	}
}
