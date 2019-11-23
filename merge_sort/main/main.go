package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{3, 2, 1, 20, 5, 6, 42, -2, 4, 3}
	result := mergeSort(array)
	fmt.Println(result)
}

func mergeSort(subArray []int) []int {
	if len(subArray) == 1 {
		return subArray
	}

	if len(subArray) == 2 {
		return sort(subArray)
	}

	mid := int(math.Floor(float64(len(subArray)) / 2))

	array1 := mergeSort(subArray[:mid])
	array2 := mergeSort(subArray[mid:])

	merge := merge(array1, array2)
	fmt.Printf("sorted merge: %v\n", merge)

	return merge
}

func merge(low, high []int) []int {
	sorted := make([]int, len(low)+len(high))
	lowIndex := 0
	highIndex := 0

	for i := range sorted {
		if lowIndex > len(low)-1 {
			sorted[i] = high[highIndex]
			highIndex++
			continue
		} else if highIndex > len(high)-1 {
			sorted[i] = low[lowIndex]
			lowIndex++
			continue
		}

		if low[lowIndex] < high[highIndex] {
			sorted[i] = low[lowIndex]
			lowIndex++
		} else {
			sorted[i] = high[highIndex]
			highIndex++
		}
	}

	return sorted
}

func sort(subArray []int) []int {
	if subArray[1] < subArray[0] {
		swap(subArray, 0, 1)
	}

	return subArray
}

func swap(array []int, first int, second int) {
	if array[first] == array[second] {
		return
	}

	swap := array[first]
	array[first] = array[second]
	array[second] = swap
}
