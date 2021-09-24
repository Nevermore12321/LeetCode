package lib

import (
	"errors"
	"fmt"
	"sync"
)

// SingleLinkListNode 节点结构
type SingleLinkListNode struct {
	Data interface{}
	Next *SingleLinkListNode
}

// SingleLinkList 单链表结构
type SingleLinkList struct {
	head *SingleLinkListNode
	mutx *sync.RWMutex
}

// MakeSingleLinkList 创建单链表初始化函数
func MakeSingleLinkList() *SingleLinkList {
	linkList := new(SingleLinkList)
	linkList.head = nil
	linkList.mutx = new(sync.RWMutex)
	return linkList
}

// IsEmpty 判断一个单链表是否为空
func (sl *SingleLinkList) IsEmpty() bool {
	return sl.head == nil
}

// Length 获取单链表的长度
// 注意，这里读数据，加读锁
func (sl *SingleLinkList) Length() int64 {
	current := sl.head
	var length int64 = 0

	//sl.mutx.RLock()
	//defer sl.mutx.RUnlock()

	for current != nil {
		length += 1
		current = current.Next
	}

	return length
}

// GetHead 获取单链表的头节点
func (sl *SingleLinkList) GetHead() *SingleLinkListNode {
	return sl.head
}

// GetNodeByIndex 获取指定位置节点，不存在返回nil
// index = 0 -> head 节点
func (sl *SingleLinkList) GetNodeByIndex(index int64) *SingleLinkListNode {
	if sl.head == nil {
		return nil
	}

	sl.mutx.RLock()
	defer sl.mutx.RUnlock()

	if index == 0 {
		return sl.head
	}

	var i int64 = 0
	current := sl.head

	for i = 0; i < index; i++ {
		current = current.Next
	}

	return current
}

// Add 从头部插入节点
func (sl *SingleLinkList) Add(data interface{}) {
	sl.mutx.Lock()
	defer sl.mutx.Unlock()

	tmp := new(SingleLinkListNode)
	tmp.Data = data
	tmp.Next = sl.head
	sl.head = tmp
}

// 从尾部插入节点
func (sl *SingleLinkList) Append(data interface{}) {
	sl.mutx.Lock()
	defer sl.mutx.Unlock()

	tmp := new(SingleLinkListNode)
	tmp.Data = data
	tmp.Next = nil

	if sl.IsEmpty() {
		sl.head = tmp
	} else {
		current := sl.head

		for current.Next != nil {
			current = current.Next
		}
		current.Next = tmp
	}
}

// Delete 删除 指定位置 的 节点
// 如果 index < 0 : 删除 头 节点
// 如果 index > length : 报错，超出范围
// 如果 index = [0, length] : 删除指定节点
func (sl *SingleLinkList) Delete(index int64) (*SingleLinkListNode, error) {
	if sl.IsEmpty() {
		return nil, errors.New("single list is nil, can't delete")
	}

	sl.mutx.Lock()
	defer sl.mutx.Unlock()

	current := sl.head
	if index <= 0 {
		sl.head = sl.head.Next
		return current, nil
	} else if index > sl.Length() {
		return nil, errors.New("index is out of bound")
	} else {
		var i int64 = 0
		for i = 0; i < index - 1; i ++ {
			current = current.Next
		}
		tmp := current.Next
		current.Next = current.Next.Next
		return tmp, nil
	}
}

// Show 单链表的打印显示
func (sl *SingleLinkList) Show() {
	if sl.head == nil {
		fmt.Println("this single link list is empty")
	} else {
		sl.mutx.RLock()
		defer sl.mutx.RUnlock()

		current := sl.head
		for current.Next != nil {
			fmt.Printf("%v -> ", current.Data)
			current = current.Next
		}
		fmt.Printf("%v\n", current.Data)
	}
}


// 测试函数

func SingleLinkTest() {
	linkList := MakeSingleLinkList()
	linkList.Append(9)
	linkList.Append(0)
	linkList.Append(6)
	linkList.Append(15)
	linkList.Append(2)
	linkList.Show()
	node := linkList.GetNodeByIndex(1)
	fmt.Println(node.Data)
	linkList.Append(4)
	linkList.Add(14)
	linkList.Show()

	deleteNode, err := linkList.Delete(1)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Println(deleteNode.Data)

	linkList.Show()

}