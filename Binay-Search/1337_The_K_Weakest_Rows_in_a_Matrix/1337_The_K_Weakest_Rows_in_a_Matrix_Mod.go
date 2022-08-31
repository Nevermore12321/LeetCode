func kWeakestRows(mat [][]int, k int) []int {
	// 临时数组
	score := []int{}
	res := make([]int, k)
	for index, row := range mat { // 计算每一行的 1 的个数
		pos := sort.Search(len(row), func(i int) bool { return row[i] == 0 }) // 利用二分找到第一个为0的元素
		score = append(score, pos*100+index)                                  // 将结果根据公式计算后加入 临时数组
	}

	// 临时数组排序
	sort.Ints(score)

	// 取出排序后的前k个元素
	for i := 0; i < k; i++ {
		res[i] = score[i] % 100 // 并且根据 value 计算出 其 行号
	}
	return res

}
