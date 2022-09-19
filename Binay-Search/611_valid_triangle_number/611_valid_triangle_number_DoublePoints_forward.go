func triangleNumber(nums []int) int {
	// 排序
	sort.Ints(nums)
	n := len(nums)
	res := 0

	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 第一重循环，边 a
	for i := 0; i < n; i++ {
		k := i + 1						// 双指针 k 初始化 i + 1 也可以初始化 i，表示 边 c
		for j := i + 1; j < n; j++ {	// 双指针 j 初始化 i + 1，也可以初始化 i， 表示 边 b
			for k < n-1 && nums[i]+nums[j] > nums[k+1] {		// a + b > c 满足条件，继续向右找，直到找到第一个不满足条件的 k
				k++
			}
			res += maxInt(k-j, 0)				// 此时 [j, k) 之间都满足条件，记录三元组个数
		}
	}
	return res
}
