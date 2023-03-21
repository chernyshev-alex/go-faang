package strs

import (
	"fmt"
	"math"
	"regexp"
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

func longestPalindromicSubstring(s string) string {
	maxLen, ps := 0, ""
	subsFn := func(l, r int) {
		for ln := len(s); l >= 0 && r < ln && s[l] == s[r]; l, r = l-1, r+1 {
			fmt.Printf("1) l=%d r=%d  sl=%c sr=%c\n", l, r, s[l], s[r])
			//	l, r = l-1, r+1
		}

		if r-l-1 >= maxLen {
			ps = s[l+1 : r]
			maxLen = r - l - 1
			fmt.Printf("\t2) ps=%s  \n", ps)
		}
	}

	for i := range s {
		subsFn(i, i)   // check for odd length
		subsFn(i, i+1) // check for event length
	}
	return ps
}

func longestPalindromicSubstring__(s string) string {
	maxLen, ps := 0, ""
	subsFn := func(l, r int) {
		for ln := len(s); l >= 0 && r < ln && s[l] == s[r]; l, r = l-1, r+1 {
			//  <-|->
		}

		if r-l-1 >= maxLen {
			ps = s[l+1 : r]
			maxLen = r - l - 1
		}
	}

	for i := range s {
		subsFn(i, i)   // check for odd length
		subsFn(i, i+1) // check for event length
	}
	return ps
}

func subseqIndexesIter(s string, subs string) []int {
	result := make([]int, 0)

	maxPos := len(s) - len(subs)
	for offset := 0; offset < maxPos; {
		for ; offset < maxPos && subs[0] != s[offset]; offset++ {
		}

		startPos := offset
		matched, i := 0, 0
		for ; i < len(subs) && s[offset] == subs[i]; offset, i = offset+1, i+1 {
			matched++
		}
		if matched == len(subs) {
			result = append(result, startPos)
		}

		offset = startPos + 1
	}

	return result
}

func subseqIndexesRegExp(s string, subs string) [][]int {
	re := regexp.MustCompile(subs)
	return re.FindAllStringIndex(s, -1)
}

func subseqIndexesReq(s string, subs string) []int {
	result := make([]int, 0)
	subseqIndexesReqPart(s, "", subs, 0, &result)
	return result
}

func subseqIndexesReqPart(s string, spart string, subs string, pos int, result *[]int) {
	if pos >= len(s)-len(subs) {
		return
	}

	if spart == subs {
		*result = append(*result, pos-1)
		subseqIndexesReqPart(s, s[pos:pos+len(subs)], subs, pos+1, result)
		return
	}

	subseqIndexesReqPart(s, s[pos:pos+len(subs)], subs, pos+1, result)
}

func allSubsets(chars []byte) [][]byte {
	all_counter := int(math.Pow(2, float64(len(chars))))
	result, partial_result := make([][]byte, 0), make([]byte, 0)

	for i := 0; i < all_counter; i++ {
		for j := 0; j < len(chars); j++ {
			if i&(1<<j) > 0 {
				partial_result = append(partial_result, chars[j])
			}
		}
		result = append(result, partial_result)
		partial_result = make([]byte, 0)

	}
	return result
}

func allPermutationsRec(s string, prefix string, output *[]string) {
	if len(s) == 0 {
		*output = append(*output, prefix)
	}
	for i := range s {
		allPermutationsRec(s[0:i]+s[i+1:], prefix+s[i:i+1], output)
	}
}

func allPermutationsIter(s []byte) []string {
	result, indexes := make([]string, 0), make([]int, len(s))
	for i := 0; i < len(s); {
		if indexes[i] < i {
			if i%2 == 0 {
				s[0], s[i] = s[i], s[0]
			} else {
				s[indexes[i]], s[i] = s[i], s[indexes[i]]
			}
			indexes[i]++
			i = 0

		} else {
			indexes[i] = 0
			i++
		}
		result = append(result, string(s))
	}
	return result
}

// "AATTAAAARYYY", "2A2T4AR3Y"
func stringCompressorRec(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	buf := ""
	stringCompressorRec__(s, 0, 1, &buf)
	return buf
}

func stringCompressorRec__(s string, l, r int, result *string) {
	fmtPart := func() string {
		if r-l == 1 {
			return fmt.Sprintf("%c", s[l])
		}
		return fmt.Sprintf("%d%c", r-l, s[l])
	}

	if r >= len(s) {
		*result = *result + fmtPart()
		return
	}

	if s[l] != s[r] {
		*result = *result + fmtPart()
		stringCompressorRec__(s, r, r+1, result)
		return
	}

	if s[l] == s[r] {
		stringCompressorRec__(s, l, r+1, result)
	}

}

func stringCompressorIter(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}

	fmtPart := func(l, r int) string {
		if r-l == 1 {
			return fmt.Sprintf("%c", s[l])
		}
		return fmt.Sprintf("%d%c", r-l, s[l])
	}

	result := ""
	l, r := 0, 0
	for r < len(s) {
		if s[l] == s[r] {
			r++
		} else {
			result += fmtPart(l, r)
			l = r
		}
	}
	result += fmtPart(l, r)
	return result
}

func reverseStrWithSpaces(s string) string {
	result, stack := make([]rune, 0), make([]rune, 0)

	for _, v := range s {
		if v != ' ' {
			stack = append([]rune{v}, stack...)

		} else {
			for len(stack) > 0 { // pop from stack
				var c rune
				c, stack = stack[0], stack[1:]
				result = append(result, c)
			}
			result = append(result, ' ')
		}
	}
	// if no space after last word
	for len(stack) > 0 {
		var c rune
		c, stack = stack[0], stack[1:]
		result = append(result, c)
	}

	return string(result)
}

func lexicographicalMaxString(s string) string {
	result := ""
	for i := range s {
		if s[i:] > result {
			result = s[i:]
		}
	}
	return result
}

func longestCommonPrefixHorz(ls []string) string {
	if len(ls) == 0 {
		return ""
	}
	prefix := ls[0]
	for i := 1; i < len(ls); i++ {
		s := ls[i]
		if len(s) == 0 || len(prefix) == 0 {
			return ""
		}

		prefix = prefix[0:min(len(prefix), len(s))]

		for k := 0; k < len(s) && k < len(prefix); k++ {
			if s[k] != prefix[k] {
				prefix = prefix[0:k]
				break
			}
		}
	}
	return prefix
}

func longestCommonPrefixVert(ls []string) string {
	if len(ls) == 0 {
		return ""
	}
	for i := 0; i < len(ls[0]); i++ {
		ch := ls[0][i]
		for j := 1; j < len(ls); j++ {
			if i == len(ls[j]) || ls[j][i] != ch {
				return ls[0][:i]
			}
		}
	}
	return ""
}

func longestCommonPrefixVertRec(ls []string) string {
	return longestCommonPrefixVertRecCmp(0, ls)
}

func longestCommonPrefixVertRecCmp(idx int, ls []string) string {
	for si := 1; si < len(ls); si++ {
		if idx == len(ls[si]) || ls[si][idx] != ls[0][idx] {
			return ls[0][:idx]
		}
	}
	return longestCommonPrefixVertRecCmp(idx+1, ls)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
