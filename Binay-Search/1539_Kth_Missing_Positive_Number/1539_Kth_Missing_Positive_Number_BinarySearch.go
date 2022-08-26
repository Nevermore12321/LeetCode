func findKthPositive(arr []int, k int) int {
	// 如果 arr 的第一个元素比k大，那么 k 就是 第k个丢失的元素
	if arr[0] > k {
		return k
	}
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		// 这里计算，arr[mid] 这个元素前丢失的元素个数
		missCounts := arr[mid] - mid - 1
		if missCounts < k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	// 这里的计算公式为：
	// arr[l-1] - (l - 1) - 1 表示 arr[l-1] 前丢失的元素，
	// k - (arr[l-1] - (l - 1) - 1) 表示 arr[l-1] 到 arr[l] 还差几个丢失的元素
	// k - (arr[l-1] - (l - 1) - 1) + arr[l-1] 就是 第k个丢失的元素值
	return k - (arr[l-1] - (l - 1) - 1) + arr[l-1]
}
