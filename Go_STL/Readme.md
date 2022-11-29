# Golang标准库

## OS模块

### 文件目录

#### 创建文件

```go
f,err := os.Create("file.txt")
if err != nil{
   
}
```

#### 创建目录

```go
err := os.Mkdir("test",os.ModePerm) // os.ModePerm是最高权限777
if err != nil{

}
err := os.MkdirAll("a/b/c",os.ModePerm)  // 创建级联目录
```

#### 删除目录\文件

```go
err := os.Remove("a.txt")
if err != nil{

}
err := os.RemoveAll("a") // a是目录
if err != nil{
    
}
```

#### 获得工作目录

```go
dir, err := os.Getwd()
```

#### 修改当前工作目录

```go
err := os.Chdir("../../a")
```

#### 获得临时目录

```go
s := os.TempDir()
```

#### 读文件

```go
b, err := os.ReadFile("test2.txt") // b是byte数组，要强转为string
if err != nil{
	...
}else{
	fmt.Printf("b: %v\n",string(b[:]))
}
```

#### 写文件

```go
s := "hello world"
os.WriteFile("test2.txt",[]byte(s),os.ModePerm)
```



### File文件读操作

```go
// 只读
f, _ := os.Open("a.txt")
for{
    buf := make([]byte,10)
    n,_ := f.Read(buf)
    fmt.Printf("%v\n",string(buf))
    if err == io.EOF{
        break
    }
}
f.Close()

// 根据第二个参数，可以读写或者创建
f2,_ := os.OpenFile("a1.txt",os.O_RDWR|os.O_CREATE,0755)
```



### File文件写操作

#### 写字节数组

```go
f,_ := os.OpenFile("a.txt",os.O_RDWR|os.O_APPEND,0775)
f.Write([]byte("hello"))
f.Close()
```



#### 写字符串

```go
f,_ := os.OpenFile("a.txt",os.O_RDWR|os.O_APPEND,0775)
f.WriteString("Hello World")
f.Close()
```



#### 随机写

```go
f,_ := os.OpenFile("a.txt",os.O_RDWR|os.O_APPEND,0775)
f.WriteAt([]byte("aaa"),3)
f.Close()
```



### 进程相关操作

#### 获得当前进程id

```go
os.Getpid()
```

#### 获得父id

```go
os.Getppid()
```

#### 开始一个新进程

```go
p, err := os.StartProcess("notepad.exe",[]string{"notepad.exe"."a.txt"},attr)
```

#### 向P进程发送退出信号

```go
p.Signal(os.Kill)
```

#### 等待进程p的退出

```go
ps, _ := p.Wait()
```



## ioutil包

### ReadAll

```go
f, _ := os.Open("a.txt")  // File实现了Reader接口
defer f.Close()
b,err := ioutil.ReadAll(f) // 可以读文件，也可以读其他的输入
if err != nil{
    log.Fatal(err)
}
fmt.Printf("%v\n",string(b))
```



### ReadDir

```go
dirs,_ := ioutil.ReadDir(".")
for , v := range(fi){
	fmt.Printf("v.Name():%v\n",v.Name())
}
```



### ReadFile

```go
b, _ := ioutil.ReadFile("a.txt")
fmt.Printf("string(b):%v\n",string(b))
```



#### WriteFile

```go
err := ioutil.WriteFile("a.txt",[]byte("hello"),0644)
if err != nil{
    log.Fatal(err)
}
```



## bufio

**实现了带有缓冲的I/O**

### bufio读

```go
f, err := os.Open("a.txt")
defer f.Close()
r2 := bufio.NewReader(f)
s, err := r2.ReadString('\n')
fmt.Printf("s:%v\n",s)
str := strings.NewReader("12345")

r2.Reset(str)
s, err := r2.ReadString('\n')
fmt.Println(s)
```



### bufio写

```go
f, err := os.OpenFile("a.txt",os.O_RDWR,0777)
defer f.Close()
w := bufio.NewWriter(f)
w.WriteString("hello world!")
w.Flush()
```

```
b := bytes.NewBuffer(make([]byte,0))
bw := bufio.NewWriter(b)
bw.WrteString("12345")
c := bytes.NewBuffer(make([]byte,0))
bw.Reset(c)
bw.WriteString("678")
bw.Flush()
```

```go
b := bytes.NewBuffer(make([]byte,0))
bw := bufio.NewWriter(b)
s := strings.NewReader("123")
br := bufio.NewReader(s)
rw := bufio.NewReadWriter(br,bw)
p, _ := rw.ReadString('\n')
fmt.Println(string(p)) // 123
rw.WriteString("asdf")
rw.Flush()
fmt.Println(b) //asdf
```



## 标准log库

### print

只打印日志

### panic

打印日志，并抛出panic异常

### fatal

打印日志，强制结束程序，**defer函数不会执行**

### 设置logFlags

```go
log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
```



## bytes包

bytes包提供了对字节切片进行读写的一系列函数

#### Contains

```
s := "abcdef"
b := []byte(s)
b1 := []byte("abcdef123")
b3 := bytes.Contains(b1,b)
```



### Replace

```go
s := []byte{"hello,world"}
old := []byte("o")
news := []byte("ee")
fmt.Println(string(bytes.Replace(s,old,news,0))) // hello,world
fmt.Println(string(bytes.Replace(s,old,news,1))) // hellee,world
fmt.Println(string(bytes.Replace(s,old,news,2))) // hellee,weerld
fmt.Println(string(bytes.Replace(s,old,news,-1))) // hellee,weerld
```



### Runes

```go
s := []byte("你好世界")
r := bytes.Runes(s)
fmt.Println(len(s)) // 12
fmt.Println(len(r)) // 4
```



### Buffer类型

#### 声明Buffer

```go
// 四种方法
var b bytes.Buffer
b := new(bytes.Buffer)
b := bytes.Buffer(s []byte)
b := bytes.BufferString(s string)
```



#### 往Buffer中写入数据

``` /
b.Write(d []byte) //将切片d写入Buffer尾部
b.WriteString(s string) //将字符串s写入Buffer尾部
b.WrtieByte(c byte)  // 将字符c写入Buffer尾部
b.WriteRune(r rune) //将一个rune类型的数据放到缓冲区的尾部
```



#### 从Buffer中读取数据到指定容器

```
c := make([]byte,10)
b.Read(c) //一次读取10个byte到c容器中
b.ReadByte() // 读取第一个byte
b.ReadRune() // 读取第一个Rune
b.ReadBytes(delimiter byte) // 需要一个byte作为分隔符，读的时候从缓冲区里找出第一个出现的分隔符
b.ReadString(delimiter byte) // 需要一个byte作为分隔符，读的时候从缓冲区里找出第一个出现的分隔符
```

