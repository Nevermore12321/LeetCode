package Queue

import (
	"Algorithm_and_Data_Structure/Ch3_DataStructure/Array"
	"Algorithm_and_Data_Structure/common"
	"fmt"
	"strings"
)

/*
数组实现队列，有一个小问题，虽然入队的时间复杂度为 O(1)，但出队的时间复杂度为 O(n)
*/
type ArrayQueue[T common.Number] struct {
	array Array.Array[T]
}

func NewArrayQueue[T common.Number](capacity int) *ArrayQueue[T] {
	return &ArrayQueue[T]{
		array: *Array.New[T](capacity),
	}
}

func (arrayQueue *ArrayQueue[T]) GetSize() int {
	return arrayQueue.array.GetSize()
}

func (arrayQueue *ArrayQueue[T]) IsEmpty() bool {
	return arrayQueue.array.IsEmpty()
}

func (arrayQueue *ArrayQueue[T]) GetCapacity() int {
	return arrayQueue.array.GetCapacity()
}

func (arrayQueue *ArrayQueue[T]) Enqueue(element T) error {
	return arrayQueue.array.AddLast(element)
}

func (arrayQueue *ArrayQueue[T]) Dequeue() T {
	return arrayQueue.array.RemoveFirst()
}

func (arrayQueue *ArrayQueue[T]) GetFront() T {
	return arrayQueue.array.GetFirst()
}

func (arrayQueue *ArrayQueue[T]) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("ArrayQueue: "))
	builder.WriteString("[")
	for i := 0; i < arrayQueue.array.GetSize(); i++ {
		builder.WriteString(fmt.Sprintf("%v", arrayQueue.array.Get(i)))
		if i != arrayQueue.array.GetSize()-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("] tail\n")

	return builder.String()
}
