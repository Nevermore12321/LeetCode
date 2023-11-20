package MergeSortBottomUp

import (
	"math"
)

/*
自底向上的归并排序
之前讨论的都是自顶向下的归并排序，也是就是 一个数组，自顶向下先分解到最小集，然后开始合并
自底向上意思是，将数组直接分解成一个一个最小的元素，直接开始合并
*/

// 合并两个有序的区间 arr[l, mid-1] 和 arr[mid, r]
func merge[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T, l, mid, r int, tmp []T) {
	// 初始化 新数组的 下标
	var (
		i, j int = l, mid
	)
	// 深拷贝一个新数组
	for index := l; index <= r; index++ {
		tmp[index] = arr[index]
	}

	// 合并
	for k := l; k <= r; k++ {
		if i >= mid { // 左边已经 merge 完毕，只剩右边
			arr[k] = tmp[j]
			j += 1
		} else if j > r { // 右边已经 merge 完毕，只剩左边
			arr[k] = tmp[i]
			i += 1
		} else if tmp[i] < tmp[j] { // 两边都没有merge完，左边更小
			arr[k] = tmp[i]
			i += 1
		} else { // 两边都没有merge完，右边更小
			arr[k] = tmp[j]
			j += 1
		}
	}
}

// 归并排序算法
func Sort[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T) {
	temp := make([]T, len(arr))

	length := len(arr)

	// 遍历合并的区间长度，每一轮合并区间长度为 sz 的数组
	for sz := 1; sz < length; sz *= 2 {
		// 遍历要合并的两个区间的起始位置
		// 即合并 [i, i+sz-1],[i+sz, min(i+sz+sz-1, length-1)]
		for i := 0; i+sz < length; i += 2 * sz {
			// r 位置有可能已经不足 i+sz+sz-1 了，因此有可能 r 为 length-1
			if arr[i+sz-1] > arr[i+sz] {
				merge(arr, i, i+sz, int(math.Min(float64(i+sz+sz-1), float64(length-1))), temp)
			}
		}
	}
}
