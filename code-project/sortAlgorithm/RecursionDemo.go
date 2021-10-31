package sortAlgorithm

import "fmt"

// GetMax 获取数组中的最大值
func GetMax(nums []int) int {
	return process(nums, 0, len(nums) - 1)
}

// 返回 L~R 上的最大值
func process(nums []int, L, R int) int {
	// 递归终止条件
	if L == R {
		return nums[L]
	}

	// 找出 middle 值
	// 右移 相当于 除以 2
	mid := L + ((R - L) >> 2)

	var left = process(nums, L, mid)
	var right = process(nums, mid + 1, R)
	if left < right {
		return right
	} else {
		return left
	}
}

func RecursionTest() {
	nums := []int {1, 4, 2, 4, 5, 7, 4, 32, 2, 6, 7, 4, 2, 6, 3, 6, 7}
	fmt.Println(GetMax(nums))
}