package main

import "fmt"

/*
Summation of primes

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

http://projecteuler.net/problem=10
*/

func IsPrime(number uint64) bool {
	switch number {
	case 1:
		return false
	case 2, 3, 5, 7:
		return true
	}
	if number == 1 || number&1 == 0 {
		return false
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

func SummationOfPrimes(limit uint64) uint64 {
	digits, sum := []uint64{1, 3, 7, 9}, uint64(0)
	if limit > 5 {
		sum += 7
	} else if limit > 2 {
		sum += 2
	}
	for base := uint64(0); base < limit; base += 10 {
		for _, digit := range digits {
			candidate := base + digit
			if IsPrime(candidate) && candidate < limit {
				sum += candidate
			}
		}
	}
	return sum
}

func main() {
	fmt.Printf("%d", SummationOfPrimes(uint64(2000000)))
}
