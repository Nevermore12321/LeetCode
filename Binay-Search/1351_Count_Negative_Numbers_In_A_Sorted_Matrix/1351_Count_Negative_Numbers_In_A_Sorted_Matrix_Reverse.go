func countNegatives(grid [][]int) int {
	// 倒序遍历
	count := 0

	// 从第一行的 最后一元素开始找
	for i, j := 0, len(grid[0])-1; i < len(grid) && j >= 0; {
		// 如果当前元素 大于等于0，说明这一整行都是大于等于 0，直接下一行
		if grid[i][j] >= 0 {
			i += 1
		} else {		// 如果当前元素 小于0，说明当前位置的下面所有元素都是 小于0，统计个数，并且继续往前找，看是否还有 小于0的元素
			count += len(grid) - i
			j -= 1
		}
	}
	return count
}
