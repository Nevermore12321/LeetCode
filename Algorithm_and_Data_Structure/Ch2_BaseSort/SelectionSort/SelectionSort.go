package SelectionSort

import "Algorithm_and_Data_Structure/common"

// Sort 选择排序，每次选择一个最小的放入 i 位置
// 循环不动量：arr[0,i) 区间是已排序的
// 函数：arr[i, n)中最小的元素放入 arr[i] 位置
func Sort[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T) {
	for i := 0; i < len(arr); i++ {
		// 找出 arr[i, n) 中最小值的索引
		var minIndex int = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		common.Swap(arr, i, minIndex)
	}
}

// CustomSort 自定义结构体排序，实现了 CompareTo 方法可以比较
func CustomSort[T common.CustomStruct](arr []T) {
	for i := 0; i < len(arr); i++ {
		var minIndex int = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j].CompareTo(arr[minIndex]) < 0 {
				minIndex = j
			}
		}
		common.CustomSwap[T](arr, i, minIndex)
	}
}
