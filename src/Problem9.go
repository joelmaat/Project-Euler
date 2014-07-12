package main

import (
	"fmt"
	"math"
)

/*
Special Pythagorean triplet

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

http://projecteuler.net/problem=9
*/

func SpecialPythagoreanTriplet(sum uint64) uint64 {
	for limit, b := (sum + 1) / 2, sum + 1 - uint64(math.Ceil(math.Sqrt(float64(sum * sum) / 2.0))); b < limit; b++ {
		a := sum * (sum - 2 * b) / (2 * (sum - b))
		c := sum - a - b
		if a*a+b*b == c*c && a < b && b < c {
			return a * b * c
		}
	}
	return 0
}

func main() {
	fmt.Printf("%d", SpecialPythagoreanTriplet(uint64(1000)))
}
