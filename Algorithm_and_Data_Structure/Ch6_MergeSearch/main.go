package main

import (
	"Algorithm_and_Data_Structure/common"
	"Algorithm_and_Data_Structure/common/SortHelper"
)

func main() {
	dataSize := []int{10000, 100000}
	ag := common.ArrayGenerator{}
	for _, d := range dataSize {
		data := ag.GenerteRandomArray(d, d)
		SortHelper.SortTest("MergeSort", data)
	}
}
