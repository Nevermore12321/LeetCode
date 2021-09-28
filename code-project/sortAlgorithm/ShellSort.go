package sortAlgorithm

import "program-algorithm/lib"

func ShellSort(nums []int) {
	var (
		h = 1
		n = len(nums)
	)

	//  计算增量，分别为 1，4，13，40，121，364....
	for h < n/3 {
		h = 3 * h + 1
	}

	// 根据增量分组
	for h >= 1 {
		// 从 第一个增量后，开始循环，与前一个 相差 增量h 的元素比较
		for i := h; i < n; i++ {
			tmp := nums[i]
			j := i
			for j >= h && tmp < nums[j - h] {
				nums[j]	= nums[j - h]
				j = j - h
			}
			nums[j]	= tmp
		}
		// 缩小 增量，继续分组排序
		h /= 3
	}

}

func ShellSortTest() {

	lib.SortDuration(ShellSort, "ShellSort")
}