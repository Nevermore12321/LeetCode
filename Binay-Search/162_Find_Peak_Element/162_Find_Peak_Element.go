func findPeakElement(nums []int) int {
	var recursionBinarySearch func(int, int) int
	recursionBinarySearch = func(l, r int) int {
		if l >= r {
			return l
		}
		mid := l + (r-l)/2
		// 始终朝着 mid 的值 增长的方向去找
		// 如果 右边 递增，那么 二分朝右找
		if nums[mid] < nums[mid+1] {
			return recursionBinarySearch(mid+1, r)
		} else { // 如果 左边 递增，那么 二分朝左找，注意，这里不能取 mid-1 边界，因为 mid 处的值有可能就是波峰
			return recursionBinarySearch(l, mid)
		}
	}
	return recursionBinarySearch(0, len(nums)-1)
}
