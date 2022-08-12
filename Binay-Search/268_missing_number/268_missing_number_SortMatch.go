func missingNumber(nums []int) int {
	// 将数组 nums 排序
	sort.Ints(nums)
	// 找到 nums[i] != i, 返回 i
	for i, v := range nums {
		if v != i {
			return i
		}
	}

	// 如果没找到，那么就说明最后一个元素丢失
	return len(nums)
}
