package BinarySearch

func Divide(dividend int, divisor int) int {

	INT_MAX := int(^uint32(0) >> 1)
	INT_MIN := ^INT_MAX

	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	} else if divisor == -1 {
		if dividend > INT_MIN {
			return -dividend
		}
		return INT_MAX
	}

	absFunc := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	divid := absFunc(dividend)
	divis := absFunc(divisor)

	signed := true

	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		signed = false
	}

	//  重点，递归求解
	var divFunc func(int, int) int			//  这里声明，是为了下面的匿名函数可以递归
 	divFunc = func(a, b int) int {			// a 是 被除数 ，b 是除数
		if a < b {
			return 0
		}
		multiple := 1			// 表示  除数的倍数 。也就是 2^0 开始
		tmpNum := b		// 除数翻倍后的数

		for (tmpNum + tmpNum) <= a {
			tmpNum += tmpNum
			multiple += multiple
		}
		return multiple + divFunc(a - tmpNum, b)
	}

	res := divFunc(divid, divis)
	if signed {
		return res
	}
	return -res
}