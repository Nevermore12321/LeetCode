package stackAndQueue

import (
	"fmt"
	"program-algorithm/lib"
)

/*
【题目】：用两个栈实现队列的基本操作，支持 add，poll、peek 等
【要求】：
*/

type TwoStackQueue struct {
	stackPush *lib.StackByLinkList
	stackPop  *lib.StackByLinkList
}

func MakeTwoStackQueue() *TwoStackQueue {
	stack := new(TwoStackQueue)

	stackPush := lib.MakeStackByLinkList()
	stackPop := lib.MakeStackByLinkList()

	stack.stackPush = stackPush
	stack.stackPop = stackPop

	return stack
}

func (tsq *TwoStackQueue) Add(data interface{}) {
	tsq.stackPush.Push(data)
}

func (tsq *TwoStackQueue) Poll() interface{} {
	if tsq.stackPop.IsEmpty() && tsq.stackPush.IsEmpty() {
		fmt.Println("this queue is empty")
		return nil
	} else if tsq.stackPop.IsEmpty() {
		for !tsq.stackPush.IsEmpty() {
			value, err := tsq.stackPop.Pop()
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			tsq.stackPop.Push(value)
		}
	}
	popValue, err := tsq.stackPop.Pop()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return popValue
}

func (tsq *TwoStackQueue) Peek() interface{} {
	if tsq.stackPop.IsEmpty() && tsq.stackPush.IsEmpty() {
		fmt.Println("this queue is empty")
		return nil
	} else if tsq.stackPop.IsEmpty() {
		for !tsq.stackPush.IsEmpty() {
			value, err := tsq.stackPush.Pop()
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			tsq.stackPop.Push(value)
		}
	}

	peekValue := tsq.stackPop.Peek()
	return peekValue
}

func TwoStackQueueTest() {
	stack := MakeTwoStackQueue()

	stack.Add(1)
	stack.Add(2)
	stack.Add(3)
	stack.Add(4)
	stack.Add(5)

	fmt.Println(stack.Peek())
	fmt.Println(stack.Poll())
	fmt.Println(stack.Peek())
	fmt.Println(stack.Poll())
	fmt.Println(stack.Peek())
	fmt.Println(stack.Poll())

}