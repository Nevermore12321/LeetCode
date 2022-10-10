func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	sum := 0
	maxDiff := 0
	n := len(nums1)
	// 辅助数组，排序
	rec := append(sort.IntSlice(nil), nums1...)
	rec.Sort()

	// 求绝对值
	var abs func(int) int
	abs = func(a int) int {
		if a >= 0 {
			return a
		}
		return -a
	}

	// 求最大值
	var maxValue func(int, int) int
	maxValue = func(a int, b int) int {
		if a >= b {
			return a
		} else {
			return b
		}
	}

	// 二分搜素
	var binarySearch func([]int, int, int, int) int
	binarySearch = func(arr []int, l int, r int, target int) int {
		if l > r {
			return l
		}
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			return binarySearch(arr, mid+1, r, target)
		} else {
			return binarySearch(arr, l, mid-1, target)
		}
	}

	// 二分搜索，每一个 i 都去找到一个新的最小的 newdiff, 如果 与 oldDiff 差别最大，那么就选这个 i
	// 因为 oldDiff - newDiff 差值越大，newDiff 就越小
	for i := 0; i < n; i += 1 {
		oldDiff := abs(nums1[i] - nums2[i])
		// index 是第一个 大于或等于 target 的索引
		index := binarySearch(rec, 0, n-1, nums2[i])

		// 最接近 target 的有可能是 第一个大于的，也有可能是最后一个小于的
		if index < n {
			maxDiff = maxValue(maxDiff, oldDiff-abs(rec[index]-nums2[i]))
		}
		if index > 0 {
			maxDiff = maxValue(maxDiff, oldDiff-abs(rec[index-1]-nums2[i]))
		}

		sum += oldDiff
	}

	return (sum - maxDiff) % (1e9 + 7)
}
