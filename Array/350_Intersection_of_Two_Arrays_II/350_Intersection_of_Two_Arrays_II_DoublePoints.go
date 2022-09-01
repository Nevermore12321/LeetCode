func intersect(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	res := []int{}
	// 将两个数组排序
	sort.Ints(nums1)
	sort.Ints(nums2)
	// 初始化双指针，分别指向两个数组的起始位置
	i, j := 0, 0

	// 只要有一个指针到结尾，停止循环
	for i < m && j < n {
		if nums1[i] == nums2[j] {			//  找到交集元素，添加到返回结果数组中，双指针都向前进一步
			res = append(res, nums1[i])
			i += 1
			j += 1
		} else if nums1[i] < nums2[j] {		//  nums1[i] 小，i 指针向前进一步
			i += 1
		} else {							//  nums1[j] 小，j 指针向前进一步
			j += 1
		}

	}
	return res
}
