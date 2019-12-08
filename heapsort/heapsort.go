package main

import "fmt"

func max_heapify(array []int, i, array_len int) {
	max := i
	right := 2*i + 1
	left := 2*i + 2

	if right <= array_len && array[right] > array[max] {
		max = right
	}
	if left <= array_len && array[left] > array[max] {
		max = left
	}
	if max != i {
		array[i], array[max] = array[max], array[i]
		max_heapify(array, max, array_len)
	}
}

func create_max_heap(array []int) {
	for i := len(array) - 1; i >= 0; i-- {
		max_heapify(array, i, len(array) - 1)
	}
}

func heapsort(array []int) {
	array_len := len(array)

	create_max_heap(array)
	for i := array_len - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		max_heapify(array, 0, i - 1)
	}
}

func main() {
	array := []int{3, 19, 1, 14, 8, 7}

	heapsort(array)
	fmt.Println(array)

	array = []int{3, 1, 7, -1, 12, -6, 111, 32}
	heapsort(array)
	fmt.Println(array)

	array = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	heapsort(array)
	fmt.Println(array)
}
