func judgeSquareSum(c int) bool {
	// 双指针初始化，从0 到 根号 c
	l, r := 0, int(math.Sqrt(float64(c)))

	// 双指针开始夹逼
	for l <= r {
		res := l*l + r*r
		if res == c {		// 如果等于 c 表示找到
			return true
		} else if res > c {		// 如果 大于 c，表示 右指针过大，右指针向左夹逼
			r -= 1
		} else {				// 否则，左指针过大，左指针向右夹逼
			l += 1
		}
	}
	return false
}
