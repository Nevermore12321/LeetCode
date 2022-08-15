func findDuplicate(nums []int) int {
	// step1 找到快慢指针相遇的地方
	slow, fast := 0, 0
	slow = nums[slow]
	fast = nums[nums[fast]]
	for nums[slow] != nums[fast] {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// 找到出现环的第一个位置
	finder := 0
	for nums[finder] != nums[slow] {
		finder = nums[finder]
		slow = nums[slow]
	}
	return nums[slow]
}
