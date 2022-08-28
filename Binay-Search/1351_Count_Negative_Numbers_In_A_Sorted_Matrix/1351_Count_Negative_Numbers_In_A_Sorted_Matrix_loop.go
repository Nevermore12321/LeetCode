func countNegatives(grid [][]int) int {
	// 暴力搜索
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] < 0 {
				count += 1
			}
		}
	}
	return count
}
