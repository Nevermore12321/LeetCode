package main

import "fmt"

/*
请实现⼀个算法，在不使⽤【额外数据结构和储存空间】的情况下，翻转⼀个给定的字符串(可以使⽤单个过程变量)。
给定⼀个string，请返回⼀个string，为翻转后的字符串。保证字符串的⻓度⼩于等于 5000
 */


func reverseStr(str string) (string, bool){
	runeStr := []rune(str)
	l := len(runeStr)

	if len(runeStr) > 5000 {
		return "", false
	}

	for i := 0; i < l/2; i++ {
		runeStr[i], runeStr[l - 1 - i] = runeStr[l - 1 - i], runeStr[i]
	}

	return string(runeStr), true
}


func main() {
	str := "abcdefghigklmn"
	rStr, flag := reverseStr(str)
	if flag {
		fmt.Println(rStr)
	}
}
