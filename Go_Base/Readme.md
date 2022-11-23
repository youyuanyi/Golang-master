# GoLang基础

## Go的源码文件

![Go源码文件分类](https://raw.githubusercontent.com/youyuanyi/Golang-master/master/Go_Base/img/Go源码文件分类.png)

### 命令源码文件

用户创建的**.go源码文件**，命令源码文件被install后，GOPATH如果只有一个工作区，那么相应的可执行文件会被存放到当前工作区的bin文件夹下；如果有多个工作区，就会安装到GOBIN指向的目录

**多个命令源码文件可以分开用go run命令运行起来，但是无法通过go build和go install**

```bash
PS D:\workspace\Go\src\Golang-Master\Go_Base\code> dir


    目录: D:\workspace\Go\src\Golang-Master\Go_Base\code


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
-a----        2022/11/23      9:48             68 Hello.go
-a----        2022/11/23      9:49             68 Hello2.go


PS D:\workspace\Go\src\Golang-Master\Go_Base\code> go run .\Hello.go
hello_1
PS D:\workspace\Go\src\Golang-Master\Go_Base\code> go run .\Hello2.go
hello_2
PS D:\workspace\Go\src\Golang-Master\Go_Base\code> go build
# Golang-Master/Go_Base/code
.\Hello2.go:5:6: main redeclared in this block   // main包重复声明
        .\Hello.go:5:6: other declaration of main
PS D:\workspace\Go\src\Golang-Master\Go_Base\code> go install
# Golang-Master/Go_Base/code
.\Hello2.go:5:6: main redeclared in this block
        .\Hello.go:5:6: other declaration of main

```

### 库源码文件

库源码文件是存在于某个代码包中的普通的源码文件。库源码文件被install后，相应的归档文件**(.a文件)**会被存放到**$GOPAHT/pkg**目录下

### 测试源码文件

名称以_testgo为后缀的代码文件，必须包含Test或者Benchmark名称前缀的函数

```go
func TestXXX( t *testing.T) {

}
```

名称以 Test 为名称前缀的函数，只能接受 *testing.T 的参数，这种测试函数是功能测试函数。

```go
func BenchmarkXXX( b *testing.B) {

}
```

名称以 Benchmark 为名称前缀的函数，只能接受 *testing.B 的参数，这种测试函数是性能测试函数。



## Go的命令

### go run

go run 命令只能接受一个命令源码文件以及若干个库源码文件，不能接受测试源码文件。

一般用于调试程序

#### 执行流程

编译(命令源文件)->链接->生成一个**临时可执行文件**

### go build

go build主要用于测试编译

- 如果是普通包，go build不会产生任何文件
- 如果是main包，go build会在**当前目录下**生成一个可执行文件
- 如果是库源码文件，go build**只会测试编译包是否有问题**，不会产生任何文件
- 可以使用go build -o执行编译输出的文件名

### go install

用于**构建+安装包**

- 对**库源码文件**，go install会直接编译链接整个包，会在**$GOPATH/pkg目录下生成.a静态文件，供其他包调用** 
- 对**命令源码文件**，go install会执行**编译+链接+生成可执行文件**的操作，**生成的可执行文件在$GOPATH/bin目录下**

### go get

go get 命令用于从远程代码仓库（比如 Github ）上下载并安装代码包，默认安装路径为$GOPATH/src



## Golang数组与切片

### 数组

#### 声明方式

```go
var a [10]int
var b = [5]float32{100.0,2.1,3.3,4.0,8.5}
c:=[...]int{12,6,3}  // 将数组长度替换为...，由编译器负责找到长度
```

#### 遍历方式

##### for

```go
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
        fmt.Printf("%d th element of a is %.2f\n", i, a[i])
    }
}
```

##### for range

```go
package main

import "fmt"

func main() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    for i,v:=range a { //looping from 0 to the length of the array
        fmt.Printf("%d th element of a is %.2f\n", i, v)
    }
}
```



#### 注意事项

数组是值类型，赋值操作是**深拷贝**



### 切片

#### 实现原理

slice是一个可变长的数组，其底层结构是一个结构体

```go
type slice struct{
	array unsafe.pointer
	len int
	cap int
}
```

- array：一个数组指针，数据实际存储在该指针指向的数组上，占用8bytes
- len：当前slice中元素的个数，8bytes
- cap：slice的最大容量，8bytes

#### slice本质

slice本质不是什么动态数组，而是一个引用类型。之所以能像创建普通数组一样创建slice，是因为golang的语法糖。

#### slice共享存储空间

多个切片如果共享同一个底层数组，这种情况下，如果对一种一个切片或者底层数组修改，会影响到其他切片

```go
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
}

```

```bash
Output:
array before change 1 [78 79 80]
array after modification to slice nums1 [100 79 80]
array after modification to slice nums2 [100 101 80]
```



#### 切片常用操作

##### 创建

```go
// 1.直接声明
var slice1 []int

// 2.使用字面量
slice2 := []int{1,2,3,4,5}

// 3.使用make
slice3 := make([]int,3,5) //最大为5，当前为3的int slice

// 4.从slice或者数组中截取
s := [5]int{1,2,3,4,5}
slice4 := s[1:3]
slice5 := make([]int,len(slice4))
copy(slice5,slice4)
```



##### 增加

```go
var a []int
a = append(a,0)
a = append(a,1,2,3)
b := make([]int,len(a),(cap(a))*2) // b是a的两倍容量
copy(b,a)
```



##### 遍历

```go
slice1 := []int{1,2,3,4,5}
// 普通for循环遍历
for i:=0;i<len(slice1);i++{
	fmt.Println(slice1[i])
}
// for range遍历
for i,v := range slice1{
	fmt.Println(i,v)
}
```



##### 深拷贝

**深拷贝会在内存中开辟一个新的地址空间用来创建一个新对象**

数组、int、string、struct、float、bool等默认是深拷贝



##### 浅拷贝

浅拷贝只拷贝了数据的地址，只复制指向对象的指针，所以新对象和源对象指向的内存是一样的。新对象修改所指向内存的值时，源对象所指向内存的值也变化。

slice、map等引用类型默认为浅拷贝

```go
slice1 := []int{1, 2, 3, 4, 5}
slice2 := slice1
fmt.Printf("slice1 address:%p\n", slice1)  // slice1 address:0xc0000164e0
fmt.Printf("slice2 address:%p\n", slice2)  // slice2 address:0xc0000164e0

slice1[0] = 100
fmt.Println("slice1:", slice1) // slice1: [100 2 3 4 5]
fmt.Println("slice2:", slice2) // slice2: [100 2 3 4 5]
```



#### 扩容

设当前容量为x，所申请容量y，扩容后容量为c，原slice长度为l

- 如果y>2x，则c=y
- 如果l<1024，则c=2x
- 如果l>1024，则c=1.25x

#### 非线程安全

slice是非线性安全的，不支持并发读写。所以使用多个go routine对slice进行操作时，每次输出的值大概率会不一样

##### 加锁实现slice线程安全

```go
func TestSliceConcurrencySafeByMutex(t *testing.T) {
 var lock sync.Mutex //互斥锁
 a := make([]int, 0)
 var wg sync.WaitGroup
 for i := 0; i < 10000; i++ {
  wg.Add(1)
  go func(i int) {
   defer wg.Done()
   lock.Lock()
   defer lock.Unlock()
   a = append(a, i)
  }(i)
 }
 wg.Wait()
 t.Log(len(a)) 
 // equal 10000
}
```



##### 通过channel实现slice线程安全（推荐）

```go
func TestSliceConcurrencySafeByChanel(t *testing.T) {
 buffer := make(chan int)
 a := make([]int, 0)
 // 消费者
 go func() {
  for v := range buffer {
   a = append(a, v)
  }
 }()
 // 生产者
 var wg sync.WaitGroup
 for i := 0; i < 10000; i++ {
  wg.Add(1)
  go func(i int) {
   defer wg.Done()
   buffer <- i
  }(i)
 }
 wg.Wait()
 t.Log(len(a)) 
 // equal 10000
}
```

