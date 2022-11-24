package main

import (
	"fmt"
	"time"
)

func sendData(ch1 chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch1 <- i
	}
	close(ch1) //关闭通道
}

func main() {
	ch1 := make(chan int)
	go sendData(ch1)
	for v := range ch1 {
		fmt.Println("读取数据:", v)
	}
	fmt.Println("main over..")
}
