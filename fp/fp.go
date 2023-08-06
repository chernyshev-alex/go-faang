package fp

import (
	"math"
	"strconv"

	"github.com/samber/lo"
)

func addNumbers(maxn int) int {
	return lo.Sum(lo.Filter(lo.Range(maxn), func(item int, pos int) bool {
		return item%3 == 0 || item%5 == 0
	}))
}

func sumFibo() int {
	f0, f1 := 0, 1
	fiboGen := func() int { f0, f1 = f1, f0+f1; return f1 }

	total := fiboGen()
	for next := fiboGen(); total <= 4000000; next = fiboGen() {
		if next%2 == 0 {
			total += next
		}
	}
	return total
}

func largestPrime(n int64) int64 {
	var acc []int64
	for v := int64(2); v < int64(math.Sqrt(float64(n))); v++ {
		if n%v == 0 {
			acc = append(acc, v)
			n /= v
		}
	}
	if n > 1 {
		acc = append(acc, n)
	}
	return acc[len(acc)-1]
}

func largest3Palindrome() int {
	nmax := 0
	for i := 100; i <= 999; i++ {
		for j := i; j <= 999; j++ {
			n := i * j
			rev := string(lo.Reverse([]byte(strconv.Itoa(n))))
			if strconv.Itoa(n) == rev && nmax < n {
				nmax = n
			}
		}
	}
	return nmax
}
