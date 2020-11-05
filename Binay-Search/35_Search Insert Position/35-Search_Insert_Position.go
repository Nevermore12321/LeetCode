package BinarySearch

func SearchInsert(nums []int, target int) int {

	var binarySearchFunc func(int, int) int
	binarySearchFunc = func(l int, r int) int {
		if l > r {
			return l
		} else {
			mid := l + (r - l) / 2
			if target == nums[mid] {
				return mid
			} else if target > nums[mid] {
				return binarySearchFunc(mid + 1, r)
			} else {
				return binarySearchFunc(l, mid - 1)
			}
		}
	}

	return binarySearchFunc(0, len(nums) - 1)

}


func SearchInsert2(nums []int, target int) int {

	var (
		l int = 0
		r int = len(nums) - 1
	)
	for l <= r {
		mid := l + (r - l) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}