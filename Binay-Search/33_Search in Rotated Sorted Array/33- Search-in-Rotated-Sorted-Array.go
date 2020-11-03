package BinarySearch

func Search(nums []int, target int) int {

	//binarySearchFunc := func(l, r int) int {
	//	var (
	//		tmpLeft int = l
	//		tmpRight int = r
	//	)
	//
	//	for tmpLeft <= tmpRight {
	//		tmpMid := tmpLeft + (tmpRight - tmpLeft) / 2
	//
	//		if target == nums[tmpMid] {
	//			return tmpMid
	//		} else if target < nums[tmpMid] {
	//			tmpRight = tmpMid - 1
	//		} else if target > nums[tmpMid] {
	//			tmpLeft = tmpMid + 1
	//		}
 	//	}
 	//	return -1
	//}

	var binaryDivideFunc func(int, int) int
	binaryDivideFunc = func(L, R int) int {
		var (
			left int = L
			right int = R
		)

		if left > right {
			return -1
		}
		mid := left + (right - left) / 2

		if target == nums[mid] {
			return mid
		}

		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				return binaryDivideFunc(left, mid - 1)
			} else {
				return binaryDivideFunc(mid + 1, right)
			}
		} else if nums[right] >= nums[mid] {
			if target > nums[mid] && target <= nums[right] {
				return binaryDivideFunc(mid + 1, right)
			} else {
				return binaryDivideFunc(left, mid - 1)
			}
		}

		return -1
	}

	res := binaryDivideFunc(0, len(nums)-1)
	return res
}