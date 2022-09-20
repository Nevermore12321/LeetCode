func findClosestElements(arr []int, k int, x int) []int {
	var binarySearch func(int, int, int) int
	// 二分查找，返回第一个等于 target 或者 第一个 大于 target 的下标
	binarySearch = func(l, r, target int) int {
		if l > r {
			return l
		}

		mid := l + (r-l)/2
		if arr[mid] >= target {
			return binarySearch(l, mid-1, target)
		} else {
			return binarySearch(mid+1, r, target)
		}
	}
	absInt := func(a int) int {
		if a >= 0 {
			return a
		} else {
			return -a
		}
	}

	// 找到 arr 中最近姐 x 的位置
	pos := binarySearch(0, len(arr)-1, x)
	// 双指针 初始化
	i, j := pos, pos

	// 找到 k 个元素就退出
	for cur := 0; cur < k; cur++ {
		left := math.MaxInt64		// 默认初始化为最大值，因为数组有可能越界
		right := math.MaxInt64		// 默认初始化为最大值，因为数组有可能越界

		// 左指针 从 i-1 开始判断， 如果下标不越界，就计算距离
		if i-1 >= 0 {
			left = absInt(arr[i-1] - x)
		}
		// 右指针 从 j 开始判断，如果下标不越界，就计算距离
		if j < len(arr) {
			right = absInt(arr[j] - x)
		}

		// 如果 左指针元素的距离小，或者 左右指针元素的距离相等，取左指针的元素
		if left <= right {
			i--
		} else {	// 否则，取右指针的元素
			j++
		}
	}

	// 取 k 个最小距离的元素，范围为 [i, j), 也就是 i ~ j-1 的元素
	return arr[i:j]
}
