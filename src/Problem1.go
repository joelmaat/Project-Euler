package main

import "fmt"

/*
Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these
multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

http://projecteuler.net/problem=1
*/

func SeriesSum(n int) int {
	return n * (n + 1) / 2
}

func SumOfMultiples(maximum int, numbers ...int) int {
	total := 0
	maximum--
	for _, number := range numbers {
		total += SeriesSum(maximum/number)*number
	}
	return total
}

func main() {
	fmt.Printf("%d", SumOfMultiples(1000, 3, 5, 7))
}
