// Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
// If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

// For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

// Evaluate the sum of all the amicable numbers under 10000.
package main

import "math"

func amicableNumbers(n int) ([]int, int) {
	sumDivisorsMem := make(map[int]int)
	var amicableNumbers []int
	sumAmicableNumbers := 0

	for i := 1; i < n; i++ {
		sumOfDivisors := findSumDivisorsOf(i)
		memSumDivisor, exists := sumDivisorsMem[sumOfDivisors]
		if !exists {
			sumDivisorsMem[i] = sumOfDivisors
			continue
		}
		if memSumDivisor == i {
			amicableNumbers = append(amicableNumbers, i)
			sumAmicableNumbers = sumAmicableNumbers + i
			amicableNumbers = append(amicableNumbers, sumOfDivisors)
			sumAmicableNumbers = sumAmicableNumbers + sumOfDivisors
		}
	}
	return amicableNumbers, sumAmicableNumbers
}

func findSumDivisorsOf(n int) int {
	out := 0
	limit := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 2; i < limit; i++ {
		if n%i == 0 {
			if n/i == i {
				out = out + i
			} else {
				out = out + i
				out = out + n/i
			}
		}
	}
	return out + 1
}
