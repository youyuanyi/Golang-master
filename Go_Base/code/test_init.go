package main

import "fmt"

var i int = initVar()

func initVar() int {
	fmt.Println("initVar...")
	return 100
}

func init() {
	fmt.Println("init...")
}

func main() {
	fmt.Println("main")
}
