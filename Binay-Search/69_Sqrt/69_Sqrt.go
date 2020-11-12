package BinarySearch


func mySqrt(x int) int {
	var binarySeach func(l, r int) int

	binarySeach = func(l, r int) int {
		mid := l + (r - l) / 2
		if l > r {
			return r
		}
		if mid * mid == x {
			return mid
		} else if mid * mid > x {
			return binarySeach(l, mid - 1)
		} else {
			return binarySeach(mid + 1, r)
		}
	}

	return binarySeach(1, x)
}