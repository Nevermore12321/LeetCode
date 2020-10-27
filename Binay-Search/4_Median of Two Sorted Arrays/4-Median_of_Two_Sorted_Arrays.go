package BinarySearch

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//  num1 的长度 m
	m := len(nums1)
	//  num2 的长度 n
	n := len(nums2)

	if m > n {
		return FindMedianSortedArrays(nums2, nums1)
	}

	//  二分查找 的 Left 和 Right， 以及 nums1 中 i 和 nums2 中的 j
	var (
		Left int = 0
		Right int = m
		k int = (m + n + 1) / 2
	)

	//  二分查找. 最后返回的 l 的值
	for Left < Right {
		//  i 是 nums1 的中位数或右中位数
		i := (Right - Left) / 2 + Left
		j := k - i
		if nums1[i] < nums2[j-1] {
			Left = i + 1
		} else {
			Right = i
		}
	}

	fnMin := func (x, y int) int {
		if x < y {
		return x
	}
		return y
	}

	fnMax := func (x, y int) int {
		if x > y {
		return x
	}
		return y
	}

	fnIf := func(condition bool, whenTrue int, nums []int, index int) int {
		if condition {
			return whenTrue
		}
		return nums[index]
	}

	i := Left
	j := k - Left
	INT_MAX := int(^uint(0) >> 1)
	INT_MIN := ^INT_MAX

	c1 := fnMax(
		fnIf(i<=0, INT_MIN, nums1, i-1),
		fnIf(j<=0, INT_MIN, nums2, j-1),
		)

	if (m + n) % 2 == 1 {
		return float64(c1)
	}

	c2 := fnMin(
		fnIf(i>=m, INT_MAX, nums1, i),
		fnIf(j>=n, INT_MAX, nums2, j),
		)

	return (float64(c1) + float64(c2)) * 0.5
}