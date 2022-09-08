func maxDistance(nums1 []int, nums2 []int) int {
	// 初始化 双指针
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] { // 如果不满足 要求，i 向前移动
			i += 1
		}
		j += 1
	}
	if j-i-1 > 0 { // 这里减一是因为 j 已经越界了
		return j - i - 1
	} else {
		return 0
	}
}
