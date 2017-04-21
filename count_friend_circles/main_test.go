package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid1(t *testing.T) {
	arr := make([][]bool, 6)
	arr[0] = []bool{true, true, false, false, false, false}
	arr[1] = []bool{true, true, false, false, false, false}
	arr[2] = []bool{false, false, true, false, false, false}
	arr[3] = []bool{false, false, false, true, true, false}
	arr[4] = []bool{false, false, false, true, true, true}
	arr[5] = []bool{false, false, false, false, true, true}

	visited := map[int]bool{}
	counter := countCircles(arr, visited)

	assert.Equal(t, counter, 3, "These two numbers should be the same")
}

func TestGrid2(t *testing.T) {
	arr := make([][]bool, 6)
	arr[0] = []bool{true, true, false, false, false, false}
	arr[1] = []bool{true, true, true, false, false, false}
	arr[2] = []bool{false, true, true, true, false, false}
	arr[3] = []bool{false, false, true, true, true, false}
	arr[4] = []bool{false, false, false, true, true, true}
	arr[5] = []bool{false, false, false, false, true, true}

	counter := 0
	visited := map[int]bool{}
	counter = countCircles(arr, visited)

	assert.Equal(t, counter, 1, "These two numbers should be the same")
}

func TestGrid3(t *testing.T) {
	arr := make([][]bool, 5)
	arr[0] = []bool{true, false, false, false, false}
	arr[1] = []bool{false, true, false, false, false}
	arr[2] = []bool{false, false, true, false, false}
	arr[3] = []bool{false, false, false, true, false}
	arr[4] = []bool{false, false, false, false, true}

	visited := map[int]bool{}
	counter := countCircles(arr, visited)

	assert.Equal(t, counter, 5, "These two numbers should be the same")
}
