package sortAlgorithm

import (
	"fmt"
	"math/rand"
	"time"
)

func DijkstraQuickSort(arr []int) {
	rand.Seed(time.Now().UnixNano())
	__DijkstraQuickSort(arr, 0, len(arr)-1)
}
func __DijkstraQuickSort(arr []int, lo, hi int) {
	// 数量小于 15 使用 插入排序
	if hi-lo <= 15 {
		__insertionSortForQuick(arr, lo, hi)
		return
	}

	//  随机切分
	randIndex := rand.Intn(hi-lo) + lo
	__swap(&arr[lo], &arr[randIndex])

	// 切分元素
	//v := arr[lo]
	v := arr[lo]

	//  三个指针初始化
	lt := lo
	gt := hi
	i := lo + 1

	for i <= gt {
		if arr[i] < v {
			__swap(&arr[lt], &arr[i])
			//arr[lt], arr[i] = arr[i], arr[lt]
			lt++
			i++
		} else if arr[i] > v {
			__swap(&arr[gt], &arr[i])
			//arr[gt], arr[i] = arr[i], arr[gt]
			gt--
		} else {
			i++
		}
	}

	// 去掉中间 =v 的部分，已经排好
	__DijkstraQuickSort(arr, lo, lt-1)
	__DijkstraQuickSort(arr, gt+1, hi)
}

func DijkstraQuickSortTest() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("三路快速排序耗时：")
	arr1 := []int{}
	for i := 0; i < 100000; i++ {
		arr1 = append(arr1, rand.Intn(100))
	}
	curTime := time.Now()
	DijkstraQuickSort(arr1)
	durTime := time.Now().Sub(curTime).Seconds()
	fmt.Printf("\t100000 elements and 100 random:  %v seconds \n", durTime)

	fmt.Println("普通快速排序耗时：")
	arr2 := []int{}
	for i := 0; i < 100000; i++ {
		arr2 = append(arr2, rand.Intn(10))
	}
	curTime = time.Now()
	QuickSortAdvanced1(arr2)
	durTime = time.Now().Sub(curTime).Seconds()
	fmt.Printf("\t100000 elements and 100 random:  %v seconds \n", durTime)
}
