# GoLang基础

## Go的源码文件

![Go源码文件分类](.\img\Go源码文件分类.png)

### 命令源码文件

用户创建的.go源码文件，命令源码文件被install后，GOPATH如果只有一个工作区，那么相应的可执行文件会被存放到当前工作区的bin文件夹下；如果有多个工作区，就会安装到GOBIN指向的目录

多个命令源码文件可以分开用go run命令运行起来，但是无法通过go build和go install

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

库源码文件是存在于某个代码包中的普通的源码文件。库源码文件被install后，相应的归档文件(.a文件)会被存放到当前工作取得pkg的平台相关目录下

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
