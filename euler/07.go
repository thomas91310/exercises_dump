package main

// findNthPrime returns the nth prime number
func findNthPrime(n int) int {
	i := 2
	count := 0
	for {
		candidates := candidates(i) // e.g [30, 29 ... 2] for 30
		canDivide := false
		for _, candidate := range candidates {
			if i%candidate == 0 {
				canDivide = true
				break
			}
		}
		if !canDivide {
			count++
		}
		if count == n {
			return i
		}
		i++
	}
}
