package QuickSort

import "Algorithm_and_Data_Structure/common"

/*
	快速排序，原理是在于分，找到元素v所在的位置 p，然后根据p的位置分成左右两部分，分别是 <=v , v , >v
	伪代码：
	QuickSort(arr, l, r) {
		if l >= r {
			return
		}

		p := partition(arr, l, r)

		// 对 [l,p-1] 排序
		QuickSort(arr, l, p - 1)
		// 对 [p+1, r] 排序
		QuickSort(arr, p + 1, r)

	}
*/

// partition 过程，循环不变量，arr[l+1 ... j] <= v, arr[j+1 ... i-1] > v，其中 j 表示 v 的位置变化，i表示前进指针
func partition[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T, l int, r int) int {
	var (
		v T   = arr[l]
		j int = l
	)

	//arr[l+1 ... j] <= v, arr[j+1 ... i-1] > v
	for i := l + 1; i <= r; i++ {
		if arr[i] <= v {
			j += 1
			common.Swap[T](arr, j, i)
		}
	}
	common.Swap[T](arr, l, j)
	return j
}

func Sort[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T) {
	sort[T](arr, 0, len(arr)-1)
}

func sort[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T, l, r int) {
	if l >= r {
		return
	}

	position := partition[T](arr, l, r)

	sort[T](arr, l, position-1)
	sort[T](arr, position+1, r)

}
