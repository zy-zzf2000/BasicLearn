package channel

import (
	"strconv"
	"time"
)

//返回一个生成1，2，3，...字符串的生成器
func Getgenerator() <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(time.Second))
			ch <- strconv.Itoa(i)
		}
	}()
	return ch
}

//多路复用函数，将两个Generator的channel合并为一个channel，防止了ch1与ch2的相互阻塞
func Merge(ch1 <-chan string, ch2 <-chan string) <-chan string {

	ch := make(chan string)

	go func() {
		for {
			ch <- <-ch1
		}
	}()
	go func() {
		for {
			ch <- <-ch2
		}
	}()

	return ch
}
