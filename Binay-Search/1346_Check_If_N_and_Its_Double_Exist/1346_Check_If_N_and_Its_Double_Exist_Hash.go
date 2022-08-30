func checkIfExist(arr []int) bool {
	// 哈希表

	// 使用 map 来作为 哈希表
	hashMap := make(map[int]int, len(arr))

	// 遍历每一个元素
	for i, v := range arr {
		// 如果当前元素 v*2 已经在 哈希表中，那么返回 True
		if _, ok := hashMap[v*2]; ok {
			return true
		}

		// 如果 当前元素可以整除2，那么判断 v/2 是否在 哈希表中，如果在返回 True
		if v%2 == 0 {
			if _, ok := hashMap[v/2]; ok {
				return true
			}
		}
		// 每次都需要把 元素 v 放进哈希表
		hashMap[v] = i
	}

	return false
}
