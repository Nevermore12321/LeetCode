package Queue

import (
	"Algorithm_and_Data_Structure/common"
	"bytes"
	"fmt"
)

/*
	循环队列：front 表示首元素位置，tail 表示末元素的后一个位置
	入队/出队 时间复杂度都是 O(1)

初始： front == tail 队列为空

	  front
		0  1  2  3  4  5  6  7  8  9
	  tail

入队,  tail 的索引后移

	  front
		0  1  2  3  4  5  6  7  8  9
	                  tail

出队，front 的索引后移

	         front
			0  1  2  3  4  5  6  7  8  9
		                  tail

tail 在队列结尾，但是队列头还有可插入的位置

	                  front
			0  1  2  3  4  5  6  7  8  9
		                              tail

tail = (tail + 1 ) % size

		                  front
				0  1  2  3  4  5  6  7  8  9
	           tail

队列满时，(tail + 1) % size = front 表示队列满，浪费一个位置

	                       front
					0  1  2  3  4  5  6  7  8  9
	                  tail

	              front
					0  1  2  3  4  5  6  7  8  9
		                                      tail
*/
type LoopQueue[T common.Number] struct {
	data  []T // 队列数据
	front int // 首元素索引
	tail  int // 末元素后一个位置索引
	size  int // 队列长度
}

func (loopQueue *LoopQueue[T]) GetSize() int {
	return loopQueue.size
}

func (loopQueue *LoopQueue[T]) IsEmpty() bool {
	return loopQueue.front == loopQueue.tail
}

func (loopQueue *LoopQueue[T]) GetCapacity() int {
	return len(loopQueue.data) - 1
}

func (loopQueue *LoopQueue[T]) Enqueue(element T) error {
	// 如果队列已满
	if (loopQueue.tail+1)%len(loopQueue.data) == loopQueue.front {
		loopQueue.resize(loopQueue.GetCapacity() * 2)
	}
	// 添加元素
	loopQueue.data[loopQueue.tail] = element
	loopQueue.tail = (loopQueue.tail + 1) % len(loopQueue.data)

	loopQueue.size += 1
	return nil
}

func (loopQueue *LoopQueue[T]) Dequeue() T {
	if loopQueue.IsEmpty() {
		panic("Queue is empty")
	}
	element := loopQueue.data[loopQueue.front]
	// 循环队列需要执行求余运算
	loopQueue.size -= 1
	loopQueue.front = (loopQueue.front + 1) % len(loopQueue.data)

	// 缩容
	if loopQueue.size == loopQueue.GetCapacity()/4 && loopQueue.GetCapacity()/2 != 0 {
		loopQueue.resize(loopQueue.GetCapacity() / 2)
	}

	return element
}

func (loopQueue *LoopQueue[T]) GetFront() T {
	if loopQueue.IsEmpty() {
		panic("Queue is empty")
	}
	return loopQueue.data[loopQueue.front]
}

func (loopQueue *LoopQueue[T]) resize(newCapacity int) {
	newData := make([]T, newCapacity+1)
	// 将 front 位置放到 0 位置
	for i := 0; i < loopQueue.size; i++ {
		newData[i] = loopQueue.data[(i+loopQueue.front)%len(loopQueue.data)]
	}
	loopQueue.data = newData
	loopQueue.front = 0
	loopQueue.tail = loopQueue.size
}
func (loopQueue *LoopQueue[T]) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("LoopQueue: size = %d, capacity = %d\n", loopQueue.size, loopQueue.GetCapacity()))
	buffer.WriteString("front [")
	for i := loopQueue.front; i != loopQueue.tail; i = (i + 1) % len(loopQueue.data) {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprintf("%v", loopQueue.data[i]))
		if (i+1)%len(loopQueue.data) != loopQueue.tail {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")

	return buffer.String()
}

func NewLoopQueue[T common.Number](capacity int) *LoopQueue[T] {
	return &LoopQueue[T]{
		data:  make([]T, capacity+1),
		front: 0,
		tail:  0,
		size:  0,
	}
}
