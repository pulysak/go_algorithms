package main

import "fmt"

func max_heapify(array []int, i int) {
	max := i
	right := 2*i + 1
	left := 2*i + 2
	array_len := len(array) - 1

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

func create_max_heap(array []int) {
	for i := len(array) - 1; i >= 0; i-- {
		max_heapify(array, i)
	}
}

func main() {
	array := []int{3, 19, 1, 14, 8, 7}
	fmt.Println(array)

	create_max_heap(array)
	fmt.Println(array)
}
