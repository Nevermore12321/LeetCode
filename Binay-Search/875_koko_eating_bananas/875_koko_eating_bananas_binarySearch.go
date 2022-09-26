func minEatingSpeed(piles []int, h int) int {
	// y=f(x) 其中 x 表示吃香蕉的速度，y 表示当吃香蕉的速度为 x 时，吃完需要的小时数
	// x 吃的速度越大，花费的时间y越小。因此 y 随着 x 的增大而减小

	// spend_times 表示计算 y，也就是计算 f(x)
	var spend_times func(int) int
	spend_times = func(x int) int {
		total_times := 0
		for _, v := range piles {
			total_times += int(math.Ceil(float64(v) / float64(x)))
		}
		return total_times
	}

	// maxPile 返回所有堆中香蕉最多的 堆大小
	var maxPile func(arr []int) int
	maxPile = func(arr []int) int {
		maxValue := 0
		for _, v := range arr {
			if v > maxValue {
				maxValue = v
			}
		}
		return maxValue
	}

	// 下面使用二分搜索，现在要找到最小的吃香蕉的速度x
	// l 最终就是第一个等于8 或者第一个小于8 的元素
	l, r := 1, maxPile(piles)

	for l <= r {
		mid := l + (r-l)/2
		if spend_times(mid) > h {
			l = mid + 1

		} else {
			r = mid - 1

		}
	}

	return l
}
