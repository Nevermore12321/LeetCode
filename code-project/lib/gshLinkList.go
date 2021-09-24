package lib

import (
	"fmt"
	"sync"
)

// LinkListNode 双向链表节点
type LinkListNode struct {
	Data interface{}
	Prev, Next *LinkListNode
}


// LinkList 双向链表结构
type LinkList struct {
	length 		int64						// 双向链表的长度
	Head 		*LinkListNode				// 双向链表的头节点
	Tail		*LinkListNode				// 双向链表的尾节点
	mutex		*sync.RWMutex				// 多线程安全
}

// MakeLinkList 创建 双向链表
func MakeLinkList() *LinkList {
	mylink := new(LinkList)
	mylink.length = 0
	mylink.Head = nil
	mylink.Tail = nil
	mylink.mutex = new(sync.RWMutex)
	return mylink
}

// GetByIndex Get 获取指定位置的节点
//  index : 0  表示 返回 head 节点
//  index : [1, ] 表示 head 下一个表示1
func (ll *LinkList) GetByIndex(index int64) *LinkListNode {
	if index < 0 || index > ll.length {
		return nil
	}

	if index == 0 {
		return ll.Head
	}

	var i int64 = 0
	current := ll.Head
	for ; i < index; i++ {
		current = current.Next
	}

	return current
}

// Append 在双向链表结尾追加节点
func (ll *LinkList) Append(node *LinkListNode) bool {
	if node == nil {
		return false
	}

	ll.mutex.Lock()
	defer ll.mutex.Unlock()

	if ll.length == 0 {
		ll.Head = node
		ll.Tail = node
		node.Prev = nil
		node.Next = nil
	} else {
		ll.Tail.Next = node
		node.Prev = ll.Tail
		node.Next = nil
		ll.Tail = node
	}

	ll.length++
	return true
}

// InsertAt 向双链表指定位置插入节点
func (ll *LinkList) InsertAt(index int64, node *LinkListNode) bool {
	if node == nil || index > ll.length || index < 0 {
		return false
	}

	if index == ll.length {
		return ll.Append(node)
	}

	ll.mutex.Lock()
	defer ll.mutex.Unlock()

	if index == 0 {
		ll.Head.Prev = node
		node.Next = ll.Head
		node.Prev = nil
		ll.Head = node
	} else {
		indexNode := ll.GetByIndex(index)
		indexNode.Prev.Next = node
		node.Prev = indexNode.Prev
		indexNode.Prev = node
		node.Next = indexNode
	}
	ll.length++
	return true
}

// DeleteAt 删除指定位置的节点
// 0 是 head 节点，1 是head的下一个节点
func (ll *LinkList) DeleteAt(index int64) bool {
	if index < 0 || index > ll.length || ll.length == 0 {
		return false
	}

	ll.mutex.Lock()
	defer ll.mutex.Unlock()

	if index == 0 {
		if ll.length == 1 {
			ll.Head = nil
			ll.Tail = nil
		} else {
			ll.Head = ll.Head.Next
			ll.Head.Prev = nil
		}

		ll.length--
		return true
	}

	if index == ll.length - 1 {
		ll.Tail = ll.Tail.Prev
		ll.Tail.Next = nil
		ll.length--
		return true
	}

	indexNode := ll.GetByIndex(index)
	indexNode.Prev.Next = indexNode.Next
	indexNode.Next.Prev = indexNode.Prev

	ll.length--
	return true
}

// Show 显示双向链表
func (ll *LinkList) Show() {
	if ll == nil || ll.length == 0 {
		fmt.Println("this double list is nil or empty")
		return
	}

	ll.mutex.RLock()
	defer ll.mutex.RUnlock()

	current := ll.Head
	for current.Next != nil {
		fmt.Printf("%v <-> ", current.Data)
		current = current.Next
	}
	fmt.Printf("%v\n", current.Data)
}

//  ReverseShow 倒序打印双链表信息
func (ll *LinkList) ReverseShow() {
	if ll == nil || ll.length == 0 {
		fmt.Println("this double list is nil or empty")
		return
	}

	ll.mutex.RLock()
	defer ll.mutex.RUnlock()

	current := ll.Tail
	for current.Prev != nil {
		fmt.Printf("%v <-> ", current.Data)
		current = current.Prev
	}
	fmt.Printf("%v\n", current.Data)
}


// 测试函数

func LinkListTest() {
	node1 := new(LinkListNode)
	node1.Data = 9
	node2 := new(LinkListNode)
	node2.Data = 0
	node3 := new(LinkListNode)
	node3.Data = 6
	node4 := new(LinkListNode)
	node4.Data = 15
	node5 := new(LinkListNode)
	node5.Data = 2
	node6 := new(LinkListNode)
	node6.Data = 1
	linkList := MakeLinkList()
	linkList.Append(node1)
	linkList.Append(node2)
	linkList.Append(node3)
	linkList.Append(node4)
	linkList.Append(node5)
	linkList.Append(node6)
	linkList.Show()

	index := linkList.GetByIndex(4)
	fmt.Println(index.Data)

	node7 := new(LinkListNode)
	node7.Data = 10
	node8 := new(LinkListNode)
	node8.Data = 10
	linkList.InsertAt(0, node7)
	linkList.Show()
	linkList.InsertAt(5, node8)
	linkList.Show()

	linkList.DeleteAt(0)
	linkList.Show()
	linkList.DeleteAt(3)
	linkList.Show()

}