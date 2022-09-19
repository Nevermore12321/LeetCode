func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	res := 0

	// 第一重循环 边a
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue
		}
		// 第一重循环 边b
		for j := i + 1; j < n; j++ {
			// 第一重循环 边c
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j] > nums[k] { // 满足两边和大于第三表的条件，结果+1
					res += 1
				}
			}
		}
	}
	return res
}
