func missingNumber(nums []int) int {
	res := 0
	// i 相当于 [0,n-1] 的元素
	// v 相当于 nums 的每个元素
	// 全部异或
	for i, v := range nums {
		res ^= i ^ v
	}
	// 这里需要注意，因为上面是 [0, n-1]，还需要异或 n
	return res ^ len(nums)
}
