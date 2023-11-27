package main

import (
	"Algorithm_and_Data_Structure/common"
	"Algorithm_and_Data_Structure/common/SortHelper"
)

func main() {
	dataSize := []int{100000, 1000000}
	ag := common.ArrayGenerator{}
	for _, d := range dataSize {
		data := ag.GenerteRandomArray(d, d)
		SortHelper.SortTest("QuickSort", data)
	}

	n := 5000000
	data1 := ag.GenerteRandomArray(n, n)
	data2 := make([]int, len(data1))
	copy(data2, data1)
	SortHelper.SortTest("MergeSort", data1)
	SortHelper.SortTest("QuickSort", data2)

}
