package main

import "fmt"

func main() {
	array := []int{3, 2, 1, 20, 5, 6, 42, -2, 4, 3}
	sort(array)
	fmt.Println(array)
}

func sort(array []int) {
	sorted := false

	for !sorted {
		withoutSwap := true

		for i := 0; i < len(array)-1; i++ {

			if array[i] > array[i+1] {
				swap(array, i, i+1)
				withoutSwap = false
			}
		}

		if withoutSwap == true {
			sorted = true
		}
	}
}

func swap(array []int, first int, second int) {
	array[first], array[second] = array[second], array[first]
}
