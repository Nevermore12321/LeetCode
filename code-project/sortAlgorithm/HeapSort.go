package sortAlgorithm

import (
	"fmt"
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




func HeapSort2(arr []int) {
	//  排序过程
	maxHeap := lib.MaxHeapInitForSort(arr, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = maxHeap.DelMax()
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}


//  堆的 下沉操作
//  arr 表示堆的实现数组，注意堆元素 从 0 开始存放
//  k 表示要下沉的节点
//  n 表示堆的长度 也就是 count
func sink(arr []int, k, n int) {
	// 2k+1 是 k 的左孩子节点
	for 2 * k + 1 < n {
		j := 2 * k + 1

		if j + 1 < n && arr[j] < arr[j + 1] {
			j = j + 1
		}

		if arr[k] > arr[j] {
			break
		}

		swap(arr, k , j)
		k = j
	}
	
}

func HeapSort3(arr []int) {
	// 建堆过程, 把原始数组抽象成堆数组
	n := len(arr)
	for i := (n - 1) / 2; i >= 0; i-- {
		sink(arr, i, n)
	}

	// 排序过程。i 表示每次堆的长度
	// 每次将最大的第一个元素放到堆末尾，然后堆长度减一
	for i := n - 1; i > 0; i-- {
		// 将最大的 0 号根节点元素，与堆长度最后的位置元素交换
		swap(arr, 0, i)
		// 交换后，破坏了堆有序化，利用 sink 操作，从 根节点开始 下沉，重新形成堆
		sink(arr, 0, i)
	}
}


func HeapSortTest() {
	arr := []int{19, 17, 16, 22, 28, 62, 30, 41, 13, 15}
	HeapSort3(arr)
	fmt.Println(arr)
}