package main

// n^2 solution
// for each number we try to divide by 1 to limit
// if one number between 1 and limit didn't evenly divide that number, we go to the next one
func smallestMultiple(limit int) int {
	result := 1
	for {
		foundSmallest := true
		for i := 1; i <= limit; i++ {
			if result%i == 0 {
				continue
			} else {
				foundSmallest = false
				result++
				break
			}
		}
		if foundSmallest {
			return result
		}
	}
}
