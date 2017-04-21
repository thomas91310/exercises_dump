package main

func collatzMain(underN int) (int, int) {
	countMaxSeq := 0
	result := 0
	remember := make(map[int]int)
	for i := 0; i <= underN; i++ {
		countSeq := collatzHelper(i, 0, remember)
		remember[i] = countSeq
		if countSeq > countMaxSeq {
			countMaxSeq = countSeq
			result = i
		}
	}
	return result, countMaxSeq
}

// 13 -> 40 -> 20 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1

func collatzHelper(n int, countSequence int, remember map[int]int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return countSequence
	}
	countCache, exists := remember[n]
	if exists {
		return countSequence + countCache
	}
	if n%2 == 0 {
		n = n / 2
	} else {
		n = 3*n + 1
	}
	return collatzHelper(n, countSequence+1, remember)
}
