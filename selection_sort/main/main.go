package main

import "fmt"

func main() {
	array := []int{3, 2, 8, 12, 4, 1, 20, 5}

	fmt.Printf("Array before sorting: %v\n", array)

	for i := 1; i < len(array); i++ {
		insert(array, i-1, array[i])
	}

	fmt.Printf("Array after sorting: %v\n", array)
}

func insert(array []int, rightIndex int, value int) {
	for i := rightIndex; array[i] > value; i-- {
		if i == 0 {
			array[1] = array[0]
			array[0] = value
			break
		}

		swapValue := array[i]

		array[i] = value
		array[i+1] = swapValue
	}
}
