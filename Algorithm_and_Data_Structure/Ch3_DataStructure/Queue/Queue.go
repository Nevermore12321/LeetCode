package Queue

import "Algorithm_and_Data_Structure/common"

type Queue[T common.Number] interface {
	GetSize() int
	IsEmpty() bool
	Enqueue(T) error
	Dequeue() T
	GetFront() T
}
