package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

//type error interface {
//	Error() string
//}

// 结构体MyError实现了error接口的Error方法，也成为了一个error
func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
