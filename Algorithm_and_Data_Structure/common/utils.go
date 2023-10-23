package common

func Swap[T comparable](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func CustomSwap[T CustomStruct](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
