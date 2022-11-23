package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Tom"
	s2 := "Jerry"
	// 字符串连接
	s3 := s1 + s2
	fmt.Printf("s3:%s\n", s3)

	// 使用Join完成字符串连接
	s4 := strings.Join([]string{s1, s2}, ",")
	fmt.Printf("s4:%s\n", s4)

	// 字符串切片:左闭右开原则
	fmt.Println("s4[1:3]:", s4[1:3])

	// 字符串分割
	s5 := strings.Split(s4, ",")
	fmt.Println("s5:", s5)

	// 是否包含某个字符串
	fmt.Println("s3.Contains('Tom')? ", strings.Contains(s3, "Tom"))

	// 字符串都转为小写
	s6 := strings.ToLower(s3)
	fmt.Println("s6:", s6)

	// 查找前后缀
	fmt.Println("s3.Prefix('Tom')", strings.HasPrefix(s3, "Tom"))
	fmt.Println("s3.Suffix('Jerry')", strings.HasSuffix(s3, "Jerry"))

	// 查找字符串中指定字符/子串的首次出现位置
	fmt.Println("s3.Index('er')", strings.Index(s3, "er"))

}
