func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		// 唯一需要注意的就是 mid + 1 有可能会越界
		if mid+1 < len(arr) && arr[mid] <= arr[mid+1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
