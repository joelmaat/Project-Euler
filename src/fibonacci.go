package main

import (
	"fmt"
	"math/big"
	"time"
	"reflect"
	"runtime"
)

type SquareRooted struct {
	// Meaning: coefficient * sqrt(rooted) + constant
	coefficient, rooted, constant *big.Rat
}

func (rooted *SquareRooted) Set(template *SquareRooted) *SquareRooted {
	if template == rooted {
		return rooted
	}
	if rooted.coefficient == nil || rooted.rooted == nil || rooted.constant == nil {
		rooted.coefficient, rooted.rooted, rooted.constant = new(big.Rat), new(big.Rat), new(big.Rat)
	}
	rooted.coefficient.Set(template.coefficient)
	rooted.rooted.Set(template.rooted)
	rooted.constant.Set(template.constant)
	return rooted
}

func (rooted *SquareRooted) Multiply(coefficient, constant *big.Rat) *SquareRooted {
	coefficient, constant = new(big.Rat).Set(coefficient), new(big.Rat).Set(constant)
	scratch, rcoefficient := big.NewRat(1, 1), new(big.Rat).Set(rooted.coefficient)
	rooted.coefficient.Mul(rcoefficient, constant)
	rooted.coefficient.Add(rooted.coefficient, scratch.Mul(rooted.constant, coefficient))
	rooted.constant.Mul(rooted.constant, constant)
	rooted.constant.Add(rooted.constant, scratch.Mul(rcoefficient, rooted.rooted).Mul(scratch, coefficient))
	return rooted
}

func (rooted *SquareRooted) Square() *SquareRooted {
	return rooted.Multiply(rooted.coefficient, rooted.constant)
}

func (rooted *SquareRooted) Exponentiate(power uint64) *SquareRooted {
	base := new(SquareRooted).Set(rooted)
	rooted.coefficient.SetInt64(0)
	rooted.constant.SetInt64(1)
	for remaining := power; remaining > 0; remaining >>= 1 {
		if remaining&1 != 0 {
			rooted.Multiply(base.coefficient, base.constant)
		}
		if remaining > 1 {
			base.Square()
		}
	}
	return rooted
}

func Fibonacci(n uint64) *big.Int {
	phi := SquareRooted{big.NewRat(1, 2), big.NewRat(5, 1), big.NewRat(1, 2)}
	return phi.Exponentiate(n).coefficient.Mul(phi.coefficient, big.NewRat(2, 1)).Num()
}

func FibonacciSum(n uint64) *big.Int {
	first, second := big.NewInt(1), big.NewInt(0)
	for i := n; i > 0; i-- {
		second.Add(first, second)
		first.Sub(second, first)
	}
	return second
}

func Multiply(result, base []*big.Int) []*big.Int {
	result[1].Mul(result[1], base[1])
	result[0].Mul(result[0], base[0]).Add(result[0], result[1])
	result[2].Mul(result[2], base[2]).Add(result[2], result[1])
	result[1].Sub(result[2], result[0])
	return result
}

func Square(base []*big.Int) []*big.Int {
	two := big.NewInt(2)
	base[1].Exp(base[1], two, nil)
	base[0].Exp(base[0], two, nil).Add(base[0], base[1])
	base[2].Exp(base[2], two, nil).Add(base[2], base[1])
	base[1].Sub(base[2], base[0])
	return base
}

func FibonacciMatrix(power uint64) *big.Int {
	base := []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(1)}
	result := []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(1)}
	if power < 3 {
		return result[power]
	}
	for remaining := power - 2; remaining > 0; remaining >>= 1 {
		if remaining&1 != 0 {
			Multiply(result, base)
		}
		if remaining > 1 {
			Square(base)
		}
	}
	return result[2]
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	//		fmt.Printf("%d\n", FibonacciMatrix(uint64(17))); return
	for _, fn := range []func(uint64) *big.Int{FibonacciMatrix} {
		delay, count := time.Now(), int64(0)
		for j := 100; j > 0; j-- {
			fn(uint64((1 << 17) - 1))
		}
		count = time.Since(delay).Nanoseconds()
		fmt.Printf("%30s: %15d\n", runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), count)
	}
}
