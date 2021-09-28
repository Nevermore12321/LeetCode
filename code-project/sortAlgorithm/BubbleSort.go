package sortAlgorithm

import "program-algorithm/lib"

func BubbleSort(nums []int) {
	var (
		swapFlag = true
		n		 = len(nums) - 1
	)

	for swapFlag {
		swapFlag = false
		for i := 0; i < n; i++ {
			if nums[i] > nums[i + 1] {
				swapFlag = true
				nums[i], nums[i + 1] = nums[i + 1], nums[i]
			}
		}
		n--
	}
}

func BubbleSortAdvanced(nums []int) {
	var (
		newN = 1
		n	 = len(nums) - 1
	)

	for newN > 0 {
		newN = 0
		for i := 0; i < n; i++ {
			if nums[i] > nums[i + 1] {
				nums[i], nums[i + 1] = nums[i + 1], nums[i]
				newN = i + 1
			}
		}
		n = newN
	}

}

func BubbleSortTest() {
	lib.SortDuration(BubbleSort, "BubbleSort")
	lib.SortDuration(BubbleSortAdvanced, "BubbleSortAdvanced")
}