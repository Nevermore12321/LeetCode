func search(nums []int, target int) int {
	var binarySearch func(int, int) int

	binarySearch = func(l, r int) int {
		if l > r {
			return -1
		}
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			return binarySearch(l, mid-1)
		} else {
			return binarySearch(mid+1, r)
		}
	}

	return binarySearch(0, len(nums)-1)
}
