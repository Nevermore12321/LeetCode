package main

import (
	"Algorithm_and_Data_Structure/Ch4_LinkList/LinkList"
	"fmt"
)

func main() {
	/* ==================LinkedList WithVirtualhead===================== */
	linkList := LinkList.NewLinkedListWithVirtualHead[int]()
	for i := 0; i < 5; i++ {
		err := linkList.AddFirst(i)
		if err != nil {
			panic(err)
		}
	}
	linkList.Add(666, 2)
	fmt.Println(linkList)

	linkList.Remove(2)
	fmt.Println(linkList)

	fmt.Println(linkList.RemoveFirst())
	fmt.Println(linkList)

	fmt.Println(linkList.RemoveLast())
	fmt.Println(linkList)
}
