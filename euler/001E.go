package euler

// https://github.com/strelec/Project-Euler-with-Scala

func eu001(value int) int {
	total := 0
	for i := 1; i < value; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}

func eu012(divisors int) int {
	var s = []int{}
	for i := 1; i < 10; i++ {
		x := i * (i + 1) / 2
		if divisors < numberOfDivisors(x) {
			s = append(s, x)
		}
	}
	return s[0]
}

// ----  utils ----

func numberOfDivisors(n int) (divisors int) {
	divisors = 1
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			divisors++
		}
	}
	return
}

func sieve(n int) []int {
	prime := make([]bool, n)
	for i := range prime {
		prime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if prime[i] {
			for k := i * i; i < n; k += i {
				prime[k] = false
			}
		}
	}
	result := make([]int, 0, n)
	for i := range prime {
		if prime[i] {
			result = append(result, i)
		}
	}
	return result
}
