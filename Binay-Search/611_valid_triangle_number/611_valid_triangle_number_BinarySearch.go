func triangleNumber(nums []int) int {
	// 排序
	sort.Ints(nums)
	n := len(nums)
	res := 0

	// 二分查找
	var binarySearch func(int, int, int) int
	binarySearch = func(l, r, target int) int {
		if l > r { // 返回 l 的下标
			return l
		}

		mid := l + (r-l)/2
		//  这里注意，要寻找第一个 c >= a + b 的索引，因此，如果 nums[mid] == target 的时候，要继续向左找到第一个等于的下标
		if nums[mid] >= target { //  nums[mid] 大，向左找
			return binarySearch(l, mid-1, target)
		} else { // nums[mid] 小，向右找
			return binarySearch(mid+1, r, target)
		}
	}

	// 第一重循环，边 a
	for i := 0; i < n; i++ {
		// 第二重循环，边 b
		for j := i + 1; j < n; j++ {
			// 第三重循环 边 c，使用二分查找
			k := binarySearch(j+1, n-1, nums[i]+nums[j])

			if k <= n {
				res += k - j - 1
			} else { // 这里注意，如果k越界了，那么说明，j 到结尾的所有元素都满足要求
				res += n - j
			}
		}
	}
	return res
}
