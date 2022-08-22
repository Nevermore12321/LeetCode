func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	// 对 arr2 排序
	sort.Ints(arr2)
	// 绝对值
	intAbs := func(a int) int {
		if a >= 0 {
			return a
		}

		return -a
	}

	// 二分搜搜，找到第一个大于等于 target 的索引值
	binarySearch := func(nums []int, target int) int {
		l, r := 0, len(nums)-1
		for l <= r {
			mid := l + (r-l)/2

			if target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return l
	}

	// v1 - v2 <= b
	count := 0
	for _, v1 := range arr1 {
		// 找到最接近 v1 的值
		index := binarySearch(arr2, v1)
		var flag bool = true
		// 注意这里需要判断两个元素，如果相等，只需要判断 index 索引即可；
		// 如果 x 是第一个大于 v1 的元素，那么就需要判断，index-1 也就是上一个小于 v1 的元素，求他们最接近 v1 的距离
		if index < len(arr2) && intAbs(v1-arr2[index]) <= d {
			flag = false
			continue
		} else if (index-1) >= 0 && intAbs(arr2[index-1]-v1) <= d {
			flag = false
			continue
		}
		if flag {
			count += 1
		}
	}
	return count
}
