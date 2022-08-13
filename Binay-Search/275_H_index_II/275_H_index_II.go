func hIndex(citations []int) int {
	MINVALUE := -1
	// 存储最大的 H 指数，初始化为 -1
	maxHIndex, n := MINVALUE, len(citations)-1
	l, r := 0, n
	max := func(a int, b int) int {
		if a < b {
			return b
		}
		return a
	}

	// 使用二分法，
	for l <= r {
		mid := l + (r-l)/2
		// 如果 citations[mid] 大于 mid ~ n 的所有元素个数 ，那么 mid 后面的所有元素都大于，因此 满足 h 指数
		if n-mid+1 <= citations[mid] {
			maxHIndex = max(n-mid+1, maxHIndex)
			r = mid - 1			//	继续寻找更大的 h，向右边寻找
		} else {
			l = mid + 1			// 否则 向左边寻找
		}
	}
	if maxHIndex == MINVALUE {
		return 0
	}
	return maxHIndex
}
