/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

func sumMultiple(number int) int {
	var toMultiply []int
	var sum int
	for i := 1; i < number; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			toMultiply = append(toMultiply, i)
		}
	}
	for j := 0; j < len(toMultiply); j++ {
		sum = sum + toMultiply[j]
	}
	return sum
}
