package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"time"
	"zinoviev_4/mergesort"
)

const (
	COUNT        = 100000
	LENGTH_SHORT = 30
	LENGTH_LONG  = 500
)

var SEED int64 = 1337

func init() {
	flag.Int64Var(&SEED, "s", 1337, "Seed of RNG that will fill the arrays.")
	flag.Parse()
}

func getArr(r *rand.Rand, length int64) (arr []float64, toFind float64) {
	out := make([]float64, length)
	for i := range out {
		out[i] = r.Float64()*200 - 100
	}
	return out, out[r.Int()%len(out)]
}

func findBruteForce(arr []float64, v float64) (index int) {
	for index = range arr {
		if arr[index] == v {
			return
		}
	}
	return -1
}

func findInArr(arr []float64, v float64, method func([]float64, float64) int, mw ...func([]float64)) (searchTime int64) {
	for i := range mw {
		mw[i](arr)
	}
	start := time.Now()
	method(arr, v)
	return time.Since(start).Nanoseconds()
}

func main() {
	r := rand.New(rand.NewSource(SEED))

	fmt.Println("// Non-sorted")
	// Search using bruteforce algorithm for COUNT arrays of lengthes LENGTH_SHORT and LENGTH_LONG
	var bruteForceShort int64
	for i := 0; i < COUNT; i++ {
		arr, toFind := getArr(r, LENGTH_SHORT)
		bruteForceShort += findInArr(arr, toFind, findBruteForce)
	}
	fmt.Printf("findBruteForce(), LENGTH_SHORT: %14f ns/op\n", float64(bruteForceShort)/COUNT)

	var bruteForceLong int64
	for i := 0; i < COUNT; i++ {
		arr, toFind := getArr(r, LENGTH_LONG)
		bruteForceLong += findInArr(arr, toFind, findBruteForce)
	}
	fmt.Printf("findBruteForce(), LENGTH_LONG:  %14f ns/op\n", float64(bruteForceLong)/COUNT)

	fmt.Println("// Sorted")
	// Search in sorted array using bruteforce
	var bruteForceShortSorted int64
	for i := 0; i < COUNT; i++ {
		arr, toFind := getArr(r, LENGTH_SHORT)
		bruteForceShortSorted += findInArr(arr, toFind, findBruteForce, mergesort.MergeSort)
	}
	fmt.Printf("findBruteForce(), LENGTH_SHORT: %14f ns/op\n", float64(bruteForceShort)/COUNT)

	var bruteForceLongSorted int64
	for i := 0; i < COUNT; i++ {
		arr, toFind := getArr(r, LENGTH_LONG)
		bruteForceLongSorted += findInArr(arr, toFind, findBruteForce, mergesort.MergeSort)
	}
	fmt.Printf("findBruteForce(), LENGTH_LONG:  %14f ns/op\n", float64(bruteForceLong)/COUNT)

	// Search in sorted array using binary search from standard library
	var bSearchShort int64
	for i := 0; i < COUNT; i++ {
		arr, toFind := getArr(r, LENGTH_SHORT)
		bSearchShort += findInArr(arr, toFind, sort.SearchFloat64s, mergesort.MergeSort)
	}
	fmt.Printf("searchFloat64s(), LENGTH_SHORT: %14f ns/op\n", float64(bSearchShort)/COUNT)

	var bSearchLong int64
	for i := 0; i < COUNT; i++ {
		arr, toFind := getArr(r, LENGTH_LONG)
		bSearchLong += findInArr(arr, toFind, sort.SearchFloat64s, mergesort.MergeSort)
	}
	fmt.Printf("searchFloat64s(), LENGTH_LONG:  %14f ns/op\n", float64(bSearchLong)/COUNT)
}
