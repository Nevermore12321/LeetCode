package Stack

import (
	"Algorithm_and_Data_Structure/Ch4_LinkList/LinkList"
	"Algorithm_and_Data_Structure/common"
	"fmt"
	"strings"
)

type LinkedListStack[T common.Number] struct {
	data LinkList.LinkedListWithVirtualHead[T]
}

func (stack *LinkedListStack[T]) GetSize() int {
	return stack.data.GetSize()
}

func (stack *LinkedListStack[T]) IsEmpty() bool {
	return stack.data.IsEmpty()
}

func (stack *LinkedListStack[T]) Push(element T) {
	err := stack.data.AddFirst(element)
	if err != nil {
		panic(err)
	}
}

func (stack *LinkedListStack[T]) Pop() T {
	return stack.data.RemoveFirst()
}

func (stack *LinkedListStack[T]) Peek() T {
	return stack.data.GetFirst()
}

func (stack *LinkedListStack[T]) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("LinkedListStack: "))
	builder.WriteString(stack.data.String())
	return builder.String()
}

func NewLinkedListStack[T common.Number]() *LinkedListStack[T] {
	return &LinkedListStack[T]{
		data: *LinkList.NewLinkedListWithVirtualHead[T](),
	}
}
