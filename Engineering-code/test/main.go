package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct {}

func (s *Student) Speak(str string) string {
	fmt.Println(str)
	return str + "new"
}

func main() {
	var p People = &Student{}

	bbb := p.Speak("aaa")
	fmt.Println(bbb)
}