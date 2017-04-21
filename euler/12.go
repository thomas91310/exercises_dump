package main

import (
	"math"
)

func computeNthArithmeticSum(n int) int {
	return n * (n + 1) / 2
}

func highlyDivisibleTriangleNumber(numberOfDivisors int) int {
	i := 0
	for {
		result := computeNthArithmeticSum(i)
		nDivisors := returnNumberOfDivisors(result)
		// fmt.Printf("nDivisors of %v : %v\n", result, nDivisors)
		if nDivisors >= numberOfDivisors {
			return result
		}
		i++
	}
}

func returnNumberOfDivisors(n int) int {
	count := 0
	limit := int(math.Sqrt(float64(n)))
	for i := 1; i <= limit; i++ {
		if n%i == 0 {
			count += 2
		}
	}
	if limit*limit == n {
		count--
	}

	return count
}
