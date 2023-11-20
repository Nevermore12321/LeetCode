package MergeSortAdvance

/*
归并排序的优化1： 在merge过程中，如果 左边最后一个元素已经比 右边第一个元素小，那么可以直接排序，不需要merge过程
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
func Sort1[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T) {
	sort1[T](arr, 0, len(arr)-1)
}

// 归并排序的 递归部分，带有下标
func sort1[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](arr []T, l, r int) {
	// 基础
	if l >= r {
		return
	}

	// mid 表示中间元素（奇数个），中间靠右元素（偶数个）
	mid := l + (r-l+1)/2

	// 递归左右部分
	sort1[T](arr, l, mid-1)
	sort1[T](arr, mid, r)

	// 优化1 ，如果已经有序，不需要merge
	if arr[mid-1] <= arr[mid] {
		return
	}
	// 合并
	merge[T](arr, l, mid, r)
}
