package Student

type Student struct {
	Name string
	Age  int
	Id   int
}

func (s Student) Equals(target any) bool {
	if target == nil {
		return false
	}

	another := target.(Student)
	return another.Name == s.Name

}
