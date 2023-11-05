package Queue

import "Algorithm_and_Data_Structure/common"

// LoopQueueWithoutSize 在这一版LoopQueue的实现中，我们将不浪费那1个空间,即 size
type LoopQueueWithoutSize[T common.Number] struct {
	data  []T // 队列数据
	front int // 首元素索引
	tail  int // 末元素后一个位置索引
}

/*
如果队列还没有开始循环，front <= tail，队列长度就是 tail - front

	         front
			0  1  2  3  4  5  6  7  8  9
		                  tail

如果队列已经开始循环，front > tail，队列长度就是 len(data)-front + tail

		                  front
				0  1  2  3  4  5  6  7  8  9
	              tail
*/
func (loopQueue *LoopQueueWithoutSize[T]) GetSize() int {
	if loopQueue.front <= loopQueue.tail {
		return loopQueue.tail - loopQueue.front
	} else {
		return len(loopQueue.data) - loopQueue.front + loopQueue.tail
	}
}

// IsEmpty front==tail 表示队列为空
func (loopQueue *LoopQueueWithoutSize[T]) IsEmpty() bool {
	return loopQueue.tail == loopQueue.front
}

func (loopQueue *LoopQueueWithoutSize[T]) GetCapacity() int {
	return len(loopQueue.data)
}

func (loopQueue *LoopQueueWithoutSize[T]) Enqueue(element T) error {
	// 如果队列已满，扩容
	if (loopQueue.tail+1)%len(loopQueue.data) == loopQueue.front {
		loopQueue.resize(loopQueue.GetCapacity() * 2)
	}

	loopQueue.data[loopQueue.tail] = element
	loopQueue.tail = (loopQueue.tail + 1) % len(loopQueue.data)
	return nil
}

func (loopQueue *LoopQueueWithoutSize[T]) Dequeue() T {
	if loopQueue.IsEmpty() {
		panic("Queue is empty")
	}
	element := loopQueue.data[loopQueue.front]
	loopQueue.front = (loopQueue.front + 1) % len(loopQueue.data)

	// 如果当前队列长度为总容量的 1/4，那么缩容到 1/2
	if loopQueue.GetSize() == loopQueue.GetCapacity()/4 && loopQueue.GetCapacity()/2 != 0 {
		loopQueue.resize(loopQueue.GetCapacity() / 2)
	}
	return element
}

func (loopQueue *LoopQueueWithoutSize[T]) GetFront() T {
	if loopQueue.IsEmpty() {
		panic("Queue is empty")
	}
	return loopQueue.data[loopQueue.front]
}

func (loopQueue *LoopQueueWithoutSize[T]) resize(newCapacity int) {
	newData := make([]T, newCapacity+1)
	for i := 0; i < loopQueue.GetSize(); i++ {
		newData[i] = loopQueue.data[(i+loopQueue.front)%len(loopQueue.data)]
	}

	loopQueue.front = 0
	loopQueue.tail = loopQueue.GetSize()
	loopQueue.data = newData

}

func NewLoopQueueWithoutSize[T common.Number](capacity int) *LoopQueueWithoutSize[T] {
	return &LoopQueueWithoutSize[T]{
		data:  make([]T, capacity+1),
		front: 0,
		tail:  0,
	}

}
