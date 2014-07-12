package main

import "fmt"
import "math/big"

/*
Smallest multiple

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

http://projecteuler.net/problem=5
*/

type Factor struct {
	prime    *big.Int
	power    *big.Int
	expanded *big.Int
}

func IsPrime(number *big.Int) bool {
	return number.ProbablyPrime(4)
}

func Min(numbers ...*big.Int) *big.Int {
	smallest := new(big.Int).Set(numbers[0])
	for _, number := range numbers {
		if number.Cmp(smallest) < 0 {
			smallest.Set(number)
		}
	}
	return smallest
}

func NextPrime(minimum, maximum *big.Int) *big.Int {
	two, four, five, ten, base := big.NewInt(2), big.NewInt(4), big.NewInt(5), big.NewInt(10), big.NewInt(0)
	if minimum.Cmp(two) == 0 {
		return Min(two, new(big.Int).Set(maximum))
	} else if minimum.Cmp(four) == 0 || minimum.Cmp(five) == 0 {
		return Min(five, new(big.Int).Set(maximum))
	}
	digits := []*big.Int{big.NewInt(1), big.NewInt(3), big.NewInt(7), big.NewInt(9)}
	for base.Mod(minimum, ten).Sub(minimum, base); base.Cmp(maximum) <= 0; base.Add(base, ten) {
		for _, digit := range digits {
			candidate := new(big.Int).Add(digit, base)
			if candidate.Cmp(maximum) <= 0 && IsPrime(candidate) && candidate.Cmp(minimum) >= 0 {
				return candidate
			}
		}
	}
	return maximum
}

func LargestPowerDivisibleBy(number, base *big.Int) (*big.Int, *big.Int) {
	zero, one, two, scratch := big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(0)
	if base.Cmp(one) <= 0 || new(big.Int).Mod(number, base).Cmp(zero) != 0 {
		return zero, one
	}
	smallest, largest, middle := big.NewInt(1), new(big.Int).Set(number), big.NewInt(0) // should be log of number
	for scratch.Sub(largest, smallest).Cmp(one) > 0 {
		middle.Add(smallest, largest).Div(middle, two)
		if scratch.Exp(base, middle, nil).Mod(number, scratch).Cmp(zero) == 0 {
			smallest.Set(middle)
		} else {
			largest.Sub(middle, one)
		}
	}
	if smallest.Cmp(largest) != 0 && scratch.Exp(base, largest, nil).Mod(number, scratch).Cmp(zero) != 0 {
		largest.Set(smallest)
	}
	return largest, scratch.Exp(base, largest, nil)
}

func PrimeFactors(number *big.Int) *[]Factor {
	zero, one, two := big.NewInt(0), big.NewInt(1), big.NewInt(2)
	factors, maximum, prime := []Factor{}, new(big.Int).Div(number, two), big.NewInt(1)
	if number.Cmp(one) < 0 {
		panic("Number needs to be greater than 0.")
	} else if number.Cmp(one) == 0 {
		return &[]Factor{Factor{big.NewInt(1), big.NewInt(1), big.NewInt(1)}}
	}
	for number.Cmp(one) > 0 {
		if IsPrime(number) {
			factors = append(factors, Factor{new(big.Int).Set(number), big.NewInt(1), new(big.Int).Set(number)})
			break
		}
		prime = NextPrime(new(big.Int).Add(prime, one), maximum)
		power, expanded := LargestPowerDivisibleBy(number, prime)
		if power.Cmp(zero) > 0 {
			factors = append(factors, Factor{new(big.Int).Set(prime), new(big.Int).Set(power), new(big.Int).Set(expanded)})
			number.Div(number, expanded)
		}
	}
	return &factors
}

func SmallestMultiple(start, end int) *big.Int {
	factors, multiple := map[*big.Int]Factor{}, big.NewInt(1)
	for ; start <= end; start++ {
		for _, factor := range *PrimeFactors(big.NewInt(int64(start))) {
			_, exists := factors[factor.prime]
			if !exists || factors[factor.prime].power.Cmp(factor.power) < 0 {
				factors[factor.prime] = factor
			}
		}
	}
	for _, factor := range factors {
		multiple.Mul(multiple, factor.expanded)
	}
	return multiple
}

func main() {
	fmt.Printf("%d", SmallestMultiple(1, 20))
}
