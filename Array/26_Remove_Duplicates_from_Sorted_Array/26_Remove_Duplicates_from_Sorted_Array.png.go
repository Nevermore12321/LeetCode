func removeDuplicates(nums []int) int {
	var pos int = 0
	for i := pos + 1; i < len(nums); i++ {
		if nums[i] != nums[pos] {
			pos += 1
			// 细节，如果时 1，2，3，4，5 就不需要赋值
			if i-pos > 0 {
				nums[pos] = nums[i]

			}
		}
	}
	// pos 表示最后一个不同的位置，但是需要返回长度，所以 +1
	return pos + 1
}
