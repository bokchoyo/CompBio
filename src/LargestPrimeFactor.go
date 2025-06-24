package main

import (
	"math"
)

func LargestPrimeFactor(n int) int {
	factors := Factorize(n)
	max := 1

	for _, i := range factors {
		if IsPrime(i) {
			if i > max {
				max = i
			}
		}
	}

	return max
}

func Factorize(num int) []int {
	factors := []int{}

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}

	return factors
}

func IsPrime(p int) bool {
	if p == 0 || p == 1 {
		return false
	} else if p == 2 {
		return true
	}

	sqrt := int(math.Sqrt(math.Abs(float64(p))))

	for i := 2; i <= sqrt; i++ {
		if p%i == 0 {
			return false
		}
	}

	return true
}
