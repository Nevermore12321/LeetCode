func lengthOfLIS(nums []int) int {
	//  base case 长度为0，则返回0
	if len(nums) == 0 {
		return 0
	}
	//  状态转移方程, dp 每个位置，表示原数组 nums 中对应位置结尾的，最大递增自序列的长度
	var dp = make([]int, len(nums))
	dp[0] = 1
	// 返回结果，最大递增子序列的长度
	maxLen := 1

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	//  动态规划，典型的双层循环，外层遍历 整个nums 数组
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		// 内层，子问题，每个原数组元素，求出一个最大递增子序列的长度
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {				// 求出 i 位置前的最大递增子序列，并且 nums[i] 比子序列的最后一个元素大，才可以合并成一个更大的子序列
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])			// 每个元素的寻找，都更新 maxLen ，如果长度更大，就更新
	}
	return maxLen
}
