package kth_smallestlargest_element_unsorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	type Input struct {
		arr []int
		i   int
		j   int
	}
	argsIn := []Input{
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 0, j: 0},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 0, j: 1},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 0, j: 6},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 1, j: 5},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 1, j: 6},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 3, j: 4},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 6, j: 6},
	}
	argsOutExpected := [][]int{
		[]int{8, 1, 3, 7, 4, 2, 9},
		[]int{1, 8, 3, 7, 4, 2, 9},
		[]int{9, 1, 3, 7, 4, 2, 8},
		[]int{8, 2, 3, 7, 4, 1, 9},
		[]int{8, 9, 3, 7, 4, 2, 1},
		[]int{8, 1, 3, 4, 7, 2, 9},
		[]int{8, 1, 3, 7, 4, 2, 9},
	}
	for idx, arg := range argsIn {
		argOutExpected := argsOutExpected[idx]

		Swap(arg.arr, arg.i, arg.j)
		argOutComputed := arg.arr

		assert.Equal(t, argOutExpected, argOutComputed, "Output at idx=%d mismatches\narr=%v\ni=%d\tj=%d", idx, arg.arr, arg.i, arg.j)
	}
}

func TestPartition(t *testing.T) {
	type Input struct {
		arr []int
		i   int
		j   int
	}
	argsIn := []Input{
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 0, j: 0},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 0, j: 1},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 0, j: 3},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 2, j: 3},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 4, j: 6},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 4, j: 4},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 5, j: 6},
		{arr: []int{8, 1, 3, 7, 4, 2, 9}, i: 6, j: 6},
	}
	argsOutExpected := [][]int{
		[]int{8, 1, 3, 7, 4, 2, 9},
		[]int{1, 8, 3, 7, 4, 2, 9},
		[]int{1, 3, 7, 8, 4, 2, 9},
		[]int{8, 1, 3, 7, 4, 2, 9},
		[]int{8, 1, 3, 7, 4, 2, 9},
		[]int{8, 1, 3, 7, 4, 2, 9},
		[]int{8, 1, 3, 7, 4, 2, 9},
		[]int{8, 1, 3, 7, 4, 2, 9},
	}
	for idx, arg := range argsIn {
		argOutExpected := argsOutExpected[idx]

		Partition(arg.arr, arg.i, arg.j)
		argOutComputed := arg.arr

		assert.Equal(t, argOutExpected, argOutComputed, "Output at idx=%d mismatches\narr=%v\ni=%d\tj=%d", idx, arg.arr, arg.i, arg.j)
	}
}
