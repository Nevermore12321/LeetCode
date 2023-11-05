package Recursion

func Sum(arr []int) int {
	return sum(arr, 0)
}

// 计算 arr[l...n) 这个区间内所有数字的和
func sum(arr []int, left int) int {
	// 求解最基本问题
	if left == len(arr) {
		return 0
	}

	// 把原问题转换成更小的问题
	return arr[left] + sum(arr, left+1)
}
