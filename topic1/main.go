package main

import (
	"fmt"
	"sync"
)

/*
第一题
使用两个goroutine交替打印序列，一个goroutine打印数字，另外一个goroutine打印字母
最终效果如下：
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

/*
解题思路：
使用两个chan来控制数字、字母协程，两个协程初始化时都是处于阻塞状态，开始先给数字协程传入内容，打印数字，向字母协程传入内容，打印字母，直到字母大于Z。
*/

func main() {
	chanNum := make(chan struct{})
	chanStr := make(chan struct{})
	wg := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-chanNum:
				fmt.Print(i)
				i++
				fmt.Print(i)
				chanStr <- struct{}{}
				i++
			}
		}
	}()
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-chanStr:
				if i >= 'Z' {
					wg.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				chanNum <- struct{}{}
			}
		}
	}(&wg)
	chanNum <- struct{}{}
	wg.Wait()
}
