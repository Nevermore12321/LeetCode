package InsertionSort

/*
插入排序特性：对于一个完全有序的数组，时间复杂度为 O(n)

因此插入排序比选择排序好一些
*/

import "Algorithm_and_Data_Structure/common"

// Sort 插入排序，每次把 i 位置的元素，插入到 i 位置之前该有的位置上
// 循环不变量：arr[0,i) 区间已排序，arr[i,n) 区间未排序
// 函数：把 arr[i] 到 0-i 区间该有的位置上
func Sort[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T) {
	for i := 0; i < len(arr); i++ {
		// arr[i] 插入到合适的位置上
		for j := i; j-1 >= 0; j-- {
			if arr[j] < arr[j-1] {
				common.Swap(arr, j, j-1)
			} else {
				break
			}
		}
	}
}

// SortAdvance 插入排序优化，上一个版本在找 i 到合适位置时，每次都交换，一个交换两次赋值
// 优化内容：将交换操作，改为后移，减少赋值操作
func SortAdvance[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T) {
	for i := 0; i < len(arr); i++ {
		// 临时变量存放 arr[j]，每次比较，如果比前一个小，后移
		var (
			j   int = 0
			tmp T   = arr[i]
		)
		for j = i; j-1 >= 0; j-- {
			if tmp < arr[j-1] {
				arr[j] = arr[j-1] // 后移
			} else {
				break
			}
		}
		arr[j] = tmp

	}
}
