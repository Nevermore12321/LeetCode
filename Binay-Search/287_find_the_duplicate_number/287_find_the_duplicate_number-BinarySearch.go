func findDuplicate(nums []int) int {

	// 找到原数组 nums 中 小于等于t 的元素的总数
	findLessNums := func(t int) int {
		total := 0
		for _, v := range nums {
			if v <= t {
				total += 1
			}
		}
		return total
	}

	l, r := 1, len(nums)
	// 使用 二分
	for l <= r {
		mid := l + (r-l)/2			//  mid 就是抽屉数
		lessNums := findLessNums(mid)
		if lessNums <= mid {		// 	抽屉多，苹果少，说明 mid 前没出现重复，继续向右找
			l = mid + 1
		} else {					//  抽屉少，苹果多，说明 mid 前就出现了重复，继续向左找
			r = mid - 1
		}
	}

	//  最后 l 返回的是 第一个出现 苹果多的位置
	return l
}
