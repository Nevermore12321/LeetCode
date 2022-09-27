func maxDistance(position []int, m int) int {
	// 排序
	sort.Ints(position)

	// y=f(x)
	// 这里如果 最小距离为 x 时，最多可以放置的球的个数
	var max_balls func(int) int
	max_balls = func(x int) int {
		balls := 1
		left, right := 0, 0
		for right+1 < len(position) {
			if position[right+1]-position[left] > x {
				left = right + 1
				balls += 1
			}
			right += 1
		}
		return balls
	}

	// 二分搜索
	l, r := 1, position[len(position)-1]-position[0]
	for l <= r {
		mid := l + (r-l)/2
		if max_balls(mid) >= m {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
