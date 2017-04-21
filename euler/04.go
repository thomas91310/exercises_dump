/*A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.*/
package main

import (
	"fmt"
	"math"
	"strconv"
)

// n ^ 2 solution
// start at 100 let's say
// a way to do it would be to compute all the numbers
// so like 100 * 101, 100 * 102 .... 999 * 101, ... 999 * 999
// result instantiated to 0
// if number palindrome and > than current palindrome return value overwrite result
func findPalindrome(numberOfDigits int) (int, int, int) {
	s := math.Pow(float64(10), float64(numberOfDigits-1))
	e := math.Pow(float64(10), float64(numberOfDigits))
	start := int(s)
	end := int(e)
	result := 0
	resFirstNumber := 0
	resSecondNumber := 0
	for i := start; i < end; i++ {
		for j := start; j < end; j++ {
			mult := i * j
			isPalindrome := palindrome(mult)
			if isPalindrome && mult > result {
				resFirstNumber = i
				resSecondNumber = j
				result = mult
			}
		}
	}
	return resFirstNumber, resSecondNumber, result
}

// just an easy way to get multiples that results in a palindrome number
func findPalindromes(numberOfDigits int) {
	s := math.Pow(float64(10), float64(numberOfDigits-1))
	e := math.Pow(float64(10), float64(numberOfDigits))
	start := int(s)
	end := int(e)
	keys := []int{}
	result := make(map[int][]int)
	for i := start; i < end; i++ {
		for j := start; j < end; j++ {
			mult := i * j
			isPalindrome := palindrome(mult)
			if isPalindrome {
				_, exists := result[i]
				if !exists {
					result[i] = []int{j, mult}
					keys = append(keys, i)
				}
			}
		}
	}
	for _, k := range keys {
		fmt.Printf("%v * %v = %v\n", k, result[k][0], result[k][1])
	}
}

func palindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
