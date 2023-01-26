package mergesort

import (
	"sort"
)

func MergeSort(arr []float64) {
	ln := len(arr)

	if ln == 1 {
		return
	}

	pivo := ln / 2

	left, right := arr[0:pivo], arr[pivo:]
	MergeSort(left)
	MergeSort(right)

	merge(arr, pivo)
}

// `arr` is assumed to consist of two non-intersecting sorted arrays
func merge(arr []float64, pivo int) {

	var sorted []float64
	if pivo < len(arr)/2 {
		// swaping arrays
		temp := make([]float64, pivo)
		copy(temp, arr[:pivo])
		copy(arr, arr[pivo:])
		pivo = len(arr) - pivo
	}
	sorted = arr[:pivo] // part of `arr` left of `pivo`

	for i := pivo; i < len(arr); i++ {

		// debug
		// if i == 23 {
		// 	fmt.Print("")
		// }

		pos := sort.SearchFloat64s(sorted, arr[i]) - 1

		switch {
		case pos < 0:
			temp := arr[i]          // arr[i] should be leftmost element of `sorted`
			copy(arr[1:], arr[0:i]) // shift right by one
			arr[0] = temp
		case pos == len(sorted)-1: // arr[i] should be rightmost element of `sorted`
			sorted = arr[:i+1]
			continue
		default:
			temp := arr[i]
			copy(arr[pos+2:], arr[pos+1:i])
			arr[pos+1] = temp
		}

		sorted = arr[:i+1]
	}
}
