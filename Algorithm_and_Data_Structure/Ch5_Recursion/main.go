package main

import (
	"Algorithm_and_Data_Structure/Ch5_Recursion/LinkListR"
	"fmt"
)

func main() {
	linkList := LinkListR.NewLinkedListR[int]()
	for i := 0; i < 5; i++ {
		err := linkList.AddFirst(i)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(linkList)

	linkList.Add(666, 2)
	fmt.Println(linkList)

	fmt.Println(linkList.Remove(2))
	fmt.Println(linkList)

	fmt.Println(linkList.RemoveFirst())
	fmt.Println(linkList)

	fmt.Println(linkList.RemoveLast())
	fmt.Println(linkList)
}
