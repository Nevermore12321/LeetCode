func intersect(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]int)
	res := []int{}
	// 遍历数组 nums1 ，将所有元素添加到 哈希表，并且计算每一个元素出现的次数
	for _, v1 := range nums1 {
		hashMap[v1] += 1
	}

	// 遍历数组 nums2
	for _, v2 := range nums2 {
		// 如果当前元素在哈希表中，那么存在交集，将该元素添加到结果列表中，并且将哈希表中的次数减一
		if v, ok := hashMap[v2]; ok && v > 0 {
			res = append(res, v2)
			hashMap[v2] -= 1
		}
	}
	return res
}
