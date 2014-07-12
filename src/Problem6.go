package main

import (
	"math/big"
	"fmt"
)

/*
Sum square difference

The sum of the squares of the first ten natural numbers is,

1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)^2 = 55^2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is
3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

http://projecteuler.net/problem=6
*/

func SeriesSum(largest *big.Int) *big.Int {
	// largest * (largest + 1) / 2
	sum := big.NewInt(1)
	sum.Add(largest, sum).Mul(sum, largest).Div(sum, big.NewInt(2))
	return sum
}

func SumOfSquares(largest *big.Int) *big.Int {
	// largest * (1 + 3 * largest + 2 * largest * largest) / 6
	two, three, six, sum, scratch := big.NewInt(2), big.NewInt(3), big.NewInt(6), big.NewInt(1), big.NewInt(0)
	sum.Add(sum, scratch.Mul(three, largest)).Add(sum, scratch.Mul(largest, largest).Mul(scratch, two))
	return sum.Mul(sum, largest).Div(sum, six)
}

func SumSquareDifference(start, end *big.Int) *big.Int {
	sum, excess := big.NewInt(0), new(big.Int).Sub(start, big.NewInt(1))
	sum.Sub(SeriesSum(end), SeriesSum(excess)).Mul(sum, sum)
	return sum.Sub(sum, SumOfSquares(end)).Add(sum, SumOfSquares(excess))
}

func SumSquareDifference2(start, end *big.Int) *big.Int {
	// (start-end-1) * (start-end) * (3*start*start + 6*start*end + start - end + 3*end*end - 2) / 12
	one, two, three, six, twelve := big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(6), big.NewInt(12)
	scratch, sum := big.NewInt(0), big.NewInt(0)
	sum.Sub(scratch.Mul(start, start).Mul(scratch, three), two).Sub(sum, end)
	sum.Add(sum, scratch.Mul(start, end).Mul(scratch, six)).Add(sum, scratch.Mul(end, end).Mul(scratch, three))
	sum.Add(sum, start).Mul(sum, scratch.Sub(start, end)).Mul(sum, scratch.Sub(start, end).Sub(scratch, one))
	return sum.Div(sum, twelve)
}

func main() {
	start, end := "1", "100"
	smallest, _ := new(big.Int).SetString(start, 10)
	largest, _ := new(big.Int).SetString(end, 10)
	fmt.Printf("%d\n", SumSquareDifference(smallest, largest))
}
