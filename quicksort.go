package main

import "fmt"

func partition(array []int, start, end int) int {

	pivot := array[end]
	var pivotIndex = start

	for i := start; i < end; i++ {
		if array[i] < pivot {
			array[pivotIndex], array[i] = array[i], array[pivotIndex]
			pivotIndex++
		}
	}
	array[pivotIndex], array[end] = pivot, array[pivotIndex]
	return pivotIndex
}

func quicksort(array []int, start, end int) {

	if start < end {
		p := partition(array, start, end)
		quicksort(array, 0, p - 1)
		quicksort(array, p + 1, end)
	}
}

func main() {
	a := []int{2, 2, 2, 2, 2, 2, 2, 2}
	quicksort(a, 0, len(a) - 1)
	fmt.Println(a)

	a = []int{3, 1, 7, -1, 12, -6, 111, 32}
	quicksort(a, 0, len(a) - 1)
	fmt.Println(a)

	a = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	quicksort(a, 0, len(a) - 1)
	fmt.Println(a)
}