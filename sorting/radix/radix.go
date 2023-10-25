package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const digit = 4
const maxbit = -1 << 31

func main() {

	var data = []int32{-29, 446, 361, 327, 320, 130, 508, -104, 67, 74, 4, -348, -457, 349, -15, 216, 12, -468, 6, 202, 170, 914, -777, 4, -316, -170, -12, -559, -322, -110}
	fmt.Println("\n--- Unsorted --- \n\n", data)
	radixSort(data)
	fmt.Println("\n--- Sorted ---\n\n", data, "\n")
}

func radixSort(data []int32) {
	buf := bytes.NewBuffer(nil)
	ds := make([][]byte, len(data))
	for i, e := range data {
		err := binary.Write(buf, binary.LittleEndian, e^maxbit)
		if err != nil {
			return
		}
		b := make([]byte, digit)
		_, err = buf.Read(b)
		if err != nil {
			return
		}
		ds[i] = b
	}
	countingSort := make([][][]byte, 256)
	for i := 0; i < digit; i++ {
		for _, b := range ds {
			countingSort[b[i]] = append(countingSort[b[i]], b)
		}
		j := 0
		for k, bs := range countingSort {
			copy(ds[j:], bs)
			j += len(bs)
			countingSort[k] = bs[:0]
		}
	}
	var w int32
	for i, b := range ds {
		buf.Write(b)
		err := binary.Read(buf, binary.LittleEndian, &w)
		if err != nil {
			return
		}
		data[i] = w ^ maxbit
	}
}
