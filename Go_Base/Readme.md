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



## Golang 数组与切片

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



## Golang Map

在golang中，map将一个key与一个value关联起来，其**底层实现为Hash表，所以是无序的**。

### 实现原理

- Golang中的**map是一个指针**，占用8 bytes，**指向hmap结构体**
- 每个map的底层结构是hmap，**hmap包含若干个结构为bmap的bucket数组**
- **每个bucket底层采用链表结构**

#### hmap结构体

```go
type hmap struct {
    // 代表哈希表中的元素个数，调用len(map)时，返回的就是该字段值。
    count     int 
     // 状态标志，下文常量中会解释四种状态位含义。
    flags     uint8 
    // buckets（桶）的对数log_2
    // 如果B=5，则buckets数组的长度 = 2^5=32，意味着有32个桶
    B         uint8  
     // 溢出桶的大概数量
    noverflow uint16 
     // 哈希种子

    hash0     uint32 
    // 指向buckets数组的指针，数组大小为2^B，如果元素个数为0，它为nil。
    buckets    unsafe.Pointer 
 	// 如果发生扩容，oldbuckets是指向老的buckets数组的指针，老的buckets数组大小是新的buckets的1/2;非扩容状态下，它为nil。
    oldbuckets unsafe.Pointer 
    // 表示扩容进度，小于此地址的buckets代表已搬迁完成。
    nevacuate  uintptr        
    // 这个字段是为了优化GC扫描而设计的。当key和value均不包含指针，并且都可以inline时使用。extra是指向mapextra类型的指针。
    extra *mapextra 
 }
```



#### bmap结构体

bmap就是桶。每个桶里面会最多装8个key，这些 key 之所以会落入同一个桶，是因为它们经过哈希计算后，**哈希结果是“一类”的。**

在桶内，又会根据 key 计算出来的 hash 值的高 8 位来决定 key 到底落入桶内的哪个位置（一个桶内最多有8个位置)

```go
// A bucket for a Go map.
type bmap struct {
    tophash [bucketCnt]uint8        
    // len为8的数组
    // 用来快速定位key是否在这个bmap中
    // 桶的槽位数组，一个桶最多8个槽位，如果key所在的槽位在tophash中，则代表该key在这个桶中
}
//底层定义的常量 
const (
    bucketCntBits = 3
    bucketCnt     = 1 << bucketCntBits
    // 一个桶最多8个位置
）

但这只是表面(src/runtime/hashmap.go)的结构，编译期间会给它加料，动态地创建一个新的结构：

type bmap struct {
  topbits  [8]uint8
  keys     [8]keytype
  values   [8]valuetype
  pad      uintptr
  overflow uintptr
  // 溢出桶
}
```



从下图中可以看出key和value是各自放在一起的，好处是在某些情况下省略掉padding字段，节省内存空间

<img src="D:\workspace\Go\src\Golang-Master\Go_Base\img\go bucket.png" alt="go bucket" style="zoom: 67%;" />

### 主要特性

#### 引用特性

map是个指针，其底层为hmap，所以map是引用类型

#### 随机性

map是无序的，如果想顺序遍历map，需要对mapdekey先排序，再按照key的顺序遍历map

#### 共享存储空间

与slice相同

#### 非线程安全

Golang中的map是非线程安全的

##### map+sync.RWMutex

```go
func BenchmarkMapConcurrencySafeByMutex(b *testing.B) {
 var lock sync.Mutex //互斥锁
 m := make(map[int]int, 0)
 var wg sync.WaitGroup
 for i := 0; i < b.N; i++ {
  wg.Add(1)
  go func(i int) {
   defer wg.Done()
   lock.Lock()
   defer lock.Unlock()
   m[i] = i
  }(i)
 }
 wg.Wait()
 b.Log(len(m), b.N)
}
```



##### sync.Map

sync.map是用**读写分离实现**的，其思想是空间换时间（一个read map，一个write map)。和map+RWLock的实现方式相比，它做了一些优化：

可以**无锁访问read map**，而且**会优先操作read map**，倘若**只操作read map就可以满足要求(增删改查遍历)，那就不用去操作write map(它的读写都要加锁)**，所以在某些特定场景中它发生锁竞争的频率会远远小于map+RWLock的实现方式。



#### 哈希冲突

golang中的map底层使用hash table，用链表来解决冲突，出现冲突时，不是每一个key都申请一个结构通过链表串起来，而是以bmap为最小粒度挂载，一个bmap可以放8个kv。在哈希函数的选择上，会在程序启动时，检测 cpu 是否支持 aes，如果支持，则使用 aes hash，否则使用 memhash。



### Map操作

#### 创建

```go
// 1.声明变量
var m map[int]string
// 2.使用make
var m2 := make(map[int]string)
```

#### 增加

```go
var m map[string]string
m["France"]="Paris"
m["Italy"]="Rome"
```

#### 删除

```go
m := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}
delete(m,"France");
```



## Golang string

Golang中的string是由多个字符组成，其不可变，采用UTF-8编码

可以使用string标准库的方法来对字符串进行操作

```go
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

```



## Golang 函数

### 函数本质

函数也是一种数据类型，可以作为另一个函数的参数，也可以作为另一个函数的返回值。

```go
func add(a int, b int) int{
	return a+b
}
func sub(a int, b int) int{
	return a-b
}

func cal(op string) func(int,int) int{
	switch op{
	case "+":
		return add
	case "-":
		return sub
    default:
        return nil
	}
}

func main(){
    ff:=cal("+")
    r:=ff(1,2)
}
```

### 闭包

可以理解为定义在**一个函数内部的函数**。在本质上，闭包是将函数内部和函数外部连接起来的桥梁。

```go
// 返回一个函数
func add() func(int) int{
	var x int
	return func(y int) int{
		x += y
		return x
	}
}
func main(){
    // 变量f是一个函数，它引用了其外部作用域中的x变量，此时f就是一个闭包。
    // 在f的生命周期内，变量x也是一直有效的
    var f = add()
    fmt.Println(f(10))  // 10
    fmt.Println(f(20))  // 30
    fmt.Println(f(30))  // 60
}
```



### defer函数

可以在函数中添加多个defer语句。当函数执行到最后时，这些defer语句会**按照逆序执行**，最后该函数返回

```go
func ReadWrite() bool {
    file.Open("file")
    defer file.Close() // 最后才执行file.Close()
    if failureX {
          return false
    } i
    f failureY {
          return false
    } 
    return true
}
```

#### defer用途

- 关闭文件句柄
- 锁资源释放
- 数据库连接释放



### init函数

#### 主要特点

- init函数先于main函数自动执行，不能被其他函数调用
- init函数没有输入参数、返回值
- 每个包可以有多个init函数
- 包的每个源文件也可以有多个init函数

#### 作用

实现包级别的一些初始化操作



#### Golang中的执行顺序

**initVar->init->main**

```go
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

```

**输出**

```bash
PS D:\workspace\Go\src\Golang-Master\Go_Base\code> go run .\test_init.go
initVar...
init...
main
```



## Golang 指针



## Golang 结构体
