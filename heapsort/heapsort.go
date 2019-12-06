package main

import "fmt"

func max_heapify(array []int, i int) {
	max := i
	right := 2*i + 1
	left := 2*i + 2
	array_len := len(array)

	if right <= array_len && array[right] > array[max] {
		max = right
	}
	if left <= array_len && array[left] > array[max] {
		max = left
	}
	if max != i {
		array[i], array[max] = array[max], array[i]
		max_heapify(array, max)
	}
}

func main() {
	array := []int{100, 1, 6, 5, 3, 5, 1, 3, 2, 0, 1}
	fmt.Println(array)

	max_heapify(array, 1)
	fmt.Println(array)
}
