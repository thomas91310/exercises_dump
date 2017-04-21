package main

import "fmt"

// (0,0) -> (2,2)
// (0,1), (1,1), (1,2), (2,2)
// (0,1), (0,2), (1,2), (2,2)
// (0,1), (1,1), (2,1), (2,2)
// (1,0), (1,1), (1,2), (2,2)
// (1,0), (2,0), (2,1), (2,2)
// (1,0), (1,1), (2,1), (2,2)

// (0,0) -> (3,3)
// (0,1), (0,2), (0,3), (1,3), (2,3), (3,3)
// (0,1), (0,2), (1,2), (1,3), (2,3), (3,3)
// (0,1), (0,2), (1,2), (2,2), (2,3), (3,3)
// (0,1), (0,2), (1,2), (2,2), (3,2), (3,3)
// (0,1), (1,1), (1,2), (1,3), (2,3), (3,3)
// (0,1), (1,1), (2,1), (2,2), (2,3), (3,3)
// (0,1), (1,1), (2,1), (2,2), (3,2), (3,3)
// (0,1), (1,1), (2,1), (3,1), (3,2), (3,3)
// (1,0), (2,0), (3,0), (3,1), (3,2), (3,3)
// (1,0), (2,0), (2,1), (3,1), (3,2), (3,3)
// (1,0), (2,0), (2,1), (2,2), (2,3), (3,3)
// (1,0), (2,0), (2,1), (2,2), (3,2), (3,3)
// (1,0), (1,1), (1,2), (1,3), (2,3), (3,3)
// (1,0), (1,1), (1,2), (2,2), (2,3), (3,3)
// (1,0), (1,1), (1,2), (2,2), (3,2), (3,3)
// (1,0), (1,1), (2,1), (3,1), (3,2), (3,3)
// (1,0), (1,1), (2,1), (2,2), (3,2), (3,3)
// (1,0), (1,1), (2,1), (2,2), (2,3), (3,3)

// 2 6
// 3 18

func numberOfValidPaths(sizeOfGrid int) int {
	// return testFunc(sizeOfGrid)
	// return helper(0, 0, sizeOfGrid)
	return routes(sizeOfGrid)
}

func testFunc(sizeOfGrid int) int {
	paths := 1
	for i := 0; i < sizeOfGrid; i++ {
		paths *= (2 * sizeOfGrid) - i
		paths /= i + 1
	}
	return paths
}

func routes(size int) int {
	grid := make([][]int, size+1)

	for i := 0; i < size+1; i++ {
		grid[i] = make([]int, size+1)
	}

	// there is only one route to the corner from the far edges
	for i := 0; i < size; i++ {
		grid[size][i] = 1
		grid[i][size] = 1
	}
	// and no routes to the corner for the corner itself
	grid[size][size] = 0

	for x := size - 1; x >= 0; x-- {
		for y := size - 1; y >= 0; y-- {
			grid[x][y] = grid[x+1][y] + grid[x][y+1]
		}
	}

	return grid[0][0]
}

func helper(x int, y int, sizeOfGrid int) int {
	paths := 0
	if x == sizeOfGrid && y == sizeOfGrid {
		return 1
	}
	if x >= 0 && x < sizeOfGrid {
		paths += helper(x+1, y, sizeOfGrid)
		fmt.Printf("x : %v, y: %v\n", x, y)
	}
	if y >= 0 && y < sizeOfGrid {
		paths += helper(x, y+1, sizeOfGrid)
		fmt.Printf("x : %v, y: %v\n", x, y)
	}
	return paths
}

// 							(0,0)
// 						/   		\
// 					(1,0)			(0,1)
//				/		\		/			\
//			(2,0)		(1,1)	(1,1)		(0,2)
//....

// for i := 0; i < size; i++ {
// 		grid[size][i] = 1
// 		grid[i][size] = 1
// 	}
// 	// and no routes to the corner for the corner itself
// 	grid[size][size] = 0

// 	for x := size - 1; x >= 0; x-- {
// 		for y := size - 1; y >= 0; y-- {
// 			grid[x][y] = grid[x+1][y] + grid[x][y+1]
// 		}
// 	}

// (20,0), (20,1), (20,2) ... (20,19) = 1
// (0,20), (1,20), (2,20) ... (19,20) = 1

// x = 19				x = 18
// y = [19...1]		y = [19...1]
// (19,19) = 2		(18,19) = 3
// (19,18) = 3
// (19,17) = 4
// (19,16) = 5
// ....
// (19,0) = 20
