package BinarySearch



func SearchRange(nums []int, target int) []int {

	var binarySearchFunc func(int, int) (int, int)
	binarySearchFunc = func(l, r int) (int, int) {
		if l > r {
			return -1, -1
		} else if l <= r {
			mid := l + (r - l) / 3
			if target == nums[mid] {
				leftPosition, _ := binarySearchFunc(l, mid - 1)
				_, rightPosition := binarySearchFunc(mid + 1, r)
				if leftPosition == -1 {
					leftPosition = mid
				}
				if rightPosition == -1 {
					rightPosition = mid
				}
				return leftPosition, rightPosition
			} else if target > nums[mid] {
				return binarySearchFunc(mid + 1, r)
			} else if target < nums[mid] {
				return binarySearchFunc(l, mid - 1)
			}
		}
		return -1, -1
	}

	left, right := binarySearchFunc(0, len(nums) - 1)
	return []int{left, right}
}