package BinarySearch


func MyPow(x float64, n int) float64 {
	var (
		ans float64 = 1.0
		k int64 = int64(n)
	)

	if n < 0 {
		k = -k
		x = 1/x
	}

	for k != 0 {
		if k & 1 == 1 {
			ans *= x
		}
		x *= x
		k >>= 1
	}
	return ans
}