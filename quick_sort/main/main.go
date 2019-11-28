package main

import (
	"fmt"
)

func main() {
	array := []int{2, 1, 20, 5, 6, 42, -2, 4, 3, 5, 3, 25, 122, 4, -4, 300, 73, -10}
	quickSort(array)
	fmt.Println(array)
}

func quickSort(subArray []int) {
	if len(subArray) <= 1 {
		return
	}

	pivot := partition(subArray)

	quickSort(subArray[0:pivot])
	quickSort(subArray[pivot:len(subArray)])
}

func partition(subArray []int) int {
	q, j := 0, 0
	r := len(subArray) - 1
	pivotValue := subArray[r]

	for ; j < r; j++ {
		if subArray[j] <= pivotValue {
			swap(subArray, j, q)
			q++
		}
	}

	swap(subArray, q, r)
	return q
}

func sort(array []int) {
	if array[1] < array[0] {
		swap(array, 0, 1)
	}
}

func swap(array []int, first int, second int) {
	if array[first] == array[second] {
		return
	}
	array[first], array[second] = array[second], array[first]
}
