func minimumSize(nums []int, maxOperations int) int {
	//【参考提示】【重点：理解题意，将题目转化为可以实现的问题】
	// 如果一个袋子最多只能装 x个，需要拆分 y 次；每个袋子能能装的球数越多，则需要拆分的次数越少（具有单调性）
	// 当 y > maxOperations 时，说明 x 不合题意
	// 则 x+=1，第一次当 y = maxOperations 时的x即为符合题意的答案（此过程可以用二分查找）
	var operation_times func(int) int
	operation_times = func(x int) int {
		// x是每个袋子最多装的球个数，返回拆分次数
		oper_times := 0
		for _, v := range nums {
			if v > x {
				oper_times += (v-1) / x
			}
		}
		return oper_times
	}

	var maxArr func([]int) int
	maxArr = func(arr []int) int {
		maxValue := 0
		for _, v := range arr {
			if v > maxValue {
				maxValue = v
			}
		}
		return maxValue
	}

	l, r := 1, maxArr(nums)
	if operation_times(l) == maxOperations {
		return l
	}

	// 进行二分查找 因为oper_time是随着x的增加单调递减的
	for l <= r {
		mid := l + (r-l)/2
		if operation_times(mid) <= maxOperations {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

