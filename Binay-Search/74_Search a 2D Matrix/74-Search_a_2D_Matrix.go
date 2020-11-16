package BinarySearch



func SearchMatrix(matrix [][]int, target int) bool {

	var (
		m int = len(matrix)
		binarySearch func(int, int) bool
	)
	if m == 0 {
		return false
	}

	var n int = len(matrix[0])

	binarySearch = func(l int, r int) bool {
		if l > r {
			return false
		}
		mid := l + (r - l ) / 2
		row := mid / n
		col := mid % n
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			return binarySearch(mid + 1, r)
		} else {
			return binarySearch(l, mid - 1)
		}
	}

	return binarySearch(0, m * n - 1)
}