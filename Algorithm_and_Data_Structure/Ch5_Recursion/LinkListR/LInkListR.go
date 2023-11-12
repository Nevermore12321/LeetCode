package LinkListR

import (
	"Algorithm_and_Data_Structure/common"
	"bytes"
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

/* ============== LinkedListR ================== */

/*
	递归模式
*/

type LinkedListR[T common.Number] struct {
	head *node[T]
	size int
}

func (linkedList *LinkedListR[T]) GetSize() int {
	return linkedList.size
}

func (linkedList *LinkedListR[T]) IsEmpty() bool {
	return linkedList.size == 0
}

// AddFirst 链表头部插入元素
func (linkedList *LinkedListR[T]) AddFirst(element T) error {
	return linkedList.Add(element, 0)
}

// Add 在链表 index 位置插入元素
func (linkedList *LinkedListR[T]) Add(element T, index int) error {
	if index < 0 || index > linkedList.size {
		return errors.New("add failed. Illegal index")
	}
	linkedList.head = linkedList.add(linkedList.head, element, index)
	linkedList.size += 1
	return nil
}

func (linkedList *LinkedListR[T]) add(head *node[T], element T, index int) *node[T] {
	if index == 0 {
		return newNodeWithElement[T](element, head)
	}
	head.next = linkedList.add(head.next, element, index-1)
	return head
}

// Get 获取 index 位置的元素
func (linkedList *LinkedListR[T]) Get(index int) T {
	if index < 0 || index >= linkedList.size {
		panic("get failed. Illegal index")
	}
	return linkedList.get(linkedList.head, index)
}

func (linkedList *LinkedListR[T]) get(head *node[T], index int) T {
	if index == 0 {
		return head.data
	}
	return linkedList.get(head.next, index-1)
}

// GetFirst 获取链表的第一个元素
func (linkedList *LinkedListR[T]) GetFirst() T {
	return linkedList.Get(0)
}

// GetLast 获取链表的第一个元素
func (linkedList *LinkedListR[T]) GetLast() T {
	return linkedList.Get(linkedList.size - 1)
}

// AddLast 在链表末尾插入元素
func (linkedList *LinkedListR[T]) AddLast(element T) error {
	return linkedList.Add(element, linkedList.size-1)
}

// Set 修改链表第 index 元素的值
func (linkedList *LinkedListR[T]) Set(element T, index int) error {
	if index < 0 || index >= linkedList.size {
		return errors.New("set failed. Illegal index")
	}
	linkedList.set(linkedList.head, element, index)
	return nil
}

func (linkedList *LinkedListR[T]) set(head *node[T], element T, index int) {
	if index == 0 {
		head.data = element
		return
	}
	linkedList.set(head.next, element, index-1)
}

// Contains 链表中是否包含元素 element
func (linkedList *LinkedListR[T]) Contains(element T) bool {
	return linkedList.contains(linkedList.head, element)
}

func (linkedList *LinkedListR[T]) contains(head *node[T], element T) bool {
	if head == nil {
		return false
	}
	if head.data == element {
		return true
	}
	return linkedList.contains(head.next, element)
}

func (linkedList *LinkedListR[T]) String() string {
	buffer := bytes.Buffer{}
	current := linkedList.head
	for current != nil {
		buffer.WriteString(fmt.Sprint(current.data) + "->")
		current = current.next
	}

	buffer.WriteString("NULL\n")

	return buffer.String()
}

// Remove 删除 index 位置的元素
func (linkedList *LinkedListR[T]) Remove(index int) T {
	if index < 0 || index >= linkedList.size {
		panic("remove failed. Illegal index")
	}
	head, ele := linkedList.remove(linkedList.head, index)
	linkedList.head = head
	linkedList.size -= 1
	return ele
}

func (linkedList *LinkedListR[T]) remove(head *node[T], index int) (*node[T], T) {
	if index == 0 {
		return head.next, head.data
	}
	nextNode, ele := linkedList.remove(head.next, index-1)
	head.next = nextNode
	return head, ele
}

// RemoveFirst 删除链表中第一个元素
func (linkedList *LinkedListR[T]) RemoveFirst() T {
	return linkedList.Remove(0)
}

// RemoveLast 删除链表中最后一个元素
func (linkedList *LinkedListR[T]) RemoveLast() T {
	return linkedList.Remove(linkedList.size - 1)
}

func NewLinkedListR[T common.Number]() *LinkedListR[T] {
	return &LinkedListR[T]{}
}
