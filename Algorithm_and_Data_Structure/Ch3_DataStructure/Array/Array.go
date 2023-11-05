package Array

import (
	"Algorithm_and_Data_Structure/common"
	"errors"
	"fmt"
	"strings"
)

type Array[T common.Number] struct {
	data []T
	size int
}

// New 初始化构造函数，出入容量
func New[T common.Number](capacity int) *Array[T] {
	return &Array[T]{
		data: make([]T, capacity),
		size: 0,
	}
}

// GetSize 获取已存元素的个数
func (array *Array[T]) GetSize() int {
	return array.size
}

// GetCapacity 获取数组的容量
func (array *Array[T]) GetCapacity() int {
	return len(array.data)
}

// IsEmpty 是否为空
func (array *Array[T]) IsEmpty() bool {
	return array.size == 0
}

// Add 在数组中的第 index 位置插入元素
func (array *Array[T]) Add(index int, element T) error {
	// 保证数组元素的紧密排列，不能跳过某些元素
	if index < 0 || index > array.size {
		return errors.New("add failed. require index >= 0 and index <= size")
	}

	// 扩容
	if array.size == len(array.data) {
		array.resize(2 * len(array.data))
	}

	// index 位置及以后的所有元素，后移一位
	for i := array.size - 1; i >= index; i-- {
		array.data[i+1] = array.data[i]
	}

	// 插入元素，放到 index 位置
	array.data[index] = element
	array.size += 1
	return nil
}

// AddLast 在数组的末尾插入元素
func (array *Array[T]) AddLast(element T) error {
	return array.Add(array.size, element)
}

// AddFirst 在数组的开头插入元素
func (array *Array[T]) AddFirst(element T) error {
	return array.Add(0, element)
}

// 数组的打印格式
func (array Array[T]) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", array.size, len(array.data)))
	builder.WriteString("[")
	for i := 0; i < array.size; i++ {
		builder.WriteString(fmt.Sprintf("%v", array.data[i]))
		if i != array.size-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("]\n")

	return builder.String()
}

// Get 获取 index 索引位置的元素
func (array *Array[T]) Get(index int) T {
	if index < 0 || index >= array.size {
		panic("Get failed. Index is illegal.")
	}
	return array.data[index]
}

func (array *Array[T]) GetLast() T {
	return array.Get(array.size - 1)
}

func (array *Array[T]) GetFirst() T {
	return array.Get(0)
}

// Set 修改 index 索引位置的元素
func (array *Array[T]) Set(index int, element T) {
	if index < 0 || index >= array.size {
		panic("Set failed. Index is illegal.")
	}
	array.data[index] = element
}

// Contains 查找数组中是否有此元素
func (array *Array[T]) Contains(element T) bool {
	for i := 0; i < array.size; i++ {
		if array.data[i] == element {
			return true
		}
	}
	return false
}

// Find 查找元素索引位置，如果存在返回索引位置，如果不存在返回 -1
func (array *Array[T]) Find(element T) int {
	for i := 0; i < array.size; i++ {
		if array.data[i] == element {
			return i
		}
	}
	return -1
}

// Remove 从数组中删除index位置的元素，并返回删除的元素
func (array *Array[T]) Remove(index int) T {
	if index < 0 || index > array.size {
		panic("remove failed. require index >= 0 and index <= size")
	}
	ret := array.data[index]

	for i := index + 1; i < array.size; i++ {
		array.data[i-1] = array.data[i]
	}
	array.size -= 1
	// 缩容
	if array.size == len(array.data)/4 && len(array.data)/2 != 0 {
		array.resize(len(array.data) / 2)
	}
	return ret
}

// RemoveFirst 删除第一个元素
func (array *Array[T]) RemoveFirst() T {
	return array.Remove(0)
}

// RemoveLast 删除最后一个元素
func (array *Array[T]) RemoveLast() T {
	return array.Remove(array.size - 1)
}

// 删除某一个值为 element 的元素
func (array *Array[T]) RemoveElement(element T) {
	index := array.Find(element)
	if index != -1 {
		array.Remove(index)
	}
}

func (array *Array[T]) resize(newCapacity int) {
	newData := make([]T, newCapacity)
	for i := 0; i < array.size; i++ {
		newData[i] = array.data[i]
	}
	array.data = newData
}
