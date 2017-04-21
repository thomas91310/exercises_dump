/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"fmt"
	"math"
)

// FindPrimeBelow returns all the primes below a certain number
func FindPrimeBelow(n int) []int {
	result := []int{}
	for i := 2; i < n; i++ {
		candidates := candidates(i) // e.g [30, 29 ... 2] for 30
		canDivide := false
		for _, candidate := range candidates {
			if i%candidate == 0 {
				canDivide = true
				break
			}
		}
		if !canDivide {
			result = append(result, i)
		}
	}
	return result
}

func candidates(n int) []int {
	var result []int
	fmt.Printf("for %v, we will go to %v max\n", n, math.Sqrt(float64(n)))
	// stop := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 2; i < n; i++ {
		result = append(result, i)
	}
	return result
}

// LargestPrime returns the largest prime of the number passed in parameter
func LargestPrime(n int) int {
	for i := n - 1; i > 1; i-- {
		if n%i == 0 {
			return i
		}
	}
	return -1
}
