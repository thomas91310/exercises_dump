package main

import "fmt"

func main() {
	A := []int{-40, -20, -10, -5, -2, -1, 2, 4, 5, 9, 10, 20, 30, 34, 37}

	// isMagic, idx := magicIndexBruteForce(A)
	// if isMagic {
	// 	fmt.Printf("Magic index at %v\n", idx)
	// } else {
	// 	fmt.Println("No magic index!")
	// }

	isMagic, idx := magicIndex(A, 0)
	if isMagic {
		fmt.Printf("Magic index at %v\n", idx)
	} else {
		fmt.Println("No magic index!")
	}
}

//pb with that solution is that if the magic index is at the end of the array
//we would loop through all of it before realizing that
//we can do probably a little bit better
func magicIndexBruteForce(A []int) (bool, int) {
	for idx, val := range A {
		if idx == val {
			return true, idx
		}
	}
	return false, 0
}

func magicIndex(A []int, startingIdx int) (bool, int) {
	middle := len(A) / 2
	if A[middle] < middle+startingIdx {
		//must be on the right side
		return magicIndex(A[middle:], middle+startingIdx)
	} else if A[middle] == middle+startingIdx {
		return true, middle + startingIdx
	} else {
		return magicIndex(A[:middle], startingIdx)
	}
}
