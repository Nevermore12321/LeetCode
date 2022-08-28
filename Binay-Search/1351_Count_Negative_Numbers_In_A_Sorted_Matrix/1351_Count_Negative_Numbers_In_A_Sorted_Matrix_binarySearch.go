func countNegatives(grid [][]int) int {
	// 二分搜索
	var binarySearch func(int, int, []int) int
	total := 0

	// 二分搜索，注意，这里返回 l ，第一个小于0的索引
	binarySearch = func(l, r int, nums []int) int {
		if l > r {
			return l
		}
		mid := l + (r-l)/2
		if nums[mid] < 0 {
			return binarySearch(l, mid-1, nums)
		} else {
			return binarySearch(mid+1, r, nums)
		}
	}

	for i := 0; i < len(grid); i++ {
		index := binarySearch(0, len(grid[0])-1, grid[i])
		// 如果 index 不是 -1，说明找到了 小于0的索引，计算当前 i行有多少小于0的元素个数
		if index != -1 {
			total += len(grid[0]) - index
		}
	}
	return total
}
