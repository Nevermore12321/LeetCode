func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 前缀和 ，sums[0] = 0, sums[1] = nums[0], sums[2] = nums[1] + nums[0] ... sums[i+1] = nums[i] + ... + nums[0]
	sums := make([]int, len(nums)+1)
	minWidth := math.MaxInt
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	// 二分查找，找到 第一个大于 t 的索引
	// 注意 这里 l>=r，l 要么是 t，没有找到，l就是第一个大于 t 的索引
	// 如果 l>r, 那么 l 要么是 t，没有找到，l 就是最后一个小于 t 的索引
	binarySearch := func(n []int, t int) int {
		l, r := 0, len(n)-1
		for l <= r {
			mid := l + (r-l)/2
			if n[mid] == t {
				return mid
			} else if n[mid] > t {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return l
	}
	
	// 取最小值
	minInt := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// 从 0 开始找，例如 sums[0], 要找到 sums[t] - sums[i] >= target, 也就是 找到 sums[t], 使得 target + sums[i] <= sums[t]
	for i := 0; i < len(sums); i++ {
		t := sums[i] + target
		index := binarySearch(sums, t)
		if index == len(sums) {
			continue
		}
		minWidth = minInt(index-i, minWidth)

	}
	if minWidth == math.MaxInt {
		return 0
	}
	return minWidth

}
