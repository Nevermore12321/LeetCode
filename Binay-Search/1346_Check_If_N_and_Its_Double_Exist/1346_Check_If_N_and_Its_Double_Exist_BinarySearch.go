func checkIfExist(arr []int) bool {
	// 二分搜索

	// 排序
	sort.Ints(arr)

	for index, value := range arr {			//	遍历每一个元素 v
		l, r := 0, len(arr)-1				// 二分搜搜，寻找是否存在 v*2 的元素存在
		for l <= r {
			mid := l + (r-l)/2
			if arr[mid] == value*2 && index != mid {	// 注意这里需要排除掉 0*2=0,元素本身
				return true
			} else if arr[mid] < value*2 {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
