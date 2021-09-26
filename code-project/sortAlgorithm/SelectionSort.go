package sortAlgorithm

func SelectionSort(nums []int) {
	var length int = len(nums)

	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}


func SelectionSortAdvanced(nums []int) {
	var (
		left int = 0
		right int = len(nums) - 1
	)


	for left < right {
		minIndex := left
		maxIndex := right

		if nums[left] > nums[right] {
			nums[left], nums[right] = nums[right], nums[left]
		}
		//  在 (left, right) 区间内找最小和最大值
		for i := left + 1; i < right; i++ {
			if  nums[i] <  nums[minIndex] {
				minIndex = i
			} else if nums[i] > nums[maxIndex] {
				maxIndex = i
			}
		}
		//  找到 最大和最小值 后，交换位置到 left 和 right 位置
		nums[left], nums[minIndex] = nums[minIndex], nums[left]
		nums[right], nums[maxIndex] = nums[maxIndex], nums[right]

		left++
		right--
	}
}
