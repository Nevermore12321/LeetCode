package main

import (
	"fmt"
	"sync"
)

/*
使⽤两个 goroutine 交替打印序列，⼀个 goroutine 打印数字， 另外⼀个 goroutine 打印字⺟， 最终效果如下：
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

func printStr(sChan <-chan bool, done, nChan chan<- bool, wg *sync.WaitGroup) {
	str := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	i := 0
	for true {
		<- sChan
		if i >= len(str) {
			done <- true
			wg.Done()
			return
		}
		fmt.Print(str[i])
		fmt.Print(str[i+1])
		i += 2
		nChan <- true
	}

}

func printNum(nChan, done <-chan bool, sChan chan<- bool, wg *sync.WaitGroup) {
	i := 0
	for true {
		select {
		case <-nChan:
			fmt.Print(i)
			fmt.Print(i+1)
			i += 2
			sChan <- true
		case <-done:
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	letter, number, done := make(chan bool), make(chan bool), make(chan bool)
	wg.Add(2)
	go printNum(number, done, letter, &wg)
	go printStr(letter, done, number, &wg)

	number <- true
	wg.Wait()
}
