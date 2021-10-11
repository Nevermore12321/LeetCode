package sortAlgorithm

import (
	"fmt"
	"program-algorithm/lib"
)

func __merge(nums []int, l, mid, r int) {
	// 将 nums[l, mid], nums[mid+1, r] 进行归并

	// 开辟辅助空间，并且将原数组复制到辅助数组
	aux := make([]int, r - l + 1)
	for i := l; i <= r; i++ {
		aux[i - l] = nums[i]
	}

	// i、j 分别表示 辅助数组 左右两部分的 下标
	// 设置下标，i 从 l 开始， j 从 mid+1 开始
	var (
		i = l
		j = mid + 1
	)

	//  k 表示 原数组 r-l 的位置
	for k := l; k <= r; k++ {
		if i > mid {				// 说明 左半部分结束，还剩下右半部分，不需要比较，直接复制，右边+1
			nums[k] = aux[j - l]	// aux 索引从 0 开始
			j++
		} else if j > r {			// 说明 右半部分结束，还剩下左半部分，不需要比较，直接复制，左边+1
			nums[k] = aux[i - l]
			i++
		} else if aux[i - l] < aux[j - l] {	// 都未完成，且左边的小，将左边的复制，左边+1
			nums[k] = aux[i - l]
			i++
		} else {					//  都未完成，且右边的小，将右边的复制，右边+1
			nums[k] = aux[j - l]
			j++
		}

	}
}

func __mergeSort(nums []int, l, r int)  {
	//  该函数利用递归，实现自顶向下的归并排序

	//  递归的终止条件，类似 二分搜索
	if l >= r {
		return
	}

	mid := l + (r - l) / 2

	//  递归，对两边分别进行 排序，排好序的两个数组，可以使用归并合成一个有序数组
	__mergeSort(nums, l, mid)
	__mergeSort(nums, mid + 1, r)
	__merge(nums, l, mid, r)
}

func MergeSort(nums []int) {
	__mergeSort(nums, 0, len(nums) - 1)
}


// ---------------------------------------------------\

func __insertionSort(nums []int, l, r int) {
	for i := l + 1; i <= r; i++ {
		var (
			tmp = nums[i]
			j	= i
		)

		for ; j > l && tmp < nums[j - 1]; j-- {
			nums[j] = nums[j - 1]
		}
		nums[j] = tmp
	}
}


//  归并排序的优化，
func __mergeSortAdvanced(nums []int, l, r int) {
	//  递归结束的条件，如果 数组长度 <= 15 ，则直接使用插入排序，并返回
	//  优化1 小规模数组 n <= 15 ，使用插入排序
	if r - l <= 15 {
		__insertionSort(nums, l, r)
		return
	}

	//  如果 长度 > 15 继续递归
	mid := l + ((r - l) >> 2)
	__mergeSortAdvanced(nums, l, mid)
	__mergeSortAdvanced(nums, mid + 1, r)

	//  由于左半部分的最大值 nums[mid] 比 右半部分的最小值 nums[mid+1] 还要小，说明整个数组有序，不需要归并，此外其他情况需要归并
	//  优化2 对有序数组的跳过
	if nums[mid] > nums[mid + 1] {
		__merge(nums, l, mid, r)
	}

}

func MergeSortAdvanced(nums []int) {
	__mergeSortAdvanced(nums, 0, len(nums) - 1)
}


func MergeSortTest() {
	nums := lib.GenerateNearlyOrderedArray(100, 100)
	fmt.Println(nums)

	MergeSortAdvanced(nums)
	fmt.Println(nums)

	lib.SortDuration(MergeSort, "MergeSort")
	lib.SortDuration(MergeSortAdvanced, "MergeSortAdvanced")
}