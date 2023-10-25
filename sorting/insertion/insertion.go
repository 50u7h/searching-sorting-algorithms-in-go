package main

import (
	"fmt"
	"sortingAlgorithms/helper"
)

func main() {

	slice := helper.GenerateSlice(30)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	insertionSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

func insertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
