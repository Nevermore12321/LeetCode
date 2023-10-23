package linearSerach

import (
	"Algorithm_and_Data_Structure/common"
)

// Search 线性查找，其实就是遍历搜索，此方法用于可比较的类型
func Search[T common.Number](data []T, target T) int {
	for i, d := range data {
		if d == target {
			return i
		}
	}

	return -1
}

// 线性查找，此方法用于可比较的 自定义结构体类型(具有 Equals 方法的自定义结构体)
func CustomSearch[T common.CustomComparable](data []T, target T) int {
	for i, d := range data {
		if d.Equals(target) {
			return i
		}
	}

	return -1
}
