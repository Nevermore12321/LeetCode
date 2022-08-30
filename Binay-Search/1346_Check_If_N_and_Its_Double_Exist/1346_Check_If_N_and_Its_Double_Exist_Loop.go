func checkIfExist(arr []int) bool {
	// 暴力搜索

	for i, v1 := range arr {	// 第一轮，遍历每一个元素 v
		for j, v2 := range arr {	// 第二轮，寻找是否存在 v*2 的元素存在
			if v1*2 == v2 && i != j {	// 如果 找到，并且 不是同一个元素，例如 0*2=0
				return true
			}
		}
	}
	return false
}
