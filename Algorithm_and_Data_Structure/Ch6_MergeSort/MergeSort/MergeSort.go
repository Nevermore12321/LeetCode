package MergeSort

/*
归并排序，伪代码：
MergeSort(arr, l, r) {

	// 基础情况
	if l >= r {
		return
	}

	mid := (r - l + 1) / 2

	// 对 arr[l, mid-1] 进行排序
	MergeSort(arr, l, mid-1)

	// 对 arr[mid, r] 进行排序
	MergeSort(arr, mid, r)

	// 合并 arr[l, mid-1] 和 arr[mid, r]
	merge(arr, l, mid, r);
}


归并排序的重要在于 合并
*/

// 合并两个有序的区间 arr[l, mid-1] 和 arr[mid, r]
func merge[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T, l, mid, r int) {
	// 初始化 新数组的 下标
	var (
		i, j int = l, mid
	)
	// 深拷贝一个新数组
	newArr := make([]T, r-l+1)
	for index := l; index <= r; index++ {
		newArr[index-l] = arr[index]
	}

	// 合并
	for k := l; k <= r; k++ {
		if i >= mid { // 左边已经 merge 完毕，只剩右边
			arr[k] = newArr[j-l]
			j += 1
		} else if j > r { // 右边已经 merge 完毕，只剩左边
			arr[k] = newArr[i-l]
			i += 1
		} else if newArr[i-l] < newArr[j-l] { // 两边都没有merge完，左边更小
			arr[k] = newArr[i-l]
			i += 1
		} else { // 两边都没有merge完，右边更小
			arr[k] = newArr[j-l]
			j += 1
		}
	}
}

// 归并排序算法
func Sort[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T) {
	sort[T](arr, 0, len(arr)-1)
}

// 归并排序的 递归部分，带有下标
func sort[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T, l, r int) {
	// 基础
	if l >= r {
		return
	}

	// mid 表示中间元素（奇数个），中间靠右元素（偶数个）
	mid := l + (r-l+1)/2

	// 递归左右部分
	sort[T](arr, l, mid-1)
	sort[T](arr, mid, r)

	// 合并
	merge[T](arr, l, mid, r)
}
