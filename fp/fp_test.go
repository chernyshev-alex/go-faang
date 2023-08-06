package fp

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
