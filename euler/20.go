// n! means n × (n − 1) × ... × 3 × 2 × 1

// For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
// and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

// Find the sum of the digits in the number 100!

package main

import "strconv"

func sumFactDigits(n int) int {
	memo := make(map[int]int)
	result := memoizedFibo(n, memo)
	resultString := strconv.Itoa(result)
	sumDigits := 0
	for _, c := range resultString {
		digit, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		sumDigits = sumDigits + digit
	}
	return sumDigits
}
