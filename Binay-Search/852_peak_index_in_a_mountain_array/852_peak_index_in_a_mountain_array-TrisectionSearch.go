func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l+1 < r {
		// 三分的左右两个分割点
		midl := l + (r-l+1)/3
		midr := r - (r-l+1)/3
		//  如果 f(midl) < f(midr) , midr 更靠近最大值，忽略 [l, midl), 令 l = midl
		if arr[midl] < arr[midr] {
			l = midl
		} else {	//  如果 f(midl) >= f(midr) , midl 更靠近最大值，忽略 (midr, r], 令 r = midr
			r = midr
		}
	}

	// 最后，l 和 r 是相邻时退出，f(l) 和 f(r) 的大者，就是最大值
	if arr[l] < arr[r] {
		return r
	} else {
		return l
	}
}
