package main

import (
	"fmt"
	"sortingAlgorithms/helper"
)

func main() {

	slice := helper.GenerateSlice(30)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubbleSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

func bubbleSort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
