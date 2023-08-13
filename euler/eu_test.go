package euler

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/youthlin/stream"
	"github.com/youthlin/stream/types"
)

// add all natural numberrs below 1000 that are mupliplies of 3 or 5
func TestAddNumbers(t *testing.T) {
	assert.Equal(t, 233168, addNumbers(1000))
}

// Find the sum of all the even-valued terms in the Fibonacci sequence which do not exceed four million.
func TestSumFibo(t *testing.T) {
	assert.Equal(t, 4613733, sumFibo())
}

// Find the largest prime factor of a composite number
func TestLargestPrime(t *testing.T) {
	assert.Equal(t, int64(6857), largestPrime(600851475143))
}

// Find the largest palindrome made from the product of two 3-digit numbers
func TestLargest3Palindrome(t *testing.T) {
	assert.Equal(t, 906609, largest3Palindrome())
}

// What is the smallest number divisible by each of the numbers 1 to 20?
func TestSmallestDivision(t *testing.T) {
	result := 0
	for i := 20; i < math.MaxInt; i++ {
		forAll := true
		for j := 2; j < 21; j++ {
			if i%j != 0 {
				forAll = false
				break
			}
		}
		if forAll {
			result = i
			break
		}
	}
	assert.Equal(t, 232792560, result)
}

// What is the difference between the sum of the squares and the square of the sums
func TestSquaresSumDiff(t *testing.T) {
	s1 := lo.Sum(lo.RangeWithSteps(1, 101, 1))
	s2 := lo.Sum(lo.Map(lo.RangeWithSteps(1, 101, 1), func(v int, i int) int { return v * v }))
	assert.Equal(t, 25164150, s1*s1-s2)
}

// Find the 10001st prime
func TestFind1000Prime(t *testing.T) {
	prime := 3
	for n := 1; n <= 10000; prime += 2 {
		isPrime := true
		for j := 3; j*j <= prime; j += 2 {
			if prime%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			if n == 10000 {
				break
			}
			n++
		}
	}
	assert.Equal(t, 104743, prime)
}

// Discover the largest product of five consecutive digits in the 1000-digit number
func TestE8_1000DigitsOf5Conseq2(t *testing.T) {
	s := "123123213212375121"
	wlen := 3
	maxp := 1
	for i := 0; i < wlen; i++ {
		n, _ := strconv.Atoi(string(s[i]))
		maxp *= n
	}
	tmp := maxp
	for i := 1; i+wlen < len(s); i++ {
		n, _ := strconv.Atoi(string(s[i-1]))
		tmp /= n
		n, _ = strconv.Atoi(string(s[i+wlen-1]))
		tmp *= n
		if tmp > maxp {
			maxp = tmp
		}
	}
	assert.Equal(t, 105, maxp)
}

// E9. Find the only Pythagorean triplet, {a, b, c}, for which a + b + c = 1000
func TestTriplet(t *testing.T) {
	limit := stream.IntRange(1, 1000).Filter(func(e types.T) bool {
		n := e.(int)
		return n+int(math.Sqrt(float64(n))) >= 1000
	}).FindFirst().Get().(int)

	rs := stream.IntRange(2, limit).FlatMap(func(b_ types.T) stream.Stream {
		return stream.IntRange(1, b_.(int)).Map(func(a_ types.T) types.R {
			a, b := a_.(int), b_.(int)
			c := 1000 - a - b
			if c*c == a*a+b*b {
				return a * b * c
			}
			return 0
		}).Filter(func(result types.T) bool { return result != 0 })
	}).ReduceFrom(0, func(acc, v types.T) types.T { return acc.(int) + v.(int) })
	assert.Equal(t, 31875000, rs)
}

// E16 What is the sum of the digits of the number ?
func TestSumOfDigits(t *testing.T) {
	s := fmt.Sprintf("%.0f", math.Pow(2, 1000))
	total := 0
	for i := 0; i < len(s); i++ {
		total += int(byte(s[i]) - byte('0'))
	}
	assert.Equal(t, 1366, total)
}

// E18. Maximum Path Sum I
func TestMaxPathSumOne(t *testing.T) {
	file := `75
	95 64
	17 47 82
	18 35 87 10
	20 04 82 47 65
	19 01 23 75 03 34
	88 02 77 73 07 63 67
	99 65 04 28 06 16 70 92
	41 41 26 56 83 40 80 70 33
	41 48 72 33 47 32 37 16 94 29
	53 71 44 65 25 43 91 52 97 51 14
	70 11 33 28 77 73 17 78 39 68 17 57
	91 71 52 38 17 14 91 43 58 50 27 29 48
	63 66 04 68 89 53 67 30 73 16 69 87 40 31
	04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

	// prepare triangle
	lines := strings.Split(file, "\n")
	triangle := make([][]int, len(lines))
	for i, line := range lines {
		splits := strings.Split(line, " ")
		triangle[i] = make([]int, len(splits))
		for j, s := range splits {
			n, _ := strconv.Atoi(strings.TrimSpace(s))
			triangle[i][j] = n
		}
	}

	get := func(r, c int) int {
		if r < 0 || c < 0 || c >= len(triangle[r]) {
			return 0
		}
		return triangle[r][c]
	}

	for i, r := range triangle {
		for j, _ := range r {
			triangle[i][j] += lo.Max([]int{get(i-1, j-1), get(i-1, j)})
		}
	}
	result := lo.Max(triangle[len(triangle)-1])
	assert.Equal(t, 1074, result)
}
