package main

import (
	"math"
)

func sumSquareDifference(limit int) float64 {
	sumSquares := sumOfTheSquares(limit)
	squareOfTheSum := squareOfTheSum(limit)
	return squareOfTheSum - sumSquares
}

func sumOfTheSquares(limit int) float64 {
	var result float64
	for i := 0; i <= limit; i++ {
		result += math.Pow(float64(i), float64(2))
	}
	return result
}

func squareOfTheSum(limit int) float64 {
	var result float64
	for i := 0; i <= limit; i++ {
		result += float64(i)
	}
	return math.Pow(result, float64(2))
}
