package main

func main() {
	var (
		M = 4
		N = 4
	)
	makeArray(M, N)
}

func makeArray(M int, N int) [][]int {
	out := make([][]int, M)
	for i := 0; i < M; i++ {
		row := make([]int, N)
		out[i] = row
	}
	fillArray(out, M, N)
	return out
}

//I noticed that it always does "right", "down", "left", "up", "right", "down" ....
//until we're done fulfilling the array (criterion for stopping is that output = N*M)
func fillArray(out [][]int, M int, N int) {
	var (
		//steps tells you what to do after you reach an edge
		//For example, if you reach an edge on the right side
		//you will need to go down the next iteration
		steps = map[string]string{
			"right": "down",
			"down":  "left",
			"left":  "up",
			"up":    "right",
		}
		i = 0
		j = 0
		//current is in direction of where we're headed
		//can be right, down, left and up (always in that precise order though)
		//starting by right because we're filling the first row first
		current = "right"
		//starting value for value in the 2D array
		output = 1
		run    = true
	)

	for run {
		out[i][j] = output
		if output == N*M {
			run = false
		}
		if current == "right" {
			if j+1 == N || out[i][j+1] != 0 {
				current = steps[current]
				continue
			}
			j++
			output++
		}
		if current == "down" {
			if i+1 == M || out[i+1][j] != 0 {
				current = steps[current]
				continue
			}
			i++
			output++
		}
		if current == "left" {
			if j-1 == -1 || out[i][j-1] != 0 {
				current = steps[current]
				continue
			}
			j--
			output++
		}
		if current == "up" {
			if i-1 == -1 || out[i-1][j] != 0 {
				current = steps[current]
				continue
			}
			i--
			output++
		}
	}
}
