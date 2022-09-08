func maxDistance(nums1 []int, nums2 []int) int {
	// 初始化最大下标索引
	maxLen := 0
	var binarySearch func([]int, int, int, int) int
	// 二分搜索
	binarySearch = func(nums []int, l, r, target int) int {
		if l > r { // 注意这里返回 l，也就是第一个 比 target 大的索引
			return l
		}
		mid := l + (r-l)/2
		if nums[mid] >= target { //  如果 nums[mid] 大于等于 target，往右边找，注意这里 等于时继续找第一个大于 target 的索引
			return binarySearch(nums, mid+1, r, target)
		} else { //  如果 nums[mid] 小于 target，往左边找
			return binarySearch(nums, l, mid-1, target)
		}
	}

	// 第一轮遍历 nums1 数组
	for i, _ := range nums1 {
		// 如果 nums1[i] 直接比 nums2[i] 大，直接继续下一个
		if i >= len(nums2) || nums1[i] > nums2[i] {
			continue
		}

		// 找到第一个大的索引
		j := binarySearch(nums2, i, len(nums2)-1, nums1[i])
		curLen := j - 1 - i
		if curLen > maxLen { // 更新 最大距离
			maxLen = curLen
		}
	}
	return maxLen
}
