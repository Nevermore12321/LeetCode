func minSubArrayLen(target int, nums []int) int {
	i := 0
	sum := 0
	minWidth := math.MaxInt
	minInt := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// i,j 双指针，j 表示右指针，每次向右滑动
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		// 如果 i，j 之间的元素和 大于等于 target，那么这个 子序列满足要求，记录长度，每次取最小的长度
		// 缩小 i，j 区间，i 向右移动，缩小范围
		for sum >= target {
			minWidth = minInt(j-i+1, minWidth)
			sum -= nums[i]
			i++
		}
	}

	if math.MaxInt == minWidth {
		return 0
	}
	return minWidth
}
