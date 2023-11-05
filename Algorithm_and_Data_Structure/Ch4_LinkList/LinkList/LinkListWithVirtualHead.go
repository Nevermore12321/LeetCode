package LinkList

import (
	"Algorithm_and_Data_Structure/common"
	"bytes"
	"errors"
	"fmt"
)

/* ============== LinkedListWithVirtualHead ================== */

/*
	使用虚拟头节点，将链表的所有操作，统一起来，不需要对在头部插入节点时特殊处理。
*/

type LinkedListWithVirtualHead[T common.Number] struct {
	dummyHead *node[T]
	size      int
}

func (linkedList *LinkedListWithVirtualHead[T]) GetSize() int {
	return linkedList.size
}

func (linkedList *LinkedListWithVirtualHead[T]) IsEmpty() bool {
	return linkedList.size == 0
}

// AddFirst 链表头部插入元素
func (linkedList *LinkedListWithVirtualHead[T]) AddFirst(element T) error {
	return linkedList.Add(element, 0)
}

// Add 在链表 index 位置插入元素
func (linkedList *LinkedListWithVirtualHead[T]) Add(element T, index int) error {
	if index < 0 || index > linkedList.size {
		return errors.New("add failed. Illegal index")
	}

	// 找到 index 前一个位置的 node 节点 prev
	var prev *node[T] = linkedList.dummyHead
	for i := 0; i < index; i++ {
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

// Get 获取 index 位置的元素
func (linkedList *LinkedListWithVirtualHead[T]) Get(index int) T {
	if index < 0 || index >= linkedList.size {
		panic("get failed. Illegal index")
	}
	current := linkedList.dummyHead.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.data
}

// GetFirst 获取链表的第一个元素
func (linkedList *LinkedListWithVirtualHead[T]) GetFirst() T {
	return linkedList.Get(0)
}

// GetLast 获取链表的第一个元素
func (linkedList *LinkedListWithVirtualHead[T]) GetLast() T {
	return linkedList.Get(linkedList.size - 1)
}

// AddLast 在链表末尾插入元素
func (linkedList *LinkedListWithVirtualHead[T]) AddLast(element T) error {
	return linkedList.Add(element, linkedList.size-1)
}

// Set 修改链表第 index 元素的值
func (linkedList *LinkedListWithVirtualHead[T]) Set(element T, index int) error {
	current := linkedList.dummyHead.next
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.data = element
	return nil
}

// Contains 链表中是否包含元素 element
func (linkedList *LinkedListWithVirtualHead[T]) Contains(element T) bool {
	current := linkedList.dummyHead.next
	for current != nil {
		if current.data == element {
			return true
		}
		current = current.next
	}
	return false
}

func (linkedList *LinkedListWithVirtualHead[T]) String() string {
	buffer := bytes.Buffer{}
	current := linkedList.dummyHead.next
	for current != nil {
		buffer.WriteString(fmt.Sprint(current.data) + "->")
		current = current.next
	}

	buffer.WriteString("NULL")

	return buffer.String()
}

// Remove 删除 index 位置的元素
func (linkedList *LinkedListWithVirtualHead[T]) Remove(index int) T {
	if index < 0 || index >= linkedList.size {
		panic("remove failed. Illegal index")
	}
	prev := linkedList.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	element := prev.next.data
	prev.next = prev.next.next
	linkedList.size -= 1
	return element
}

// RemoveFirst 删除链表中第一个元素
func (linkedList *LinkedListWithVirtualHead[T]) RemoveFirst() T {
	return linkedList.Remove(0)
}

// RemoveLast 删除链表中最后一个元素
func (linkedList *LinkedListWithVirtualHead[T]) RemoveLast() T {
	return linkedList.Remove(linkedList.size - 1)
}

func NewLinkedListWithVirtualHead[T common.Number]() *LinkedListWithVirtualHead[T] {
	return &LinkedListWithVirtualHead[T]{
		dummyHead: newNode[T](),
		size:      0,
	}
}
