package main

import "fmt"

func main() {
	list := data{-29, 446, 361, 327, 320, 130, 508, -104, 67, 74, 4, -348, -457, 349, -15, 216, 12, -468, 6, 202, 170, 914, -777, 4, -316, -170, -12, -559, -322, -110}
	fmt.Println("\n--- Unsorted --- \n\n", list)

	list.pancakeSort()
	fmt.Println("\n--- Sorted ---\n\n", list, "\n")
}

type data []int32

func (dataList data) pancakeSort() {
	for uns := len(dataList) - 1; uns > 0; uns-- {
		// find largest in unsorted range
		lx, lg := 0, dataList[0]
		for i := 1; i <= uns; i++ {
			if dataList[i] > lg {
				lx, lg = i, dataList[i]
			}
		}
		// move to final position in two flips
		dataList.flip(lx)
		dataList.flip(uns)
	}
}

func (dataList data) flip(r int) {
	for l := 0; l < r; l, r = l+1, r-1 {
		dataList[l], dataList[r] = dataList[r], dataList[l]
	}
}
