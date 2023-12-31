package main

import (
	"fmt"
	"math/rand"
	"sortingAlgorithms/helper"
)

func main() {

	slice := helper.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quickSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
