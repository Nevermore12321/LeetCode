package main

import (
	"Algorithm_and_Data_Structure/Ch2_BaseSort/InsertionSort"
	"Algorithm_and_Data_Structure/Ch2_BaseSort/SelectionSort"
	"Algorithm_and_Data_Structure/Ch2_BaseSort/Student"
	"Algorithm_and_Data_Structure/common"
	"Algorithm_and_Data_Structure/common/SortHelper"
	"fmt"
)

func main() {

	/* ====================== SelectionSort ============================== */
	fmt.Println("====================== SelectionSort ==============================")
	var arr = []int{4, 3, 5, 8, 1, 7, 2, 9, 6}
	SelectionSort.Sort(arr)
	fmt.Println(arr)

	fmt.Println()

	students := []Student.Student{
		Student.Student{Name: "Alice", Score: 92},
		Student.Student{Name: "Smith", Score: 56},
		Student.Student{Name: "Jackson", Score: 98},
		Student.Student{Name: "Jack", Score: 82},
		Student.Student{Name: "Bob", Score: 66},
	}

	SelectionSort.CustomSort[Student.Student](students)
	fmt.Println(students)

	dataSize := []int{10000, 100000}
	ag := common.ArrayGenerator{}
	for _, d := range dataSize {
		data := ag.GenerteRandomArray(d, d)
		SortHelper.SortTest("SelectionSort", data)
	}

	/* ====================== InsertionSort ============================== */
	fmt.Println("====================== InsertionSort ==============================")
	var arr1 = []int{4, 3, 5, 8, 1, 7, 2, 9, 6}
	InsertionSort.Sort(arr1)
	fmt.Println(arr1)

	fmt.Println()

	for _, d := range dataSize {
		data := ag.GenerteRandomArray(d, d)
		SortHelper.SortTest("InsertionSort", data)
	}

	var arr2 = []int{4, 3, 5, 8, 1, 7, 2, 9, 6}
	InsertionSort.SortAdvance(arr2)

	fmt.Println("InsertionSort & InsertionSortAdvance")
	for _, d := range dataSize {
		data := ag.GenerteRandomArray(d, d)
		data2 := make([]int, len(data))
		copy(data2, data)
		SortHelper.SortTest("InsertionSort", data)
		SortHelper.SortTest("InsertionSortAdvance", data2)
	}

}
