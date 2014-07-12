package main

import "fmt"
import "math"
import "math/big"

/*
Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

http://projecteuler.net/problem=3
*/

func IsPrime(number uint64) bool {
	return big.NewInt(int64(number)).ProbablyPrime(4)
}

func NextPrime(minimum, maximum uint64) uint64 {
	if minimum == 2 {
		return uint64(math.Min(2, float64(maximum)))
	} else if (minimum == 4 || minimum == 5) {
		return uint64(math.Min(5, float64(maximum)))
	}
	digits := []uint64{1, 3, 7, 9}
	for base := minimum - (minimum % 10); base <= maximum; base += 10 {
		for _, digit := range digits {
			candidate := base + digit
			if candidate <= maximum && IsPrime(candidate) && candidate >= minimum {
				return candidate
			}
		}
	}
	return maximum
}

func LargestPowerDivisibleBy(number, base uint64) uint64 {
	if base <= 1 || number%base != 0 {
		return 0
	}
	smallest, largest := uint64(1), uint64(math.Log(float64(number)) / math.Log(float64(base)))
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
	maximum := number / 2
	prime := uint64(1)
	for number > 1 {
		if IsPrime(number) {
			return number
		}
		prime = NextPrime(prime+1, maximum)
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
