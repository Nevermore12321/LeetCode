func twoSum(nums []int, target int) []int {
	var (
		i int = 0
		j int = len(nums) - 1
	)
	for i < j {
		// 如果 i 和 j 位置的元素相加为 target 直接返回
		if nums[i]+nums[j] == target {
			return []int{i, j}
		}
		//  从 j~j 中间找
		for index := i + 1; index < j; index++ {
			left := target - nums[index]
			if left == nums[i] {
				return []int{i, index}
			} else if left == nums[j] {
				return []int{index, j}
			}
		}

		//  如果没找到，缩小 i-j
		i++
		j--
	}
	return nil
}
