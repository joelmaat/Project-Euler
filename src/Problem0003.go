package main

import "fmt"
import "math"

/*
Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

http://projecteuler.net/problem=3
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

func NextPrime(minimum uint64) uint64 {
	digits := []uint64{1, 3, 7, 9}
	for base := minimum - (minimum % 10); ; base += 10 {
		for _, digit := range digits {
			candidate := base + digit
			if IsPrime(candidate) && candidate >= minimum {
				return candidate
			}
		}
	}
	return 0
}

func LargestPowerDivisibleBy(number, base uint64) uint64 {
	if base <= 1 || number%base != 0 {
		return 0
	}
	var smallest, largest uint64 = 1, uint64(math.Log(float64(number)) / math.Log(float64(base)))
	for largest-smallest > 1 {
		middle := (smallest + largest) / 2
		if number%uint64(math.Pow(float64(base), float64(middle))) == 0 {
			smallest = middle
		} else {
			largest = middle-1
		}
	}
	if smallest != largest && number%uint64(math.Pow(float64(base), float64(largest))) != 0 {
		largest = smallest
	}
	return largest
}

func LargestPrimeFactor(number uint64) uint64 {
	var prime uint64 = 1
	for number > 1 {
		if IsPrime(number) {
			return number
		}
		prime = NextPrime(prime+1)
		power := LargestPowerDivisibleBy(number, prime)
		if power > 0 {
			number /= uint64(math.Pow(float64(prime), float64(power)))
		}
	}
	return prime
}

func main() {
	fmt.Printf("%d", LargestPrimeFactor(uint64(600851475143)))
}
