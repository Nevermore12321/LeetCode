func chalkReplacer(chalk []int, k int) int {
	// 注意：前缀和从 第二个元素开始累加，这里要判断第一个元素是不是已经超过 k 了
	if chalk[0] > k {
		return 0
	}
	n := len(chalk)

	// 计算前缀和，同样是判断第一轮中有没有累加超过 k 的元素
	// 这里循环结束会将 chalk 数组，对应元素下标 i 变成 [0.i] 的元素累加和，也就是 前缀和
	for i := 1; i < n; i++ {
		chalk[i] += chalk[i-1]
		if chalk[i] > k {
			return i
		}
	}

	// 计算最后一轮中的 k’
	restK := k % chalk[n-1]

	// 使用二分搜索找到 前缀和数组中，第一个大于 k' 的元素下标
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if chalk[mid] <= restK {	// 这里相等也要继续找，要找到大于 target 的元素下标
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}
