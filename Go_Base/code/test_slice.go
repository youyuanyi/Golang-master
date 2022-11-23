package main

import (
	"fmt"
)

func main() {
	s := [3]int{78, 79, 80}
	nums1 := s[:]
	nums2 := s[:] //多个切片共享同一个底层数组
	fmt.Println("array before change 1", s)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", s)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", s)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1
	fmt.Printf("slice1 address:%p\n", slice1)
	fmt.Printf("slice2 address:%p\n", slice2)

	slice1[0] = 100
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
}
