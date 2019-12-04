package main

import "fmt"

func merge(left []int, right []int) []int {
	left_len, right_len := len(left), len(right)
	var result = make([]int, left_len + right_len)
	left_i,right_i, curren_position := 0, 0, 0
	for right_i <= right_len - 1  && left_i <= left_len - 1 {
		if left[left_i] < right[right_i] {
			result[curren_position] = left[left_i]
			left_i++
		} else {
			result[curren_position] = right[right_i]
			right_i++
		}
		curren_position++
	}

	if right_i > right_len - 1 {
		for left_i <= left_len - 1 {
			result[curren_position] = left[left_i]
			left_i++
			curren_position++
		}
	} else {
		for right_i <= right_len - 1 {
			result[curren_position] = right[right_i]
			right_i++
			curren_position++
		}
	}
	return result
}


func mergesort(array []int) []int {
	array_len := len(array)
	if array_len > 1{
		mid := len(array) / 2
		return merge(mergesort(array[:mid]), mergesort(array[mid:]))
	}
	return array
}

func main() {
	a := []int{3, 1, 7, -1, 12, -6, 111, 32, 0}
	fmt.Println(mergesort(a))
}