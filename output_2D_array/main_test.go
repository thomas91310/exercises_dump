package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDummyCase(t *testing.T) {
	var (
		M = 1
		N = 1
	)
	out := makeArray(M, N)
	expected := [][]int{
		{1},
	}
	assertEqualArrays(t, out, expected)
}

func TestBasicCase(t *testing.T) {
	var (
		M = 4
		N = 4
	)
	out := makeArray(M, N)
	expected := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}
	assertEqualArrays(t, out, expected)
}

func TestBiggerValuesCase(t *testing.T) {
	var (
		M = 10
		N = 4
	)
	out := makeArray(M, N)
	expected := [][]int{
		{1, 2, 3, 4},
		{24, 25, 26, 5},
		{23, 40, 27, 6},
		{22, 39, 28, 7},
		{21, 38, 29, 8},
		{20, 37, 30, 9},
		{19, 36, 31, 10},
		{18, 35, 32, 11},
		{17, 34, 33, 12},
		{16, 15, 14, 13},
	}
	assertEqualArrays(t, out, expected)
}

func assertEqualArrays(t *testing.T, out [][]int, expected [][]int) {
	nRows := len(out)
	nExpectedRows := len(expected)
	require.Equal(t, nRows, nExpectedRows, "The number of rows of these arrays should be the same")
	for i := 0; i < nExpectedRows; i++ {
		nCols := len(out[i])
		nExpectedCols := len(expected[i])
		require.Equal(t, nCols, nExpectedCols, "The number of columns of these arrays should be the same")
		for j := 0; j < nExpectedCols; j++ {
			require.Equal(t, out[i][j], expected[i][j], "The array values don't match the expected ones")
		}
	}
}
