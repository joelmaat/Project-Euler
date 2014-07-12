package main

import "fmt"

/*
10001st prime

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

http://projecteuler.net/problem=7
*/

func IsPrime(number uint64) bool {
	if number == 1 || number&1 == 0 {
		return number == 2
	}
	var divisor uint64 = 3
	for divisor*divisor <= number {
		if number%divisor == 0 {
			return false
		}
		divisor += 2
	}
	return true
}

func NthPrimeNumber(n uint64) uint64 {
	cache, digits := []uint64{1, 2, 3, 5, 7}, []uint64{1, 3, 7, 9}
	var total, prime, length, base uint64 = uint64(len(cache) - 1), 0, uint64(len(digits)), 10
	if n <= total {
		return cache[n]
	}
	for index := uint64(0); total < n; index++ {
		if index >= length {
			index, base = 0, base+10
		}
		prime = base+digits[index]
		if IsPrime(prime) {
			total++
		}
	}
	return prime
}

func main() {
	fmt.Printf("%d", NthPrimeNumber(uint64(10001)))
}
