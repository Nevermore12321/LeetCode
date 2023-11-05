package Stack

import "Algorithm_and_Data_Structure/common"

type Stack[T common.Number] interface {
	GetSize() int
	IsEmpty() bool
	Push(T)
	Pop() T
	Peek() T
}
