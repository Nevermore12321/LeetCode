package common

import (
	"math/rand"
	"time"
)

type ArrayGenerator struct {
}

func (ag ArrayGenerator) GeneratOrderedArray(n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, i)
	}
	return result

}

func (ag ArrayGenerator) GenerteRandomArray(n, bound int) []int {
	arr := make([]int, n)
	// 使用当前的纳秒作为随机数种子
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(bound)
	}
	return arr
}
