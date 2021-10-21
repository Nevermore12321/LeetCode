package main

import (
	"fmt"
	"strings"
)

/*
给定两个字符串，请编写程序，确定其中⼀个字符串的字符重新排列后，能否变成另⼀个字符串。
这⾥规定【⼤⼩写为不同字符】，且考虑字符串重点空格。

给定⼀个strings1和⼀个string s2，请返回⼀个bool，代表两串是否重新排列后可相同。 保证两串的⻓度都⼩于等于5000。
 */

func isSameGroup(str1, str2 string) bool {
	l1 := len(str1)
	l2 := len(str2)

	if l1 > 5000 || l2 > 5000 || l1 != l2 {
		return false
	}

	for _, item := range str1 {
		if strings.Count(str1, string(item)) != strings.Count(str2, string(item)) {
			return false
		}
	}
	return true
}

func main() {
	str1 := "asdfg 24fcasde"
	str2 := "2a des4acdsffg"
	fmt.Println(isSameGroup(str1,str2))
}
