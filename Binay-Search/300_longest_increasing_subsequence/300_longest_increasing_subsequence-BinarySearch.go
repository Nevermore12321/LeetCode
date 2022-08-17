func lengthOfLIS(nums []int) int {
	//  base case 长度为0，则返回0
	if len(nums) == 0 {
		return 0
	}
	// 贪心的 tails 数组，存放当前位置结尾的，长度最长的，值最小子序列
	var tails = make([]int, 1)
	maxLen := 1
	tails[0] = nums[0]

	// 类似动态规划，遍历最外层
	for i := 1; i < len(nums); i++ {
		// 如果 当前位置 i 的元素，比 tails 的最后一个元素还大，直接放到结尾，形成一个更长的子序列
		if nums[i] > tails[maxLen-1] {
			tails = append(tails, nums[i])
			maxLen += 1
		} else {				// 否则使用二分，找到第一个比 nums[i] 的位置，替换
			l, r := 0, maxLen-1
			for l <= r {
				mid := l + (r-l)/2
				if nums[i] > tails[mid] {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			tails[l] = nums[i]
		}
	}
	return maxLen
}
