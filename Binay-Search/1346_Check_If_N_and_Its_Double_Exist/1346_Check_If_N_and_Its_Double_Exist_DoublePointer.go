func checkIfExist(arr []int) bool {
	// 双指针

	// 排序
	sort.Ints(arr)

	// v > 0 的情况，从前往后
	fast := 0		// fast 指针寻找 v*2
	for slow := fast; slow < len(arr); slow++ {
		// 找到第一个 arr[slow]*2 大于等于 arr[fast] 的 fast 索引
		for fast < len(arr) && arr[slow]*2 < arr[fast] {
			fast += 1
		}
		// 如果 arr[slow]*2 == arr[fast] 返回
		if fast != len(arr) && slow != fast && arr[slow]*2 == arr[fast] {
			return true
		}
	}

	// v > 0 的情况，从后往前
	fast = len(arr) - 1
	for slow := fast; slow >= 0; slow-- {
		// 找到第一个 arr[slow]*2 小于等于 arr[fast] 的 fast 索引
		for fast >= 0 && arr[slow]*2 < arr[fast] {
			fast -= 1
		}
		// 如果 arr[slow]*2 == arr[fast] 返回
		if fast >= 0 && slow != fast && arr[slow]*2 == arr[fast] {
			return true
		}
	}

	return false
}
