func chalkReplacer(chalk []int, k int) int {
	n := len(chalk)
	total := 0

	// 同样计算 chalk 数组的综合，这里在有个注意点，如果在第一轮就能找到第一个大于 k 的元素
	for i := 0; i < n; i++ {
		total += chalk[i]
		if total > k {
			return i
		}
	}

	// 通过 递归 ，将k值变为最后一轮还剩的 k'
	return chalkReplacer(chalk, k%total)

}
