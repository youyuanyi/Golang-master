package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	a := math.MaxFloat64
	b := math.Abs(a)
	pi := math.Pi
	fmt.Println(a, b, pi)

	// 随机数
	rand.Seed(time.Now().UnixMicro())
	r := rand.Int()
	c := rand.Intn(100)
	fmt.Println(r, c)

	f := rand.Float32()
	fmt.Println(f)
}
