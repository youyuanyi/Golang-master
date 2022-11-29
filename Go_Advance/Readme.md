# Golang进阶

## Golang 并发

### 协程

协程是一种**用户态**的轻量级线程，又称为微线程。

**协程的调度完全由用户控制**

与传统的系统级线程和进程相比，**协程的最大优势在于其轻量级**，可以轻松创建上百万个而不会导致系统资源衰竭。

### 并发模型

#### 线程模型

在现代OS中，**线程是处理器调度和分配的基本单位，进程是作为资源拥有的基本单位**。**每个进程有独立的进程地址空间、代码、数据、各种系统资源等。线程是进程内部的一个执行单元，每个进程内至少有一个主线程，它在这个进程创建时就产生，无需用户创建**。用户可根据需要在同一进程中创建多个线程。

**无论编程语言使用何种并发模型，到了OS层面，一定是以线程的形态存在**。**OS根据资源访问权限的不同，体系架构可分为用户空间和内核空间**。内核空间主要操作访问CPU资源、I/O资源、内存资源等硬件资源，为上层应用程序提供最基本的基础资源，用户空间是上层应用程序的固定活动空间，**用户空间不可以直接访问系统资源，必须通过“系统调用”、“库函数”或“Shell脚本”来调用内核空间提供的资源。**

**Go并发编程模型在底层是由OS所提供的线程库支撑的**。

**线程的主要3个实现模型：**

- 用户级线程模型
- 内核级线程模型（线程与内核调度实体[KSE]之间的关系）
- 两级线程模型

##### 内核级线程模型

程序创建的**用户线程与内核调度实体(KSE)是1:1关系**。大部分编程语言的线程库(如linux的pthread，Java的java.lang.Thread，C++11的std::thread等等)都是**对操作系统的线程（内核级线程）的一层封装，创建出来的每个线程与一个不同的KSE静态关联**，因此**其调度完全由OS调度器来做**。

