func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	// y=f(x) 计算当时速为 x 时，一共需要花费多长时间
	var spend_times func(int) float64
	spend_times = func(x int) float64 {
		var total float64
		for i := 0; i < n; i++ {
			if i == n-1 { // 最后一段距离，不需要取整
				total += float64(float64(dist[i]) / float64(x))
			} else { // 否则，花费的时间为 距离 dist[i] / 时速 x，结果向上取整，因为整点发车
				total += math.Ceil(float64(dist[i]) / float64(x))
			}
		}
		return total
	}

	// 计算 dist 中的最长的距离
	var maxValue func([]int) int
	maxValue = func(arr []int) int {
		maxValue := 0
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > maxValue {
				maxValue = arr[i]
			}
		}
		return maxValue
	}

	var myMax func(int, int) int
	myMax = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	if hour <= float64(n-1) {
		return -1
	}

	// 这里是二分搜索
	l, r := 1, myMax(maxValue(dist), dist[n-1]*100)

	for l <= r {
		mid := l + (r-l)/2
		if spend_times(mid) <= hour { // 花费时间小于等于 hour，说明时速 x 可以继续减小，使得花费的时间增大
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}
