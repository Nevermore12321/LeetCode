package sortAlgorithm

import (
	"fmt"
	"math/rand"
	"program-algorithm/lib"
	"time"
)

func __swap(s, d *int) {
	*s, *d = *d, *s
}

// 将数组切分为 arr[lo...j-1], a[j], a[j+1...hi]
func partition(arr []int, lo, hi int) int {
	// 选择第一个为 切分元素
	v := arr[lo]

	// 左右扫描指针
	i := lo + 1
	j := hi
	
	// 循环从两边找
	for true {

		// 从左向右扫描，找到左边第一个 >= v 的元素
		for i <= hi && arr[i] < v {
			i++
		}

		// 从右向左扫描，找到右边第一个 < v 的元素
		for j >= lo && arr[j] > v {
			j--
		}

		// 控制循环结束
		if i > j {
			break
		}

		//  现在 arr[i] > v arr[j] <= v ，交换后，继续上述步骤
		__swap(&arr[i], &arr[j])

		i++
		j--
	}

	// 跳出循环后，
	// i 表示 左边 + 1 元素 -> 也就是右边第一个元素
	// j 表示 右边 + 1 元素 -> 也就是左边最后一个元素
	__swap(&arr[j], &arr[lo])
	return j
}

func __quickSort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	j := partition(arr, lo, hi)
	__quickSort(arr, lo, j)
	__quickSort(arr, j+1, hi)
}

func QuickSort(arr []int) {
	__quickSort(arr, 0, len(arr) - 1)
}


func __insertionSortForQuick(arr []int, lo, hi int) {
	for i := lo; i <= hi; i++ {
		tmp := arr[i]
		j := i
		for ; j > lo && arr[j - 1] > tmp; j-- {
			arr[j] = arr[j - 1]
		}

		arr[j] = tmp
	}
}

func __partitionAdvanced1(arr []int, lo, hi int) int {
	// 随机选择一个数，并将与第一个位置的元素交换
	__swap(&arr[lo], &arr[rand.Intn(hi - lo) + lo])
	// 选择第一个为 切分元素
	v := arr[lo]

	// 左右扫描指针
	i := lo + 1
	j := hi

	// 循环从两边找
	for true {

		// 从左向右扫描，找到左边第一个 >= v 的元素
		for i <= hi && arr[i] < v {
			i++
		}

		// 从右向左扫描，找到右边第一个 < v 的元素
		for j >= lo && arr[j] > v {
			j--
		}

		// 控制循环结束
		if i > j {
			break
		}

		//  现在 arr[i] > v arr[j] <= v ，交换后，继续上述步骤
		__swap(&arr[i], &arr[j])

		i++
		j--
	}

	return j
}

// 采用随机切分
func __quickSortAdvanced1(arr []int, lo, hi int) {
	if lo - hi <= 15 {
		__insertionSortForQuick(arr, lo, hi)
		return
	}

	j := __partitionAdvanced1(arr, lo, hi)

	__quickSortAdvanced1(arr, lo, j - 1)
	__quickSortAdvanced1(arr, j + 1, hi)

}

func QuickSortAdvanced1(arr []int) {
	rand.Seed(time.Now().UnixNano())
	__quickSortAdvanced1(arr, 0, len(arr) - 1)
}


// ==============================

// 三取样切分算法，返回三取样切分元素索引
func __threeMedianIndex(arr []int, lo, hi int) int {
	//子数组少于3个元素时，第一个元素作为切分元素
	if hi - lo + 1 < 3 {
		return lo
	}

	//子数组有3个或以上元素时，取子数组前三个元素的中位数作为切分元素
	tmpArr := [3]int{lo, lo + 1, lo + 2}

	//使用插入排序法排序新数组b,按原数组的值进行排序。排序后的结果是原数组中小中大值对应的索引
	for i := 0; i < len(tmpArr); i++ {
		for j := i; j > 0; j-- {
			if arr[tmpArr[j]] < arr[tmpArr[j - 1]] {
				__swap(&tmpArr[j], &tmpArr[j - 1])
			}
		}
	}

	return tmpArr[1]
}

func __partitionAdvanced2(arr []int, lo, hi int) int {
	partMed := __threeMedianIndex(arr, lo, hi)
	__swap(&arr[partMed], &arr[lo])

	v := arr[lo]
	i := lo + 1
	j := hi

	for true {
		if i > j {
			break
		}

		for i <= hi && arr[i] < v {
			i++
		}
		for j >= lo && arr[j] > v {
			j--
		}

		__swap(&arr[i], &arr[j])
		i++
		j--
	}
	return j
}

func __quickSortAdvanced2(arr []int, lo, hi int) {
	if lo - hi <= 15 {
		__insertionSortForQuick(arr, lo, hi)
		return
	}

	j := __partitionAdvanced2(arr, lo, hi)

	__quickSortAdvanced2(arr, lo, j - 1)
	__quickSortAdvanced2(arr, j + 1, hi)
}

// 采用三取样切分
func QuickSortAdvanced2(arr []int) {
	rand.Seed(time.Now().UnixNano())
	__quickSortAdvanced2(arr, 0, len(arr) - 1)
}



func QuickSortTest() {
	var arr []int
	for i := 0; i < 100000; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	QuickSortAdvanced2(arr)
	fmt.Println(arr)

	lib.SortDuration(QuickSort, "QuickSort")
	lib.SortDuration(QuickSortAdvanced1, "QuickSortAdvanced1-random partition")
	lib.SortDuration(QuickSortAdvanced2, "QuickSortAdvanced2-three sample partition")
}
