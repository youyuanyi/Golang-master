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
