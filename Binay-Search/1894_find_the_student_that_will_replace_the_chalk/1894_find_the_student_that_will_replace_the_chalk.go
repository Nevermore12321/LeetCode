func chalkReplacer(chalk []int, k int) int {
	total := 0
	// 求 chalk 数组的所有元素和
	for _, v := range chalk {
		total += v
	}

	// 这里计算最后一轮还剩下的 k'
	round := k % total

	// 继续从 chalk 数组的 0 下标开始找，找到累加和第一个大于 k' 的元素下标
	for i, v := range chalk {
		if round < v {
			return i
		} else {
			round -= v
		}
	}
	return 0
}
