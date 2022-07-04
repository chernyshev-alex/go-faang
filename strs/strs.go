package strs

import (
	"sort"
	"strings"
	"unicode"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func alloc_ints_with(size uint, values int) []int {
	arr := make([]int, 255)
	arr[0] = -1
	for i := 1; i < len(arr); i *= 2 {
		copy(arr[i:], arr[:i])
	}
	return arr
}

func backSpaceCompare(s string, t string) bool {
	fbuild := func(s string) (res []byte) {
		for _, c := range s {
			if c != '#' {
				res = append(res, byte(c))
			} else if len(res) > 0 {
				res = res[:len(res)-1]
			}
		}
		return
	}
	return string(fbuild(s)) == string(fbuild(t))
}

func backSpaceCompareRev(s string, t string) bool {
	fcmp := func(i, j int) bool { return i > j }
	ftransform := func(s []rune) []rune {
		result := make([]rune, len(s))
		skip := 0
		for _, c := range s {
			if c == '#' {
				skip += 1
			} else if skip > 0 {
				skip -= 1
			} else {
				result = append(result, rune(c))
			}
		}
		return result
	}

	var sr, tr = []rune(s), []rune(t)
	sort.SliceStable(sr, fcmp)
	sort.SliceStable(tr, fcmp)

	return string(ftransform(sr)) == string(ftransform(tr))

}

func reverse(s string) string {
	a := []byte(s)
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return string(a)
}

func isPalindrome(s string) bool {
	for b, l, r := []byte(strings.ToLower(s)), 0, len(s)-1; l < r; l, r = l+1, r-1 {
		for l < r && !(unicode.IsDigit(rune(b[l])) || unicode.IsLetter(rune(b[l]))) {
			l++
		}
		for l < r && !(unicode.IsDigit(rune(b[r])) || unicode.IsLetter(rune(b[r]))) {
			r--
		}
		if b[l] != b[r] {
			return false
		}
	}
	return true
}

func validPalindrome(s string) bool {
	isPalindrome := func(b []byte, l, r int) bool {
		for ; l < r; l, r = l+1, r-1 {
			if b[l] != b[r] {
				return false
			}
		}
		return true
	}

	for b, l, r := []byte(s), 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if b[l] != b[r] {
			return isPalindrome(b, l+1, r) || isPalindrome(b, l, r-1)
		}
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	ln := len(s)
	if ln <= 1 {
		return ln
	}

	var result, l, seen = 0, 0, alloc_ints_with(255, -1)
	for r, c := range s {
		seen_index := seen[c]
		if seen_index >= l && seen_index != -1 {
			l = seen_index + 1
		}
		seen[c] = r
		result = max(result, r-l+1)
	}
	return result
}
