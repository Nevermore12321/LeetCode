package sortAlgorithm

import (
	"program-algorithm/lib"
)

func HeapSort1(arr []int) {
	// 使用 最大堆优先队列 的构造函数创建堆
	maxHeap := lib.MaxHeapInit(len(arr))

	// 将待排序数组的所有元素，逐个插入到堆中
	for _, item := range arr {
		maxHeap.Insert(item)
	}

	// 同样逐个删除堆中的最大值，逆序放入到待排序数组中
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = maxHeap.DelMax()
	}
}

func HeapSortTest() {

	lib.SortDuration(HeapSort1, "HeapSort1")
}
