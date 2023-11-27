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

	n := 5000000
	data1 := ag.GenerteRandomArray(n, n)
	data2 := make([]int, len(data1))
	copy(data2, data1)
	SortHelper.SortTest("MergeSort", data1)
	SortHelper.SortTest("MergeSortAdvance1", data2)

	data3 := ag.GenerteRandomArray(n, n)
	data4 := make([]int, len(data3))
	copy(data4, data3)
	SortHelper.SortTest("MergeSort", data3)
	SortHelper.SortTest("MergeSortAdvance2", data4)

	data5 := ag.GenerteRandomArray(n, n)
	data6 := make([]int, len(data5))
	copy(data6, data5)
	SortHelper.SortTest("MergeSort", data5)
	SortHelper.SortTest("MergeSortAdvance3", data6)

	data7 := ag.GenerteRandomArray(n, n)
	data8 := make([]int, len(data7))
	copy(data8, data7)
	SortHelper.SortTest("MergeSort", data7)
	SortHelper.SortTest("MergeSortBottomUp", data8)
}
