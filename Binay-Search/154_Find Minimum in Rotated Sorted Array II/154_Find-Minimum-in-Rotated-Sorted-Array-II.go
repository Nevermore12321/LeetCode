package BinarySearch


func FindMin2(nums []int) int {
	if len(nums) == 1 || nums[0] < nums[len(nums) - 1] {
		return nums[0]
	}
	var binarySearch func(left, right int) int


	binarySearch = func(left, right int) int {
		if left > right {
			return nums[left]
		}

		mid := left + (right - left) / 2

		if mid >= 1 && nums[mid] < nums[mid - 1] {
			return nums[mid]
		} else if mid < right && nums[mid] > nums[mid + 1] {
			return nums[mid + 1]
		}

		if nums[mid] < nums[right] {
			return binarySearch(left, mid - 1)
		} else if nums[mid] > nums[right] {
			return binarySearch(mid + 1, right)
		} else {
			return binarySearch(left, right - 1)
		}
	}

	return binarySearch(0, len(nums) - 1)
}