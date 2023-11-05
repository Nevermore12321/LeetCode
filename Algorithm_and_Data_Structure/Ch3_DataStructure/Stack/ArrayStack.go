package Stack

import (
	"Algorithm_and_Data_Structure/Ch3_DataStructure/Array"
	"Algorithm_and_Data_Structure/common"
	"fmt"
	"strings"
)

type ArrayStack[T common.Number] struct {
	array Array.Array[T]
}

func ArrayStackNew[T common.Number](capacity int) *ArrayStack[T] {
	return &ArrayStack[T]{
		array: *Array.New[T](capacity),
	}
}

func (arrayStack *ArrayStack[T]) GetSize() int {
	return arrayStack.array.GetSize()
}

func (arrayStack *ArrayStack[T]) IsEmpty() bool {
	return arrayStack.array.IsEmpty()
}

func (arrayStack *ArrayStack[T]) GetCapacity() int {
	return arrayStack.array.GetCapacity()
}

func (arrayStack *ArrayStack[T]) Push(element T) {
	err := arrayStack.array.AddLast(element)
	if err != nil {
		panic(err)
	}
}

func (arrayStack *ArrayStack[T]) Pop() T {
	return arrayStack.array.RemoveLast()
}

func (arrayStack *ArrayStack[T]) Peek() T {
	return arrayStack.array.GetLast()
}

func (arrayStack *ArrayStack[T]) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Stack: "))
	builder.WriteString("[")
	for i := 0; i < arrayStack.array.GetSize(); i++ {
		builder.WriteString(fmt.Sprintf("%v", arrayStack.array.Get(i)))
		if i != arrayStack.array.GetSize()-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("] top \n")

	return builder.String()
}
