package main

import (
	"fmt"
	"sortingAlgorithms/helper"
)

func main() {

	slice := helper.GenerateSlice(30)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	selectionSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

func selectionSort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}
