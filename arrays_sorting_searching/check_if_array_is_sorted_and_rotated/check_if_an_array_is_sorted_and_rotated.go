package check_if_array_is_sorted_and_rotated

import (
	"errors"
	"fmt"
	"math"
)

func FindMinEleIndex(arr []int) int {
	minEle, minEleIndex := math.MaxInt32, math.MinInt32
	for idx, val := range arr {
		if val < minEle {
			minEle = val
			minEleIndex = idx
		}
	}
	return minEleIndex
}

func CheckAscendingOrder(arr []int, lo int, hi int, trueOnError bool) (bool, error) {
	var err error
	if lo > hi {
		// error
		err = errors.New(fmt.Sprintf("Invalid range lo=%d < hi=%d; lo can't be greater than hi", lo, hi))
	} else if (lo < 0) || ((len(arr) > 0) && (len(arr) <= hi)) {
		// error
		err = errors.New(fmt.Sprintf("Invalid range lo=%d, hi=%d for arr[] of len=%d", lo, hi, len(arr)))
	}

	if err != nil {
		// return errors
		if trueOnError {
			return true, err
		} else {
			return false, err
		}
	}

	// happy cases
	if len(arr) > 0 {
		// normal case
		prev := math.MinInt32
		for idx := lo; idx <= hi; idx += 1 {
			if prev > arr[idx] {
				// if non-increasing order found, return immediately
				return false, nil
			} else {
				prev = arr[idx]
			}
		}
		// otherwise it must be ascending order
		return true, nil
	} else {
		// empty array is implicitly ascending order
		return true, nil
	}
}

func IsSortedRotated(arr []int) (bool, int) {
	// find index of smallest element
	//minEleInd := FindMinEleIndex(arr)
	// from here, 3 things should hold
	// 1. elements before minEleInd should be in ascending order
	//elesBeforeMinEleAscending, _ := CheckAscendingOrder(arr, 0, minEleInd-1, true)
	// 2. elements after minEleInd should be in ascending order
	// 3. last element should be greater than the element at minEleInd - 1
	return true, 1
}
