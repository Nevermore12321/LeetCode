package stackAndQueue

import (
	"fmt"
	"program-algorithm/lib"
)

/*
【题目】：仅使用递归函数和栈操作，实现逆序一个栈
【要求】：只能使用递归函数实现，不能使用其他数据结构
【示例】：一个栈依次压入 1，2，3，4，5，从栈顶到栈底分别是 5，4，3，2，1， 逆序后，从栈顶到栈底应该为 1，2，3，4，5
 */

func getAndRemoveLastElement(stack *lib.StackByLinkList) interface{} {
	res, err := stack.Pop()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if stack.IsEmpty() {
		return res
	} else {
		last := getAndRemoveLastElement(stack)
		stack.Push(res)
		return last
	}
}

func Reverse(stack *lib.StackByLinkList) {
	if stack.IsEmpty()  {
		return
	}

	cur_bottom := getAndRemoveLastElement(stack)
	Reverse(stack)
	stack.Push(cur_bottom)
}

func ReverseTest() {
	stack := lib.MakeStackByLinkList()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Show()

	Reverse(stack)

	stack.Show()
}



