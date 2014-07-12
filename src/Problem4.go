package main

import (
	"bytes"
	"fmt"
	"math/big"
	"runtime"
)

/*
Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is
9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

http://projecteuler.net/problem=4
*/

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

func LargestPalindrome2(digits uint64, ch chan string) *big.Int {
	if digits == 0 {
		panic("Number of digits has to be greater than or equal to 1.")
	} else if digits == 1 {
		return big.NewInt(9)
	}
	one, ten, palindrome, left, limit := big.NewInt(1), big.NewInt(10), big.NewInt(0), big.NewInt(0), big.NewInt(0)
	two, power, scratch := big.NewInt(2), new(big.Int).SetUint64(digits), big.NewInt(0)
	maximum := new(big.Int).Exp(ten, power, nil)
	maximum.Sub(maximum, one)
	minimum := new(big.Int).Exp(ten, power.Sub(power, one), nil)
	ten_m := new(big.Int).Mul(ten, minimum)
	counter, counter2 := 0, 0
	for half := new(big.Int).Set(maximum); half.Cmp(minimum) >= 0; half.Sub(half, one) {
		palindrome = MakePalindrome(half)
		counter++
		left.Sub(ten_m, one)
		limit.Add(half, left).Div(limit, two)
		for ; left.Cmp(limit) >= 0; left.Sub(left, one) {
			counter2++
			scratch.Sub(left, limit).Sub(limit, scratch)
			if scratch.Div(palindrome, left).Mul(scratch, left).Cmp(palindrome) > 0 {
				break
			} else if scratch.Cmp(palindrome) == 0 {
				ch <- fmt.Sprintf("%10d %15d %20d %10d %20d %20d %20d %20d %3d %d\n", counter, counter2, minimum, new(big.Int).Sub(maximum, left), maximum, limit, left, scratch.Div(scratch, left), digits, palindrome)
				return palindrome
			}
			//			if > palindrome || scratch.Div(palindrome, left).Cmp(left) > 0 {
			//				break
			//			}
		}
	}
	return palindrome
}

func LargestPalindrome3(digits uint64, ch chan string) *big.Int {
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
	counter, counter2 := 0, 0 // ####
	for half := new(big.Int).Sub(maximum, skip); half.Cmp(minimum) >= 0; half.Sub(half, one) {
		counter++
		counter2++
		palindrome = MakePalindrome(half)
		sqrt.Set(SqrtBig(palindrome)).Add(sqrt, one)
		high.Sub(ten_m, one)
		low.Set(sqrt)
		right.Sub(high, sqrt).Sub(sqrt, right)
		if high.Cmp(low) < 0 || scratch.Add(high, half).Add(scratch, one).Div(scratch, two).Cmp(sqrt) != 0 || scratch.Mul(high, right).Cmp(palindrome) > 0 {
			continue
		} else if scratch.Cmp(palindrome) == 0 {
			ch <- fmt.Sprintf("h%10d %15d %10d %50d %50d %3d %d\n", counter, counter2, new(big.Int).Sub(maximum, high), high, scratch.Div(scratch, high), digits, palindrome)
			return palindrome
		}
		for high.Cmp(low) >= 0 {
			counter2++
			middle.Add(high, low).Div(middle, two)
			right.Sub(middle, sqrt).Sub(sqrt, right)
			if scratch.Mul(middle, right).Cmp(palindrome) == 0 {
				ch <- fmt.Sprintf("n%10d %15d %10d %50d %50d %3d %d\n", counter, counter2, new(big.Int).Sub(maximum, high), high, scratch.Div(scratch, high), digits, palindrome)
				return palindrome
			} else if scratch.Cmp(palindrome) > 0 {
				low.Add(middle, one)
			} else {
				high.Sub(middle, one)
			}
		}
		counter2--
	}
	return palindrome
}

func LargestPalindrome(digits uint64, ch chan string) *big.Int {
	var string bytes.Buffer
	var palindrome *big.Int
	string.Grow(int(digits * 2 + 1))
	if digits&3 == 0 {
		template := "9009"
		times := int(digits / 2)
		for _, digit := range template {
			for i := 0; i < times; i++ {
				string.WriteRune(digit)
			}
		}
		palindrome, _ = new(big.Int).SetString(string.String(), 10)
	} else {
		palindrome = LargestPalindrome2(digits, ch)
	}
	return palindrome
}

func main() {
	var low, high uint64 = 2, 50
	runtime.GOMAXPROCS(int(high - low + 2))
	ch := make(chan string, high-low)
	for digits := low; digits < high; digits++ {
		go func(digits uint64) { LargestPalindrome3(digits, ch) }(digits)
	}
	for i := low; i < high; i++ {
		fmt.Printf(<-ch)
	}

}
