func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// 从 右上角开始搜索
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {		//	target 大于当前元素，向下移动，去掉当前行
			x += 1
		} else {								//	target 小于当前元素，向左移动，去掉当前列
			y -= 1
		}
	}
	return false

}
