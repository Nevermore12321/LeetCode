package sortAlgorithm

import "program-algorithm/lib"

func InsertionSort(nums []int)  {
	//  i 从第一个元素开始，第0个表示已排序好
	for i := 1; i < len(nums); i++ {
		//  新元素  j ，往前搜索位置
		for j := i; j > 0; j-- {
			if nums[j] < nums[j - 1] {
				nums[j], nums[j - 1] = nums[j - 1], nums[j]
			} else {
				break
			}
		}
	}
}

func InsertionSortAdvanced(nums []int) {
	//  i 从第一个元素开始，第0个表示已排序好
	for i := 1; i < len(nums); i++ {
		//  存放 副本，和初始化最终位置的 j
		tmp := i
		j := i

		//  寻找 tmp 的位置，如果 tmp 小于当前的 j-1 元素，j-1元素就要后移
		for ; j > 0; j-- {
			if nums[j] < nums[j - 1] {
				nums[j] = nums[j - 1]
			} else {
				break
			}
		}
		//  将 tmp 放入到 最终 j 的位置
		nums[j]	= tmp
	}
}


func InsertionSortTest() {
	lib.SortDuration(InsertionSort, "InsertionSort")
	lib.SortDuration(InsertionSortAdvanced, "InsertionSortAdvanced")
}