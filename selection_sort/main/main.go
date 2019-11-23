package main

import "fmt"

func main() {
	array := []int{3, 2, 1, 20, 5, 6, 42, -2, 4, 3, 19}
	sort(array)

	fmt.Println(array)
}

func sort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		min := i

		for i2 := i; i2 < len(array); i2++ {
			if array[i2] < array[min] {
				min = i2
			}
		}

		fmt.Printf("Swap %d with %d\n", array[i], array[min])
		swap(array, i, min)
	}
}

func swap(array []int, first int, second int) {
	if array[first] == array[second] {
		return
	}

	swap := array[first]
	array[first] = array[second]
	array[second] = swap
}
