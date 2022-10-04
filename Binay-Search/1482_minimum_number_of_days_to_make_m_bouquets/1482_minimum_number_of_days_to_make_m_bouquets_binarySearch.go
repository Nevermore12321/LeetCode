func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}

	// 判断 x  天是否可以制作 m 束花
	var check func(int) bool
	check = func(x int) bool {
		// 总共可以使用的花束数
		total := 0
		// 连续相邻的花朵数
		flowers := 0
		for _, v := range bloomDay {
			if v <= x { // 如果 当前花朵成长的天数小于所需天数 x，满足条件，相邻的花朵加1
				flowers += 1
				if flowers == k { //  如果连续的花朵数 等于 k，可以制作成一束花，total 花束数加1
					total += 1
					flowers = 0
					if total == m {
						break
					}
				}
			} else {
				flowers = 0
			}
		}
		// 最多可以制作的花束数如果 大于 m，说明在 x 天内可以制作 m 束花
		return total >= m
	}

	// 找出数组中最大的元素
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

	// 二分的界限 [1, bloomDay 中最大值]
	l, r := 1, maxValue(bloomDay)
	for l <= r {
		mid := l + (r-l)/2
		if check(mid) { // 如果 mid 天可以指制作 m 束花，那么继续缩短天数
			r = mid - 1
		} else { // 如果 mid 天不可以指制作 m 束花，那么继续增大天数
			l = mid + 1
		}
	}

	// 返回第一个可以制作的天数
	return l
}
