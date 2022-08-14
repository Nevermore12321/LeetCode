/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l, r := 1, n
	// 使用 二分法
	for l <= r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {	// 如果是坏的，不确定是不是第一个坏的，因此需要向左边继续找
			r = mid - 1
		} else {				// 如果不是坏的，那么肯定在当前版本的后面，因此向右边找
			l = mid + 1
		}
	}

	// 因为 l<=r ，因此 l 就是最后一个好版本的后一个位置，也就是 第一个 坏版本
	return l
}
