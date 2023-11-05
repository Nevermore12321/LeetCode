package LinkList

import (
	"Algorithm_and_Data_Structure/common"
	"errors"
	"fmt"
)

type node[T common.Number] struct {
	data T
	next *node[T]
}

func (n node[T]) String() string {
	return fmt.Sprintf("%v", n.data)
}

func newNode[T common.Number]() *node[T] {
	return &node[T]{}
}

func newNodeWithElement[T common.Number](element T, next *node[T]) *node[T] {
	return &node[T]{
		data: element,
		next: next,
	}
}

/* ============== LinkedList ================== */

/*
	不使用虚拟头节点的链表，在往头部插入节点时，需要特殊处理
*/

type LinkedList[T common.Number] struct {
	head *node[T]
	size int
}

func (linkedList *LinkedList[T]) GetSize() int {
	return linkedList.size
}

func (linkedList *LinkedList[T]) IsEmpty() bool {
	return linkedList.size == 0
}

// AddFirst 链表头部插入元素
func (linkedList *LinkedList[T]) AddFirst(element T) {
	//n := newNode[T]()
	//n.data = element
	//n.next = linkedList.head
	//linkedList.head = n
	// 上面几行简化为
	linkedList.head = newNodeWithElement[T](element, linkedList.head)
	linkedList.size += 1
}

// Add 在链表 index 位置插入元素
func (linkedList *LinkedList[T]) Add(element T, index int) error {
	if index < 0 || index > linkedList.size {
		return errors.New("add failed. Illegal index")
	}

	// 如果是在 0 位置插入，也就是在头部插入
	if index == 0 {
		linkedList.AddFirst(element)
		return nil
	}

	// 找到 index 前一个位置的 node 节点 prev
	var prev *node[T] = linkedList.head
	for i := 0; i < index-1; i++ {
		prev = prev.next
	}

	//n := newNode[T]()
	//n.data = element
	//n.next = prev.next
	//prev.next = n
	// 上面逻辑可以简化为：
	prev.next = newNodeWithElement(element, prev.next)
	linkedList.size += 1
	return nil
}

// Add 在链表末尾插入元素
func (linkedList *LinkedList[T]) AddLast(element T) error {
	return linkedList.Add(element, linkedList.size)
}
func NewLinkedList[T common.Number]() *LinkedList[T] {
	return &LinkedList[T]{}
}
