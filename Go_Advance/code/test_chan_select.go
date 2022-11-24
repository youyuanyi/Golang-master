package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()
	select {
	case num1 := <-ch1:
		fmt.Println("从ch1中取得数据:", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("ch2中取数据:", num2)
		} else {
			fmt.Println("ch2通道已经关闭")
		}
	}

}
