/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	var binarySearch func(int, int) int
	binarySearch = func(l, r int) int {
		mid := l + (r-l)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == 1 {
			return binarySearch(mid+1, r)
		} else {
			return binarySearch(l, mid-1)
		}
	}
	return binarySearch(1, n)
}
