func specialArray(nums []int) int {
	l, r := 1, len(nums)
	for l <= r {
		mid := l + (r-l)/2
		count := 0
		//  计算 arr 中 大于等于 mid 的元素个数
		for i := 0; i < len(nums); i++ {
			if nums[i] >= mid {
				count += 1
			}
		}
		if count == mid { 				// 如果 count == mid 找到直接返回
			return mid
		} else if count > mid {			// 如果 count > mid ,说明 arr 中比x大的元素多，扩大 x
			l = mid + 1
		} else {						// 如果 count < mid ,说明 arr 中比x大的元素少，缩小 x
			r = mid - 1
		}
	}
	return -1
}
