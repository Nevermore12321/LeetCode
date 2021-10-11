package lib

import (
	"fmt"
	"math"
)

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

func (m *MaxHeap) __swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

// 由底向上 上浮操作
// k 表示插入节点 在堆数组中的碎银
func (m *MaxHeap) __swim(k int) {
	// k/2 =》 表示 k 索引位置的父节点的索引
	// 比较插入元素是否大于父节点元素，如果大，交换，否则退出
	// 如果插入元素比父节点元素大，交换位置，然后再以交换后的索引位置，与其父节点在次比较，以此类推
	// k=1 是根节点，因此退出条件必须大于根节点索引
	for k > 1 && m.data[k/2] < m.data[k] {
		m.__swap(k/2, k)
		k /= 2
	}
}

// 插入元素
func (m *MaxHeap) Insert(item int) {
	//  堆元素个数增加1
	m.count++
	//  新元素插入到堆数组
	m.data[m.count] = item
	//  上浮操作，满足堆有序化
	m.__swim(m.count)
}

//  由顶向下  下沉操作
//  k -> 表示要下沉的某个元素，将最后一个叶子节点 与 根节点交换，然后将交换后的根节点下沉
func (m *MaxHeap) __sink(k int) {
	// 2*k 表示 k 的左孩子，是否存在左孩子
	for 2*k <= m.count {
		// j 表示索引 k 的左右孩子中较大者
		// 先初始化成左孩子
		j := 2*k

		// 如果索引 k 的右孩子大，j 表示右孩子的索引
		if j+1 <= m.count && m.data[j+1] > m.data[j] {
			j += 1
		}

		// 然后判断 父节点 与 左右孩子较大者，谁更大，如果 父节点比左右孩子都大，直接退出，否则，交换
		if m.data[k] > m.data[j] {
			break
		}

		m.__swap(k, j)

		k = j
	}
}

// 删除堆中的最大元素，如果堆为空，返回 -1
func (m *MaxHeap) DelMax() int {
	if m.count == 0 {
		fmt.Println("堆为空，无法删除！")
		return -1
	}

	// 取出最大元素
	max := m.data[1]

	// 交换 最后一个元素 与 根元素
	m.__swap(1, m.count)

	// 减少堆元素
	m.count--

	// 将新的根元素下沉
	m.__sink(1)

	return max
}

// 判断最后一行是否满元素
func isPower(value int) bool {
	n := math.Log2(float64(value));

	return (math.Pow(2, n) == float64(value));
}

func (m *MaxHeap) Show() {

	heapHeight := math.Floor(math.Log2(float64(m.count)));
	var (
		// 第一行的 tab 个数
		firstTabs float64
		intervalTabs float64
		currentHeight float64
		i int
		k float64
	)

	for i = 1; i <= m.count; i++ {
		currentHeight = math.Floor(math.Log2(float64(i)));

		// 判断当前行是否是满行
		if ( isPower(i) ) {
			/* print first tabs */
			fmt.Println()
			// 最开始有多少 tab
			firstTabs = math.Pow(2, heapHeight-currentHeight) - 1
			if firstTabs != 0 {
				for k = 0; k < firstTabs; k++ {
					fmt.Print("\t")
				}
			}
			fmt.Print(m.data[i])
		} else {
			// 如果不是满行，只能是最后一行，
			/* print interval tabs */
			intervalTabs = math.Pow(2, heapHeight - currentHeight + 1);
			for k = 0; k < intervalTabs; k++ {
				fmt.Print("\t")
			}
			fmt.Print(m.data[i])
		}
	}

	fmt.Println()
}

func MaxHeapTest() {
	mp := MaxHeapInit(11)
	mp.Insert(62)
	mp.Insert(41)
	mp.Insert(30)
	mp.Insert(28)
	mp.Insert(16)
	mp.Insert(22)
	mp.Insert(13)
	mp.Show()

	fmt.Println()
	fmt.Println(mp.DelMax())
	mp.Show()

	fmt.Println()
	fmt.Println(mp.DelMax())
	mp.Show()

	fmt.Println()
	mp.Insert(29)
	mp.Show()

}