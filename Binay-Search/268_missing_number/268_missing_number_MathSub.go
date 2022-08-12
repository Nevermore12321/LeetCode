func missingNumber(nums []int) int {
	n := len(nums)
	// 利用等差数列的求和公式  Sn = n*a1+n(n-1)d/2 = n(a1+an)/2
	// 求出 [0, n] 所有元素和 totalSum
	totalSum := n * (n + 1) / 2
	numsSum := 0
	// 求出数组 nums 的所有元素和 numsSum
	for _, v := range nums {
		numsSum += v
	}
	// 作差
	return totalSum - numsSum
}
