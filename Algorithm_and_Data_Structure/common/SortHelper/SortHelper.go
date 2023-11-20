package SortHelper

import (
	"Algorithm_and_Data_Structure/Ch2_BaseSort/InsertionSort"
	"Algorithm_and_Data_Structure/Ch2_BaseSort/SelectionSort"
	"Algorithm_and_Data_Structure/Ch6_MergeSearch/MergeSort"
	"Algorithm_and_Data_Structure/Ch6_MergeSearch/MergeSortAdvance"
	"Algorithm_and_Data_Structure/Ch6_MergeSearch/MergeSortAdvance2"
	"Algorithm_and_Data_Structure/Ch6_MergeSearch/MergeSortAdvance3"
	"Algorithm_and_Data_Structure/Ch6_MergeSearch/MergeSortBottomUp"
	"fmt"
	"time"
)

func IsSorted[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func SortTest[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](sortName string, arr []T) {

	startTime := time.Now()
	switch sortName {
	case "SelectionSort":
		SelectionSort.Sort(arr)
	case "InsertionSort":
		InsertionSort.Sort(arr)
	case "InsertionSortAdvance":
		InsertionSort.SortAdvance(arr)
	case "MergeSort":
		MergeSort.Sort(arr)
	case "MergeSortAdvance1":
		MergeSortAdvance.Sort1(arr)
	case "MergeSortAdvance2":
		MergeSortAdvance2.Sort(arr)
	case "MergeSortAdvance3":
		MergeSortAdvance3.Sort(arr)
	case "MergeSortBottomUp":
		MergeSortBottomUp.Sort(arr)
	}
	duration := time.Since(startTime)
	fmt.Printf("%s, %d elements, use: %v\n", sortName, len(arr), duration)
	fmt.Printf("Is Sorted? %v\n\n", IsSorted(arr))
}
