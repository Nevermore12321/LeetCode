func arrangeCoins(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)/2
		// 计算 mid 的前 n 项 和 ，也就是 1+2+3+...+mid
		res := mid * (mid + 1) / 2
		if res == n {
			return mid
		} else if res < n {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	// 注意，这里 l 是第一个 1+2+..+l > n 的 列，因此，l 的前一列，是完全堆满的最后一列
	return l - 1
}
