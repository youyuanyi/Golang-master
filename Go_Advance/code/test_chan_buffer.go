package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch3 := make(chan string, 4) // buffer为4
	go sendData3(ch3)

	for v := range ch3 {
		fmt.Println("\t读取的数据是:", v)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("main over")
}
func sendData3(ch chan string) {
	for i := 0; i < 10; i++ { // ch的buffer为4，子goroutine可连续往ch里传入4个数据
		ch <- "数据" + strconv.Itoa(i)
		fmt.Println("子goroutine,写出第", i, "个数据")
	}
	close(ch)
}
