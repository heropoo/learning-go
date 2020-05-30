---
title: 搭建Go语言开发环境
date: 2020-05-30 12:46:21
---

## 下载安装
官方网站：[https://golang.org/](https://golang.org/)

下载地址：[https://golang.org/dl/](https://golang.org/dl/)

当前Go语言的最新稳定版是v1.14.3，我们可以直接从[官方网站](https://golang.org/dl/)直接下载对应自己电脑系统的安装包安装。

安装之后打开终端执行：
```sh
go version
go version go1.14.3 windows/amd64
```
显示出你的Go版本信息以及操作系统信息表示安装成功

## Hello World
现在我们来创建第一个Go项目，找个合适的地方创建`hello`目录。

在这个目录中创建一个`main.o`文件，写入以下代码：
```go
package main    // 声明 main 包，表明当前是一个可执行程序

import "fmt"    // 导入标准库fmt包   

func main() {   // main函数，是程序执行的入口
    fmt.Println("Hello World!")     // 输出 Hello World!
}
```

我们执行`go build`将源代码编译成可执行文件：
```sh
go build
```

如果你是windows系统，将会生成一个`hello.exe`的可执行文件，我们执行它：
```sh
hello.exe
Hello World!
```

如果你是MacOS或者Linux系统，将会生成一个`hello`的可执行文件，我们执行它：
```sh
./hello
Hello World!
```

输出结果是 `Hello World!` ，那么就是执行成功了。 

哈哈~欢迎入坑！