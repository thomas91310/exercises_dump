//By starting at the top of the triangle below and moving to adjacent numbers on the row below,
//the maximum total from top to bottom is 23. That is, 3 + 7 + 4 + 9 = 23.
//Find the maximum total from top to bottom of the triangle below:

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	triangleInput = "75\n95,64\n17,47,82\n18,35,87,10\n20,04,82,47,65\n19,01,23,75,03,34\n88,02,77,73,07,63,67\n99,65,04,28,06,16,70,92\n41,41,26,56,83,40,80,70,33\n41,48,72,33,47,32,37,16,94,29\n53,71,44,65,25,43,91,52,97,51,14\n70,11,33,28,77,73,17,78,39,68,17,57\n91,71,52,38,17,14,91,43,58,50,27,29,48\n63,66,04,68,89,53,67,30,73,16,69,87,40,31\n04,62,98,27,23,09,70,98,73,93,38,53,60,04,2"
	smallTriangle = "3\n7,4\n2,4,6\n8,5,9,3"
)

// maxPathSum recursion tries everything
func maxPathSum(triangle string) (uint64, error) {
	lines := strings.Split(triangle, "\n")
	grid := make([][]uint64, len(lines))
	for idx, line := range lines {
		grid[idx] = make([]uint64, idx+1)
		numbers := strings.Split(line, ",")
		for column, n := range numbers {
			number, err := strconv.ParseUint(n, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("Error when converting the string number into an integer")
			}
			grid[idx][column] = number
		}
	}
	var results []uint64
	results = maxPathSumRecursion(grid, 0, 0, 0, results)
	fmt.Println("results : ", results)
	fmt.Println("biggest sum : ", findBiggestNumber(results))
	return 0, nil
}

// "3\n7,4\n2,4,6\n8,5,9,3"
// 3
// 7 4
// 2 4 6
// 8 5 9 3

// 0,0 -> 1,0 -> 2,0 -> 3,0
//

// step 1 3
// step 2
// (3) + 7 = 10
// (3) + 4 = 7
// (3 + 7) + 2 = 12
// (3 + 7) + 4 = 14
// (3 + 4) + 4 = 11
// (3 + 4) + 6 = 13
// (3 + 7 + 2) + 8 = 20
// (3 + 7 + 2) + 5 = 17
// (3 + 7 + 4) + 8 = 22
// (3 + 7 + 4) + 5 = 19
// (3 + 7 + 4) + 9 = 23
// (3 + 4 + 4) + 8 = 19
// (3 + 4 + 4) + 5 = 16
// (3 + 4 + 4) + 9 = 20
// (3 + 4 + 6) + 9 = 22
// (3 + 4 + 6) + 3 = 16

func maxPathSumRecursion(grid [][]uint64, line int, col int, result uint64, results []uint64) []uint64 {
	// fmt.Println("line ::: ", line)
	nLines := len(grid)
	if line >= nLines {
		results = append(results, result)
		return results
	}
	nCols := len(grid[line])
	// fmt.Printf("for line %v and col %v\n", line, col)
	// fmt.Printf("lim line : %v\nlim col : %v\n", nLines, nCols)
	if col-1 >= 0 && col-1 < nCols && line+1 <= nLines {
		fmt.Printf("left for line %v and col %v\n", line, col-1)
		results = maxPathSumRecursion(grid, line+1, col-1, result+grid[line][col-1], results)
	}
	if col >= 0 && col < nCols && line+1 <= nLines {
		fmt.Printf("middle for line %v and col %v\n", line, col)
		results = maxPathSumRecursion(grid, line+1, col, result+grid[line][col], results)
	}
	if col+1 >= 0 && col+1 < nCols && line+1 <= nLines {
		fmt.Printf("right for line %v and col %v\n", line, col+1)
		results = maxPathSumRecursion(grid, line+1, col+1, result+grid[line][col+1], results)
	}
	return results
}

// misunderstood, adjacent means that you can't go more than 1 column further to where you were at at the
// step before. e.g (Line 1, column 2 second step, third step is either Line 2, column 1 or Line 2, column 2.
func maximumPathMisunderstood(triangle string) (int, error) {
	lines := strings.Split(triangle, "\n")
	result := 0
	for _, line := range lines {
		numbers := strings.Split(line, ",")
		// keep the maximum number encountered for this line
		max := 0
		for _, n := range numbers {
			number, err := strconv.Atoi(n)
			if err != nil {
				return 0, fmt.Errorf("Error when converting the string number into an integer")
			}
			if number < max {
				continue
			}
			max = number
		}
		result += max
	}
	return result, nil
}

func findBiggestNumber(numbers []uint64) uint64 {
	var max uint64
	for _, n := range numbers {
		if n < max {
			continue
		}
		max = n
	}
	return max
}
