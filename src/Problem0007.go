package main

import (
	"fmt"
	"math/big"
)

/*
10001st prime

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

http://projecteuler.net/problem=7
*/

func IsPrime(number int) bool {
	return big.NewInt(int64(number)).ProbablyPrime(4)
}

func NthPrimeNumber(n uint64) int {
	cache, digits := []int{1, 2, 3, 5, 7}, []int{1, 3, 7, 9}
	total, prime, digits_length := uint64(len(cache) - 1), 0, len(digits)
	if n <= total {
		return cache[n]
	}
	for base, next := 10, 0; total < n; next++ {
		if next >= digits_length {
			next, base = 0, base+10
		}
		candidate := base + digits[next]
		if IsPrime(candidate) {
			prime = candidate
			total++
		}
	}
	return prime
}

func main() {
	fmt.Printf("%d", NthPrimeNumber(uint64(10001)))
}
