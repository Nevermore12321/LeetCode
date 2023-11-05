package main

import (
	"Algorithm_and_Data_Structure/Ch3_DataStructure/Array"
	"Algorithm_and_Data_Structure/Ch3_DataStructure/Queue"
	"Algorithm_and_Data_Structure/Ch3_DataStructure/Stack"
	"fmt"
	"math/rand"
	"time"
)

// 测试使用 队列 q运行 opCount 个 enqueue 和 dequeue 的操作所需的时间
func testQueue(queue Queue.Queue[int], opCount int) {
	startTime := time.Now()
	for i := 0; i < opCount; i++ {
		err := queue.Enqueue(rand.Intn(10000))
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < opCount; i++ {
		queue.Dequeue()
	}

	duration := time.Since(startTime)
	fmt.Printf("%d elements, use: %v\n", opCount, duration)
}

func testStack(stack Stack.Stack[int], opCount int) {
	startTime := time.Now()
	for i := 0; i < opCount; i++ {
		stack.Push(rand.Intn(10000))
	}

	for i := 0; i < opCount; i++ {
		stack.Pop()
	}

	duration := time.Since(startTime)
	fmt.Printf("%d elements, use: %v\n", opCount, duration)
}

func main() {
	/* ================Array====================== */
	arr := Array.New[int](10)
	for i := 0; i < 10; i++ {
		err := arr.AddLast(i)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println(arr)
	err := arr.Add(1, 200)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(arr)
	err = arr.AddFirst(100)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(arr)

	fmt.Println()

	arr.Remove(2)
	fmt.Println(arr)

	fmt.Println()

	arr.RemoveElement(100)
	fmt.Println(arr)

	fmt.Println()

	arr.RemoveFirst()
	fmt.Println(arr)

	fmt.Println()

	/* ================ArrayStack====================== */
	arrayStack := Stack.ArrayStackNew[int](10)
	for i := 0; i < 5; i++ {
		arrayStack.Push(i)
		fmt.Println(arrayStack)
	}

	fmt.Println(arrayStack.Pop())
	fmt.Println(arrayStack)

	/* ================LinkedListStack====================== */
	linkedListStack := Stack.NewLinkedListStack[int]()
	for i := 0; i < 5; i++ {
		linkedListStack.Push(i)
		fmt.Println(linkedListStack)
	}

	fmt.Println(linkedListStack.Pop())
	fmt.Println(linkedListStack)

	/* ================Leetcode====================== */
	fmt.Println(Stack.IsValid("({})[{}]"))

	fmt.Println()

	/* ================ArrayQueue====================== */
	queue := Queue.NewArrayQueue[int](10)
	for i := 0; i < 10; i++ {
		err := queue.Enqueue(i)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(queue)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue)
	queue.Enqueue(11)
	fmt.Println(queue)

	/* ================LoopQueue====================== */
	loopQueue := Queue.NewLoopQueue[int](10)
	for i := 0; i < 11; i++ {
		err := loopQueue.Enqueue(i)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(loopQueue)
	fmt.Println(loopQueue.Dequeue())
	fmt.Println(loopQueue.Dequeue())
	fmt.Println(loopQueue.Dequeue())
	fmt.Println(loopQueue.Dequeue())
	fmt.Println(loopQueue.Dequeue())
	fmt.Println(loopQueue.Dequeue())
	fmt.Println(loopQueue.Dequeue())
	fmt.Println(loopQueue.Dequeue())
	fmt.Println(loopQueue)
	loopQueue.Enqueue(11)
	fmt.Println(loopQueue)

	fmt.Println()
	testaq := Queue.NewArrayQueue[int](10)
	testlq := Queue.NewLoopQueue[int](10)

	testQueue(testaq, 100000)
	testQueue(testlq, 100000)

	fmt.Println()
	testas := Stack.ArrayStackNew[int](10)
	testls := Stack.NewLinkedListStack[int]()

	testStack(testas, 100000)
	testStack(testls, 100000)

}
