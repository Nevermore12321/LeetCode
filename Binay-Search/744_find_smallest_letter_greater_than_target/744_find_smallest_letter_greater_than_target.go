func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	// 二分搜索，注意点就是，求找 第一个比目标字母大的字母
	// 因此如果找到目标字母，继续向右寻找
	for l <= r {
		mid := l + (r-l)/2
		if letters[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	// 如果 l 比最后一个还大，那么就返回第一个 字母
	if l >= len(letters) || l < 0 {
		return letters[0]
	}

	return letters[l]
}
