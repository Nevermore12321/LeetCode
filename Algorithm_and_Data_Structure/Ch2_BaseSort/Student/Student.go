package Student

type Student struct {
	Name  string
	Id    int
	Age   int
	Score int
}

// CompareTo 自定义结构体的比较方法
// 0 相等，负数 比another小， 正数 比another大
func (s Student) CompareTo(another any) int {
	other := another.(Student)
	return s.Score - other.Score
}
