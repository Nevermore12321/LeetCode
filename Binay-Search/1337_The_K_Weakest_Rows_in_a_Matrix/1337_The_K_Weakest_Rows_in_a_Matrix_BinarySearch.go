type pair struct {
	index int
	value int
}

type MinHeap struct {
	data  []pair
	count int
}

func MinHeapInit(capacity int) *MinHeap {
	heap := new(MinHeap)
	data := make([]pair, capacity+1)
	heap.data = data
	heap.count = 0
	return heap
}

func (mh *MinHeap) IsEmpty() bool {
	return mh.count == 0
}

func (mh *MinHeap) Size() int {
	return mh.count
}

func (mh *MinHeap) __swap(i, j int) {
	mh.data[i], mh.data[j] = mh.data[j], mh.data[i]
}

func (mh *MinHeap) __swim(k int) {
	for k > 1 && mh.data[k].value < mh.data[k/2].value {
		mh.__swap(k, k/2)
		k = k / 2
	}
}

func (mh *MinHeap) Insert(item pair) {
	mh.count += 1
	mh.data[mh.count] = item
	mh.__swim(mh.count)
}

func (mh *MinHeap) __sink(k int) {
	for 2*k <= mh.count {
		j := 2 * k
		if j+1 <= mh.count && mh.data[j].value > mh.data[j+1].value {
			j += 1
		} else if j+1 <= mh.count && mh.data[j].value == mh.data[j+1].value { // 这里需要注意，如果 value 的值相同，那么 索引 index 小的先出
			if mh.data[j].index > mh.data[j+1].index {
				j += 1
			}
		}
		if mh.data[k].value < mh.data[j].value {
			break
		} else if mh.data[k].value == mh.data[j].value { // 同样，如果两个子节点 value 相同，index 索引小的节点更小
			if mh.data[k].index < mh.data[j].index {
				break
			}
		}

		mh.__swap(j, k)
		k = j
	}
}

func (mh *MinHeap) DelMin() pair {
	if mh.count == 0 {
		return pair{-1, -1}
	}

	min := mh.data[1]
	mh.__swap(1, mh.count)

	mh.count -= 1
	mh.__sink(1)

	return min
}

func kWeakestRows(mat [][]int, k int) []int {
	minHeap := MinHeapInit(len(mat))
	// 使用二分搜索，找到第一个为 0 的索引
	// 这里的 sort.Search 就是使用的 二分搜索
	for index, row := range mat {
		pos := sort.Search(len(row), func(i int) bool { return row[i] == 0 })
		minHeap.Insert(pair{index: index, value: pos}) // 每计算出一个，都将{index，value}加入到最小堆中
	}
	// 找到前 k 个最小的元素
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = minHeap.DelMin().index // 从最小堆中，从小到大，删除元素。
	}
	return res
}