![kernel thread](https://github.com/youyuanyi/Golang-master/raw/master/Go_Advance/img/kernel%20thread.jpg)

**优点**

实现方式简单，直接借助OS提供的线程能力

不同用户线程之间一般也不会影响，当一个线程被阻塞后，允许另一个线程继续执行

**缺点**

每创建一个用户级线程都需要创建一个内核级线程与其对应，这样创建线程的开销比较大，会影响到应用程序的性能。

创建，销毁以及多个线程之间的上下文切换等操作都是直接由OS层面来操作，在需要使用大量线程的场景下对OS的性能影响很大。



##### 用户级线程模型

用户线程与KSE是M：1关系，一个进程中所有创建的线程都与同一个KSE在运行时动态关联。

![kernel thread](https://github.com/youyuanyi/Golang-master/raw/master/Go_Advance/img/user%20thread.jpg)

**优点：**

相比于内核级线程更轻量，线程的创建、销毁、调度等操作由用户层面的线程库来实现，减少系统开销

**缺点：**

所有线程基于一个KSE，这意味着只有一个处理器可以被利用，在多处理器环境下这是不能够被接受的。本质上，用户线程只解决了并发问题，但是没有解决并行问题。如果线程因为 I/O 操作陷入了内核态，内核态线程阻塞等待 I/O 数据，则所有的线程都将会被阻塞，用户空间也可以使用非阻塞而 I/O，但是不能避免性能及复杂度问题。



##### 两级线程模型

用户线程与KSE是M：N关系。当某个KSE由于其工作的线程的阻塞操作而被内核调度出CPU时，当前与其关联的其余用户线程可以与其他的KSE建立关联关系。Go语言中的并发就是使用的这种实现方式。Go为了实现该模型自己实现了一个运行时调度器来负责Go中的"线程"与KSE的动态关联。

![kernel thread](https://github.com/youyuanyi/Golang-master/raw/master/Go_Advance/img/2%20thread.jpg)





## Golang 协程

### 创建协程

```go
package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg:%v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go showMsg("aaa") // 创建了一个协程
	showMsg("bbb")
}

```

### Goroutine的底层实现

Goroutine底层使用协程(coroutine)实现并发，coroutine是一种运行在用户态的用户线程，类似于绿色线程(greenthread)，其具有以下特点：

- 用户空间，避免了内核态和用户态的切换导致的成本
- 可以由语言和框架进行调度
- 更小的栈空间允许创建大量的实例



## Golang G-P-M模型

Go语言的go routine机制实现了M:N的线程模型。go routine是协程的一种实现，golang内置的调度器可以让多核CPU中每个CPU执行一个协程



## Golang 通道

通道(channel)是golang routines通信的管道。Go语言中，要传递某个数据给另一个go routine，可以把这个数据封装成一个对象，然后把这个对象的指针传入某个channel中，另外一个go routine就可以从这个channel中读出这个指针，并处理其指向的内存对象。

### channel线程安全

Golang channel是Golang中的一个数据类型，它是线程安全的。

因为channel的底层实现维护了一个互斥锁mutex，不需要人为实现加锁操作，所以channel是线程安全的

```go
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	lock mutex // 互斥锁
}
```



### 通道声明

```go
var 通道名 chan 传输的数据类型
例如: var a chan int

通道名 := make(chan float32)
例如: b := make(chan float32)
```

### channel的数据类型

**channel是引用类型，在作为参数传递时，传递的是内存地址。**



### channel的使用语法

#### 发送和接收

```go
data := <- a // 从chan a中读取数据
a <- data // 向chan a写入data
```

#### 发送和接收默认是阻塞的

一个通道发送和接收数据，默认是阻塞的.

当一个数据被发送到通道时，在发送语句中被阻塞，直到另一个Goroutine从该通道读取数据

当从通道读取数据时，读取被阻塞，直到一个Goroutine将数据写入该通道

```go
package main

import "fmt"

func main() {
	var ch1 chan bool       //声明，没有创建
	fmt.Println(ch1)        //<nil>
	fmt.Printf("%T\n", ch1) //chan bool
	ch1 = make(chan bool)   //0xc0000a4000,是引用类型的数据
	fmt.Println(ch1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine中，i：", i)
		}
		// 循环结束后，向通道中写数据，表示要结束了。。
		ch1 <- true

		fmt.Println("结束。。")

	}()

	data := <-ch1 // 从ch1通道中读取数据，这句话是阻塞的，如果ch1中没有数据，则主routine会阻塞在这里
	fmt.Println("data-->", data)
	fmt.Println("main。。over。。。。")
}
```

**输出**

```bash
<nil>
子goroutine中，i： 1
子goroutine中，i： 2
子goroutine中，i： 3
子goroutine中，i： 4
子goroutine中，i： 5
子goroutine中，i： 6
子goroutine中，i： 7
子goroutine中，i： 8
子goroutine中，i： 9
结束。。
data--> true
main。。over。。。。
```



#### channel的遍历

##### 使用for死循环

```go
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch3 := make(chan string, 4)
	go sendData3(ch3)
	for {
		time.Sleep(1*time.Second)
		v, ok := <-ch3  // !ok表示ch3关闭了
		if !ok {
			fmt.Println("读完了，，", ok)
			break
		}
		fmt.Println("\t读取的数据是：", v)
	}

	fmt.Println("main...over...")
}

func sendData3(ch3 chan string) {
	for i := 0; i < 10; i++ {
		ch3 <- "数据" + strconv.Itoa(i)
		fmt.Println("子goroutine，写出第", i, "个数据")
	}
	close(ch3) //子goroutines一定要关闭chan,否则会死锁
}


```



##### 使用for range形式

```go
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
```

执行结果

```bash
读取数据: 0
读取数据: 1
读取数据: 2
读取数据: 3
读取数据: 4
读取数据: 5
读取数据: 6
读取数据: 7
读取数据: 8
读取数据: 9
main over..
```



### 缓冲channel

上述所提的channel都是非缓冲channel，发送和接收到一个未缓冲的channel是阻塞的

Golang中的channel可以带有一个缓冲区，一个go routine发送数据到channel时，只有在这个channel的缓冲区满了时才会被阻塞。类似的，从缓冲通道接收的信息只有在缓冲区为空时才会被阻塞。

```go
ch := make(chan type,capacity)
```



### 单向channel

#### 只写

```go
ch := make(chan <- int)
```



#### 只读

```go
ch := make(<- chan int)
```



### Golang WaitGroup实现同步

**sync.WaitGroup**

#### 结构体

```go
type WaitGroup struct {
    noCopy noCopy
    state1 [3]uint32
}
```

noCopy是golang源码中检测禁止拷贝的技术。如果程序中有 WaitGroup 的赋值行为，使用 `go vet` 检查程序时，就会发现有报错。但需要注意的是，noCopy 不会影响程序正常的编译和运行。

整个结构体可以**简化为**

```go
type WaitGroup struct {
    counter int32
    waiter  uint32
    sema    uint32
}
```

- couter：当前尚未完成的个数。Add(n)会导致counter+=n，Done()会导致couter--
- waiter：表示当前调用了WaitGroup.Wait的goroutine的个数
- sema：对应于 golang 中 runtime 内部的信号量的实现。WaitGroup 中会用到 sema 的两个相关函数，`runtime_Semacquire` 和 `runtime_Semrelease`。`runtime_Semacquire` 表示**增加一个信号量，并挂起 当前 goroutine**。`runtime_Semrelease` 表示**减少一个信号量，并唤醒 sema 上其中一个正在等待的 goroutine**。

**WaitGroup的整个调用过程可以简单地描述成**：

1. 当调用 `WaitGroup.Add(n)` 时，counter 将会自增: `counter += n`

2. 当调用 `WaitGroup.Wait()` 时，会将 `waiter++`。同时调用 `runtime_Semacquire(semap)`, **增加信号量，并挂起当前 goroutine。**

3. 当调用 `WaitGroup.Done()` 时，将会 `counter--`。如果自减后的 **counter 等于 0，说明 WaitGroup 的等待过程已经结束，则需要调用 runtime_Semrelease 释放信号量，唤醒正在 `WaitGroup.Wait` 的 goroutine。**



#### 作用

使得主线程一直阻塞等待直到所有相关的子goroutines都结束

#### API

##### Add()

添加计数

##### Done()

减掉计数

##### Wait()

阻塞直到计数为0



#### 如果goroutine中不使用waitgroup

```go
package main

import (
	"fmt"
)

func hello(i int) {
	fmt.Printf("i: %v\n", i)
}

func main() {
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	fmt.Println("main over..")
}

```

**执行结果**

```bash
PS D:\workspace\Go\src\Golang-Master\Go_Advance\code> go run .\test_wg.go
i: 1
main over..
PS D:\workspace\Go\src\Golang-Master\Go_Advance\code> go run .\test_wg.go
main over..
```

可能子goroutine还没结束，main routine就结束了



### Golang runtime包

Go编译器产生的是本地可执行代码，但这些**代码仍然运行在Go的runtime调度器中**。它负责**内存分配、垃圾回收、栈处理、goroutine、channel、slice、map、反射**等等。

runtime包里面定义了一些channel管理相关的API

#### runtime.Gosched()

当前线程让出CPU时间片以让其它线程运行,它不会挂起当前线程，因此当前线程未来会继续执行

当一个 `goroutine` 发生阻塞，`Go` 会自动地把与该 `goroutine` 处于同一系统线程的其他 `goroutine` 转移到另一个系统线程上去，以使这些 `goroutine` 不阻塞

```go
package main

import (
	"fmt"
	"runtime"
)

func show(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("msg: %v\n", msg)
	}
}

func main() {
	go show("aaa")
	for i := 0; i < 2; i++ {
		runtime.Gosched()  // 把cpu时间片让给其他channel
		fmt.Println("bbb")
	}
	fmt.Println("main over")
}

```

**执行结果**

```bash
PS D:\workspace\Go\src\Golang-Master\Go_Advance\code> go run .\test_runtime.go
msg: aaa
msg: aaa
bbb
bbb
main over
```





#### runtime.Goexit()

**退出当前channel**



#### runtime.NumCPU()

返回当前系统的CPU核数量



#### runtime.GOMAXPROCS()

设置最大的可同时使用的CPU核数



#### runtime.NumGoroutine()

返回正在执行核排队的总goroutine数



### Golang Mutext互斥锁实现

```go
package main

import (
	"fmt"
	"sync"
)

var i int = 100
var wg sync.WaitGroup
var m sync.Mutex

func add() {
	defer wg.Done()
	m.Lock()
	i += 1 //对临界资源进行操作
	m.Unlock()
	fmt.Printf("i++:%v\n", i)
}
func sub() {
	defer wg.Done()
	m.Lock()
	i -= 1
	m.Unlock()
	fmt.Printf("i--:%v\n", i)
}

func main() {
	for i := 0; i < 10; i++ {
        wg.Add(1)  //开启一个子goroutine就要Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	wg.Wait()   // 等待所有子goroutine结束
	fmt.Println(i)
}

```

**执行结果**

```bash
PS D:\workspace\Go\src\Golang-Master\Go_Advance\code> go run .\test_mutex.go
i--:100
i--:99
i++:100
i--:99
i++:100
i--:99
i++:100
i--:99
i++:100
i--:99
i++:100
i--:99
i++:100
i++:101
i++:101
i--:100
i++:101
i++:101
i--:100
i--:100
100
```



### Golang select

select是Golang中的一个控制结构，类似于switch语句，但是**select会随机执行一个可运行的case**，**如果没有case可运行，它将执行default（如果有，否则阻塞），直到有case可运行**。

```go
select {
    case communication clause  :
       statement(s);      
    case communication clause  :
       statement(s); 
    /* 你可以定义任意数量的 case */
    default : /* 可选 */
       statement(s);
}
```

**每个case必须是一个channel表达式**

```go
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
```

**执行结果**

```bash
PS D:\workspace\Go\src\Golang-Master\Go_Advance\code> go run .\test_chan_select.go
ch2中取数据: 200
PS D:\workspace\Go\src\Golang-Master\Go_Advance\code> go run .\test_chan_select.go
从ch1中取得数据: 100
```



### Golang Timer

Golang中的定时器，内部也是通过channel实现的

**Timer是一次性的时间触发**

#### 创建Timer

```
t:=time.NewTimer(d) 
t:=time.AfterFunc(d,f)
c:=time.After(d)
//d：定时时间
//f：触发动作
```

#### time.NewTimer()

创建一个新的计时器，该计时器在其内部的channel上至少持续d之后发送当前时间



#### time.Stop()

计时器停止



#### time.After()

**在等待持续时间之后，然后在返回的通道上发送当前时间**，相当于NewTimer(d).C



### Golang Ticker

**Timer只执行一次，Ticker可以周期的执行**

#### 创建ticker

```
ticker := time.NewTicker(d Duration) // d为间隔时间
```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second) //间隔时间为1s
	counter := 1
	for _ = range ticker.C {
		fmt.Println("ticker..")
		counter++
		if counter >= 5 {
			ticker.Stop()
			break
		}
	}
}
```

#### 用Ticker实现周期性地在两个goroutine间收发数据

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second) //间隔时间为1s
	chanInt := make(chan int)
	go func() {  // sendData需要用goroutine
		defer close(chanInt)
		for _ = range ticker.C {  // 每隔1s执行一次
			select {
			case chanInt <- 1:
				fmt.Println("send 1")
			case chanInt <- 2:
				fmt.Println("send 2")
			}
		}
	}()

	sum := 0
	for v := range chanInt {
		fmt.Println("收到: ", v)
		sum += v
		fmt.Println("sum: ", sum)
		if sum >= 10 {
			ticker.Stop()
			break
		}
	}

}

```

**执行结果**

```bash
收到:  2
sum:  2
send 2
send 2
收到:  2
sum:  4
send 2
收到:  2
sum:  6
send 1
收到:  1
sum:  7
send 2
收到:  2
sum:  9
send 1
收到:  1
sum:  10
```



### Golang 原子操作

sync.atomic提供的原子操作保证任一时刻只有一个goroutine对变量进行操作



#### 增减操作

- atomic.AddInt32 
- atomic.AddInt64
- atomic.AddUInt32 
- atomic.AddUInt64

```
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i int32 = 100
	atomic.AddInt32(&i, 1)
	fmt.Println("i:", i)

	atomic.AddInt32(&i, -1)
	fmt.Println("i:", i)

	var j int64 = 200
	atomic.AddInt64(&j, 1)
	fmt.Println("j:", j)

}

```



#### 读写操作

- atomic.LoadInt32
- atomic.LoadInt64
- atomic.StoreInt32
- atomic.StoreInt64

```go
func test_load() {
	var i int32 = 100
	j := atomic.LoadInt32(&i)
	fmt.Println("j: ", j)  // 100
	atomic.StoreInt32(&i, 200)
	fmt.Println("i: ", i)  // 200
}
func main() {
	test_load()
}
```



#### 比较修改

atomic.CompareAndSwapInt32

```go
func test_cas() {
	var i int32 = 100
	// 修改i之前比较i和100( oldValue)是否一样，一样则改为200
	b := atomic.CompareAndSwapInt32(&i, 100, 200)
	fmt.Println("b: ", b)   // b: True
	fmt.Println("i: ", i)   // i: 200
}
func main() {
	test_cas()
}
```





## Golang CSP模型

### CSP是什么

Communicating Sequential Process (CSP，通信顺序进程)，是一种并发编程模型，用于描述**两个独立的并发实体通过共享的channel进行通信**的并发模型。

在Golang中，只用到了CSP的Process和Channel（对应到Golang中的goroutine/channel）

## Golang 错误处理

Golang中没有`try...catch`这类语法，而是鼓励每个函数都返回一个`error`，通过判断`error != nil`来捕捉错误

### 错误类型

```go
type error interface{
	Error() string
}
```

实现这个接口的类型都可以作为一个错误使用

### 自定义错误

#### errors New()

```go
package errors

// New returns an error that formats as the given text.
func New(text string) error {
    return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```

##### 例子

```go
package main

import (  
    "errors"
    "fmt"
    "math"
)

func circleArea(radius float64) (float64, error) {  
    if radius < 0 {
        return 0, errors.New("Area calculation failed, radius is less than zero")
    }
    return math.Pi * radius * radius, nil
}

func main() {  
    radius := -20.0
    area, err := circleArea(radius)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Area of circle %0.2f", area)
}
```



#### fmt Errorf()

```go
func Errorf(format string, a ...interface{}) error{}
```

##### 例子

```go
package main

import (  
    "fmt"
    "math"
)

func circleArea(radius float64) (float64, error) {  
    if radius < 0 {
        return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
    }
    return math.Pi * radius * radius, nil
}

func main() {  
    radius := -20.0
    area, err := circleArea(radius)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Area of circle %0.2f", area)
}
```



### panic

Golang的内建函数，如果函数F中书写了panic语句，会终止其后要执行的代码；如果函数F内存在defer函数，则按照defer的逆序执行；函数F的调用者G，调用完函数F后不会继续执行，直到整个goroutine退出并报告错误

### recover

Golang的内建函数，用来控制goroutine的panicking行为，捕获panic，一般用在defer中，通过recover来终止一个goroutine的panicking过程，从而恢复正常代码的执行，且可以获取通过panic传递的error

简单来说，go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理



### 异常处理场景

1. 空指针引用
2. 下标越界
3. 除数为0
4. 不应该出现的分支
5. 输入不应该引起函数错误

对于异常，我们可以选择在一个合适的上游去recover，并打印堆栈信息，使得部署后的程序不会终止。



## Golang 反射



