func triangleNumber(nums []int) int {
	// 排序
	sort.Ints(nums)
	n := len(nums)
	res := 0

	// 第一重循环，边 c，从右往左遍历
	for i := n - 1; i >= 0; i-- {
		j := i - 1			//	双指针 j，从i - 1 ，右边向左遍历，表示 边 b
		k := 0				//	双指针 k，从 0 ，左边向右遍历，表示 边 a
		for k < j {			// 如果 k == j, 遍历结束
			if nums[k]+nums[j] > nums[i] {		// 如果 a + b > c 满足条件，
				res += j - k					// [k, j] 之间的所有元素满足条件，记录
				j--								// 此时换下一个 边 b，j向左移动
			} else {							// 不满足 a +b > c 条件，k 继续向右移动，直到找到（或者停止循环）
				k++
			}
		}
	}
	return res
}
