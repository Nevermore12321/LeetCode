func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	// 每一行都需要查找
	for x := 0; x < m; x++ {
		l, r := 0, n-1
		for l <= r { // 二分法
			mid := l + (r-l)/2
			if matrix[x][mid] == target {
				return true
			} else if matrix[x][mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
