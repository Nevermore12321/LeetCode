package BinarySearch


func Search2(nums []int, target int) bool {

	var binarySearch func(left, right int) bool

	binarySearch = func(left, right int) bool {

		if left > right {
			return false
		}

		mid := left + (right - left) / 2

		if nums[mid] == target {
			return true
		}

		//  注意不能等于，因为例如 56755 5 55555 , mid = 5 , nums[mid]==nums[left], 但左边还是无序的
		if nums[mid] > nums[left] {
			if nums[left] <= target && target < nums[mid] {
				return binarySearch(left, mid - 1)
			} else {
				return binarySearch(mid + 1, right)
			}
		} else if nums[mid] < nums[left] {
			if nums[mid] < target && target <= nums[right] {
				return binarySearch(mid + 1, right)
			} else {
				return binarySearch(left, mid - 1)
			}
		} else{
			return binarySearch(left + 1, right)
		}
 	}

 	return binarySearch(0, len(nums) - 1)
}