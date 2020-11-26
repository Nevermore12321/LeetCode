package BinarySearch



func FindMin(nums []int) int {
	if len(nums) == 0 || nums[0] < nums[len(nums) - 1]{
		return nums[0]
	}

	var binarySearch func(l, r int) int
	binarySearch = func(l, r int) int {
		if l > r {
			return nums[0]
		}

		mid := l + (r - l) / 2

		if mid >= 1 && nums[mid] < nums[mid - 1] {
			return nums[mid]
		} else if mid < r && nums[mid] > nums[mid + 1] {
			return nums[mid + 1]
		}

		if nums[mid] < nums[r] {
			return binarySearch(l, mid - 1)
		} else {
			return binarySearch(mid + 1, r)
		}
	}


	return binarySearch(0, len(nums) - 1)
}