package common

type ArrayGenerator struct {
}

func (ag ArrayGenerator) GeneratOrderedArray(n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, i)
	}
	return result

}
