func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l <= r {
		mid := l + (r-l)/2
		res := mid * mid
		if res == num {
			return true
		} else if res < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
