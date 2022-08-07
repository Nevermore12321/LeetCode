func twoSum(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}
	var recursionBinarySearch func(int, int, int) int
	recursionBinarySearch = func(l, r, t int) int {
		if l > r {
			return -1
		}
		mid := l + (r-l)/2
		if numbers[mid] == t {
			return mid
		} else if numbers[mid] > t {
			return recursionBinarySearch(l, mid-1, t)
		} else {
			return recursionBinarySearch(mid+1, r, t)
		}
	}

	// 固定第一个数，numbers[i]
	for i := 0; i < len(numbers); i++ {
		// 使用 二分查找 查找 第二个数，目标值为 target-numbers[i]
		index := recursionBinarySearch(i+1, len(numbers)-1, target-numbers[i])
		if index != -1 {
			return []int{i + 1, index + 1}
		}
	}
	return []int{-1, -1}
}
