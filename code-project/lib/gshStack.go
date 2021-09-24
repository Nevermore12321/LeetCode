package lib

import (
	"errors"
	"fmt"
	"reflect"
)

// StackByLinkList 单链表实现栈 的结构
type StackByLinkList struct {
	stack *SingleLinkList
}

// MakeStackByLinkList 栈的创建
func MakeStackByLinkList() *StackByLinkList {
	stack := new(StackByLinkList)
	singleList := MakeSingleLinkList()

	stack.stack = singleList
	return stack
}

// IsEmpty 判断 stack 是否为空
func (stack *StackByLinkList) IsEmpty() bool {
	return stack.stack.IsEmpty()
}

// Push 进栈
//  进栈，也就是不停的 从 头 进栈
//  出栈，也是 从 头 出栈
func (stack *StackByLinkList) Push(data interface{}) {
	stack.stack.Add(data)
}

// Pop 出栈
func (stack *StackByLinkList) Pop() (interface{}, error){
	popNode := stack.stack.GetHead()

	if popNode == nil {
		err := errors.New("stack is empty, can't pop node")
		return nil, err
	} else {
		_, err := stack.stack.Delete(0)
		if err != nil {
			return nil, err
		}
		return popNode.Data, nil
	}
}

// StackLen 栈长度
func (stack *StackByLinkList) StackLen() int64 {
	var length int64 = 0
	current := stack.stack.head

	for current.Next != nil {
		length++
		current = current.Next
	}
	return length
}

// Clear 清空栈
func (stack *StackByLinkList) Clear() {
	stack.stack = MakeSingleLinkList()
}

// 打印栈
func (stack *StackByLinkList) Show() {
	if stack.stack.IsEmpty() {
		fmt.Println("The stack is empty")
	} else {
		current := stack.stack.head
		for current.Next != nil {
			fmt.Printf("%v -> ", current.Data)
			current = current.Next
		}
		fmt.Printf("%v\n", current.Data)
	}
}

// Traverser 遍历，
//  fn 表示 对每个 栈元素 的处理函数
//  isBottomToTop : 默认是 自 顶 向 下，如果isBottomToTop设置为 true，则是 自 下 向 上
func (stack *StackByLinkList) Traverser(fn func(data interface{})) {
	//  正序
	current := stack.stack.head
	for ; current != nil; current = current.Next {
		fn(current.Data)
	}
}

//  获取 栈顶 元素
func (stack *StackByLinkList) Peek() interface{} {
	if !stack.stack.IsEmpty() {
		return stack.stack.head.Data
	} else {
		return nil
	}
}

// Search 搜索某一个item在该栈中的位置, 位置为离栈顶最近的item与栈顶间距离
//  item : 要搜索的 元素 为 item
//  返回 搜索到的值 距离 栈顶的 距离
func (stack *StackByLinkList) Search(item interface{}) int64 {
	if stack.stack.IsEmpty() {
		fmt.Println("Stack is Null")
		return 0
	}
	var index int64 = 1
	current := stack.stack.head
	for ; current != nil; current = current.Next {
		if reflect.DeepEqual(item, current.Data) {
			return index
		}
		index++
	}

	return 0
}

// StackTest 测试函数
func StackTest() {
	stack := MakeStackByLinkList()
	stack.Push(3)
	stack.Push(5)
	stack.Push(9)
	stack.Push(2)
	stack.Push(1)
	stack.Push(8)
	stack.Show()
	fmt.Println(stack.Pop())
	stack.Show()
	fmt.Println(stack.Peek())
}