package main

import (
	"context"
	"fmt"
)

func main() {
	//  gen 函数，返回一个只读 chan，这个函数的作用就是，启动一个Goroutine，这个Goroutine中循环 写入 到chan，如果done chan 关闭了，就返回
	gen := func(ctx context.Context) <- chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				//  如果 context cancel，那么就goroutine结束
				case <- ctx.Done():
					return
				//  如果没有 cancel，那么就阻塞写入，直到写入成功
				case dst <- n:
					n ++
				}
			}
		}()
		return dst
	}

	//  创建 具有 cancel 的 Context
	ctx, cancel := context.WithCancel(context.Background())
	//  main 函数执行完后，调用 context 的 cancel 方法
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}