// https://www.geeksforgeeks.org/kth-smallestlargest-element-unsorted-array/
// Method 4 (QuickSelect)

package kth_smallestlargest_element_unsorted_array

func Swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func Partition(arr []int, lo int, hi int) int {
	pivotInd := hi
	pivot := arr[hi]

	j := lo
	for i := lo; i < pivotInd; i++ {
		if arr[i] <= pivot {
			Swap(arr, j, i)
			j++
		}
	}

	Swap(arr, j, pivotInd)
	j++

	return j
}
