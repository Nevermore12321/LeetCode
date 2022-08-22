func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	count := 0

	// 求 绝对值
	intAbs := func(a int) int {
		if a > 0 {
			return a
		}
		return -a
	}

	// arr1 数组的每一个元素，都去减一遍 arr2
	for _, v1 := range arr1 {
		var flag bool = true
		for _, v2 := range arr2 {
			// 如果 ∣ x − y ∣ <= d , 不满足 距离值定义，直接进行 arr1 的下一个元素
			if intAbs(v2-v1) <= d {
				flag = false
				break
			}
		}
		// 如果 都满足 ∣ x − y ∣ > d, 满足距离值定义，加过 + 1
		if flag {
			count += 1
		}
	}
	return count
}
