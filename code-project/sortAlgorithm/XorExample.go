package sortAlgorithm

import "fmt"

func PrintOddTimesNum1(nums []int) {
	var eor int = 0

	// 所有元素异或，得到奇数次的数
	for _, v := range nums {
		eor ^= v
	}

	print(eor)

}

func PrintOddTimesNums1Test() {
	nums := []int {2,1,3,1,2,3,1,3,1}

	PrintOddTimesNum1(nums)
}


func PrintOddTimesNum2(nums []int) {
	var (
		eor1 int = 0
		eor2 int = 0
	)

	// 计算所有元素异或，结果也就是 a^b
	for _, v := range nums {
		eor1 ^= v
	}

	// eor = a ^ b
	// eor != 0
	// eor 必然有一个bit位为 1

	// 找出 a^b 结果最右边为 1 的 bit 位
	// 取反加一在与本身
	rightOne := eor1 & (^eor1 + 1)

	// eor2 表示所有 rightOne 位为 1 的异或的结果，也就是 a 或者 b
	for _, v := range nums {
		if (v & rightOne) == 0 {
			eor2 ^= v
		}
	}

	fmt.Printf("a = %d, b = %d\n", eor2, eor2 ^ eor1)


}


func PrintOddTimesNums2Test() {
	nums := []int {2,1,3,1,2,3,1,3,1,5,4,4,5,5}

	PrintOddTimesNum2(nums)
}