package Stack

/*
	括号匹配
	([{}])
	栈顶元素反映了在嵌套层次中的，最近的一个需要匹配的元素
*/

func IsValid(s string) bool {
	brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if len(stack) > 0 && brackets[c] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0

}
