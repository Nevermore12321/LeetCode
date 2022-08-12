func missingNumber(nums []int) int {
	n := len(nums)
	// 使用数组代表 hash map，使用 map 一样
	var hashArray = make([]bool, n+1)
	var missing int
	// 出现的元素在 hashArray 中设置为 true
	for _, v := range nums {
		hashArray[v] = true
	}

	// 找到为 false 的值返回
	for i, _ := range hashArray {
		if !hashArray[i] {
			missing = i
			break
		}
	}
	return missing
}
