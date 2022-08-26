func findKthPositive(arr []int, k int) int {
	// 初始化
	missCounts := 0			// 当前丢失的元素个数
	curValue := 1			// 当前丢失的元素 值
	index := 0				// 当前 遍历到了 arr 的索引
	lastMiss := 0			// 当前最后一个丢失的元素值

	// curValue 从 1 开始找丢失值
	for missCounts < k {
		//  如果 当前丢失的元素 不等于 数组元素，数组元素前还有 丢失值，
		if curValue != arr[index] {
			missCounts += 1
			lastMiss = curValue
		} else {
			// 这里需要注意，如果找到了数组的结尾，还是没有找到 第 k 个丢失值，那么就从元素最后一个元素开始 +1
			// 如果还在数组内，那么继续查找
			if index+1 < len(arr) {
				index += 1
			}
		}
		curValue += 1
	}
	return lastMiss
}
