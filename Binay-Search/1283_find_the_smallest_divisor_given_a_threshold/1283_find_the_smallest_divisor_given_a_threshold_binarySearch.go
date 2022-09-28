func smallestDivisor(nums []int, threshold int) int {
	// y=f(x) 也就是计算 除数为x时，所有元素除以x向上取整后的和
	var compute_total func(int) int
	compute_total = func(x int) int {
		total := 0
		for _, v := range nums {
			total += int(math.Ceil(float64(v) / float64(x)))
		}
		return total

	}

	// 返回数组中的最大值
	var maxValue func([]int) int
	maxValue = func(arr []int) int {
		maxValue := 0
		for _, v := range arr {
			if v > maxValue {
				maxValue = v
			}
		}
		return maxValue
	}

	// 二分搜索
	l, r := 1, maxValue(nums)
	for l <= r {
		mid := l + (r-l)/2
		if compute_total(mid) <= threshold {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l

}
