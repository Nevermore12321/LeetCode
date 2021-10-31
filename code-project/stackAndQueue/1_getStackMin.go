package stackAndQueue

import (
	"fmt"
	"program-algorithm/lib"
)

/*
【题目】 实现一个特殊的栈，添加获取栈中最小值的功能
【要求】  Push Pop GetMin 的复杂度都为 O(1)

 */

type MinStack struct {
	stackData *lib.StackByLinkList
	stackMin *lib.StackByLinkList
}

func MakeMinStack() *MinStack {
	minStack := new(MinStack)
	minStack.stackMin = lib.MakeStackByLinkList()
	minStack.stackData = lib.MakeStackByLinkList()
	return minStack
}

func (ms *MinStack) Push(data interface{}) {
	ms.stackData.Push(data)

	if ms.stackMin.IsEmpty() {
		ms.stackMin.Push(data)
	} else if ms.stackMin.Peek().(int) >= data.(int) {
		ms.stackMin.Push(data)
	}
}

func (ms *MinStack) Pop() interface{} {
	if ms.stackData.IsEmpty() {
		fmt.Println("This stack is empty.")
		return nil
	}
	value, err := ms.stackData.Pop()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	top := ms.stackMin.Peek().(int)
	if value.(int) == top {
		_, err := ms.stackMin.Pop()
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
	}

	return value
}

func (ms *MinStack) GetMin() interface{} {
	if ms.stackMin.IsEmpty() {
		fmt.Println("This stack is empty.")
		return nil
	}

	return ms.stackMin.Peek()
}

func MinStackTest() {

	stack := MakeMinStack()
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(1)
	stack.Push(6)

	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
}