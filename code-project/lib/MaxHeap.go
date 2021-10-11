package lib

// 最大堆
type MaxHeap struct {
	// 存放所有元素
	data []int
	// 数组中由多少元素，也就是堆中的元素个数
	count int
}

func MaxHeapInit(capacity int) *MaxHeap {
	data := make([]int, capacity)
	maxHeap := new(MaxHeap)

	maxHeap.data = data
	maxHeap.count = 0

	return maxHeap
}

func (m *MaxHeap) IsEmpty() bool {
	return m.count == 0
}

func (m *MaxHeap) size() int {
	return m.count
}
