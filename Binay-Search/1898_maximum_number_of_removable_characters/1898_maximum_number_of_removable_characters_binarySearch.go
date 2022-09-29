func maximumRemovals(s string, p string, removable []int) int {
	// S 和 P 的长度
	lenS, lenP := len(s), len(p)
	var isSubstr func(int) bool
	isSubstr = func(x int) bool {
		// 如果 删除 ，则 removeState 对应下标的值为 True，表示删除
		removeState := make([]bool, lenS)
		// 前 x 个 removable 的值作为下标，从 s 字符串中删除，也就是更改s对应下标的状态为 True
		for i := 0; i < x; i++ {
			removeState[removable[i]] = true
		}

		// 判断 p 是否仍是 s 的子字符串
		// j 表示 p 的索引，i 表示 s 的索引，这里用到了 s 字符互不相同的特性。
		j := 0
		for i := 0; i < lenS; i++ {
			if !removeState[i] && s[i] == p[j] {
				j += 1
				if j == lenP { // 这里表示 j 已经匹配到了 p 的结尾，那么匹配成功，p 是 s 的子字符串
					return true
				}
			}
		}
		return false
	}

	// 二分搜索
	// l 左界限为 0，可以不去掉
	// r 右界限为 removable 的长度 - 1，如果全去掉，p 不可能是一个空串的子字符串
	l, r := 0, len(removable)

	for l <= r {
		mid := l + (r-l)/2
		if isSubstr(mid) { // 如果前x个去掉后，p还是s的子字符串，那么 x 可以继续增大，去掉更多的
			l = mid + 1
		} else { // 如果去掉x个后，p已经不是s的子字符串，那么去掉的太多了，需要减小 x
			r = mid - 1
		}
	}

	// 返回使得 p 是 s 子字符串的 最大的 x，l 表示第一个 p 不是 s 子字符串 的 x
	return r
}
