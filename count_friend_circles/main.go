package main

import "fmt"

func main() {
	//dummy main, just to have something running
	//Please look at the tests for some more elaborate examples
	arr := makeDummyGrid()
	visited := map[int]bool{}
	fmt.Println(countCircles(arr, visited))
}

func countCircles(arr [][]bool, visited map[int]bool) int {
	counter := 0
	for i := 0; i < len(arr); i++ {
		_, hasBeenVisited := visited[i]
		if hasBeenVisited {
			continue
		}
		friendsRecursion(arr, i, visited)
		counter++
	}
	return counter
}

func friendsRecursion(arr [][]bool, row int, visited map[int]bool) {
	for j := 0; j < len(arr[row]); j++ {
		_, hasBeenVisited := visited[j]
		if j == row || hasBeenVisited {
			continue
		}

		if arr[row][j] {
			visited[row] = true
			friendsRecursion(arr, j, visited)
		}
	}
	visited[row] = true
}

//makeDummyGrid returns a 2D slice of n*n elements
//if arr[i][j] is true then i et j are friends
//otherwise i and j are not friends
func makeDummyGrid() [][]bool {
	arr := make([][]bool, 6)
	arr[0] = []bool{true, true, false, false, false, false}
	arr[1] = []bool{true, true, false, false, false, false}
	arr[2] = []bool{false, false, true, false, false, false}
	arr[3] = []bool{false, false, false, true, true, false}
	arr[4] = []bool{false, false, false, true, true, true}
	arr[5] = []bool{false, false, false, false, true, true}
	return arr
}
