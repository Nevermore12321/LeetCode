package lib

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	GenerateNearlyOrderedArray
@title    数组生成函数
@description   生成一个近乎有序的数组
@auth	Shaohe Guo
@param     length        int         "数组长度"
@param    swapTimes        int         "交换次数，也就是有序性，越小越有序，如果为0，表示完全有序"
@return		无序数组			[]int		"返回生成的无序数组"
*/
func GenerateNearlyOrderedArray(length, swapTimes int) []int {
	// 生成有序的数组
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = i
	}

	//  随机交换 swapTimes 次
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < swapTimes; i++ {
		posx := rand.Intn(length)
		posy := rand.Intn(length)
		arr[posx], arr[posy] = arr[posy], arr[posx]
	}

	return arr
}


/*
	SortDuration
@title    测试排序算法的时间
@description   随机生成无序数组，并测试排序算法时间
@auth	Shaohe Guo
@param     mySort        int         "排序算法的回调函数"
@param    SortName        int         "排序算法名称"
*/
func SortDuration(mySort func([]int), SortName string) {
	tinyArr := GenerateNearlyOrderedArray(1000, 1000)
	hugeArr := GenerateNearlyOrderedArray(100000, 10)
	hugaArrRand := GenerateNearlyOrderedArray(100000, 100000)

	fmt.Println(SortName, ":")

	curTime := time.Now()
	mySort(tinyArr)
	durTime := time.Now().Sub(curTime).Seconds()
	fmt.Printf("\t1000 elements and 1000 random:  %v seconds \n", durTime)

	curTime = time.Now()
	mySort(hugeArr)
	durTime = time.Now().Sub(curTime).Seconds()
	fmt.Printf("\t100000 elements and 10 random:  %v seconds \n", durTime)

	curTime = time.Now()
	mySort(hugaArrRand)
	durTime = time.Now().Sub(curTime).Seconds()
	fmt.Printf("\t100000 elements and 100000 random:  %v seconds \n", durTime)

}
