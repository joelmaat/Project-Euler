package main

import (
	"fmt"
	"math/big"
)

/*
Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is
9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

http://projecteuler.net/problem=4
*/

// Reused from some other person
func SqrtBig(n *big.Int) (x *big.Int) {
	switch n.Sign() {
	case -1:
		panic(-1)
	case 0:
		return big.NewInt(0)
	}
	var px, nx big.Int
	x = big.NewInt(0)
	x.SetBit(x, n.BitLen()/2+1, 1)
	for {
		nx.Rsh(nx.Add(x, nx.Div(n, x)), 1)
		if nx.Cmp(x) == 0 || nx.Cmp(&px) == 0 {
			break
		}
		px.Set(x)
		x.Set(&nx)
	}
	return
}

func GreatestDivisor(number, base *big.Int) *big.Int {
	divisor, number, base := big.NewInt(1), new(big.Int).Set(number), new(big.Int).Set(base)
	if number.Cmp(base) < 0 {
		return divisor
	}
	divisor.Set(base)
	candidate := new(big.Int).Mul(divisor, base)
	for ; number.Cmp(candidate) >= 0 && candidate.Cmp(divisor) > 0; candidate.Mul(candidate, base) {
		divisor.Set(candidate)
	}
	return divisor
}

func MakePalindrome(half *big.Int) *big.Int {
	ten, zero, digit, half := big.NewInt(10), big.NewInt(0), big.NewInt(0), new(big.Int).Set(half)
	palindrome, multiplier, divisor := new(big.Int).Set(half), big.NewInt(1), GreatestDivisor(half, ten)
	palindrome.Mul(half, divisor).Mul(palindrome, ten)
	for divisor.Cmp(zero) > 0 {
		digit.DivMod(half, divisor, half)
		if digit.Cmp(zero) > 0 {
			palindrome.Add(palindrome, digit.Mul(digit, multiplier))
		}
		divisor.Div(divisor, ten)
		multiplier.Mul(multiplier, ten)
	}
	return palindrome
}

func LargestPalindrome(digits uint64) *big.Int {
	if digits == 0 {
		panic("Number of digits has to be greater than or equal to 1.")
	} else if digits == 1 {
		return big.NewInt(9)
	}
	one, two, ten, palindrome, high := big.NewInt(1), big.NewInt(2), big.NewInt(10), big.NewInt(0), big.NewInt(0)
	low, sqrt, power, scratch := big.NewInt(0), big.NewInt(0), new(big.Int).SetUint64(digits), big.NewInt(0)
	middle, right, maximum := big.NewInt(0), big.NewInt(0), big.NewInt(0)
	maximum.Exp(ten, power, nil).Sub(maximum, one)
	minimum := new(big.Int).Exp(ten, scratch.Sub(power, one), nil)
	ten_m := new(big.Int).Mul(ten, minimum)
	skip := big.NewInt(int64(digits / 2))
	if digits == 10 || digits == 11 {
		skip.Sub(skip, one)
	}
	skip.Exp(ten, skip, nil).Sub(skip, one)
	for half := new(big.Int).Sub(maximum, skip); half.Cmp(minimum) >= 0; half.Sub(half, one) {
		palindrome = MakePalindrome(half)
		sqrt.Set(SqrtBig(palindrome)).Add(sqrt, one)
		high.Sub(ten_m, one)
		low.Set(sqrt)
		right.Sub(high, sqrt).Sub(sqrt, right)
		if high.Cmp(low) < 0 ||
			scratch.Add(high, half).Add(scratch, one).Div(scratch, two).Cmp(sqrt) != 0 ||
			scratch.Mul(high, right).Cmp(palindrome) > 0 {
			continue
		} else if scratch.Cmp(palindrome) == 0 {
			return palindrome
		}
		for high.Cmp(low) >= 0 {
			middle.Add(high, low).Div(middle, two)
			right.Sub(middle, sqrt).Sub(sqrt, right)
			if scratch.Mul(middle, right).Cmp(palindrome) == 0 {
				return palindrome
			} else if scratch.Cmp(palindrome) > 0 {
				low.Add(middle, one)
			} else {
				high.Sub(middle, one)
			}
		}
	}
	return palindrome
}

func main() {
	fmt.Printf("%d", LargestPalindrome(uint64(3)))
}
