---
title: Go语言基础-变量
date: 2020-05-30 12:30:30
---

## Go语言基础-变量

## 标识符
在编程语言中标识符就是程序员定义的具有特殊意义的词，比如变量名、常量名、函数名等等。 Go语言中标识符由 **字母** 、 **数字** 和 **`_`(下划线）** 组成，并且只能以**字母**和 **`_`** 开头。 举几个例子：`abc`，`_`，`_123`, `a123`。

## 关键字
关键字是指编程语言中预先定义好的具有特殊含义的标识符。 

Go语言中有25个关键字：
```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

此外，Go语言中还有37个保留字:
```go
Constants:  true  false  iota  nil

Types:      int  int8  int16  int32  int64  
            uint  uint8  uint16  uint32  uint64  uintptr
            float32  float64  complex128  complex64
            bool  byte  rune  string  error

Functions:  make  len  cap  new  append  copy  close  delete
            complex  real  imag
            panic  recover
```

关键字和保留字都不能用作变量名。

## 变量
变量是存储位置的名称，用于存储特定类型的值。在 go 中声明变量有多种语法。

我们使用 `var` 关键字申明一个变量，语法是：
```go
var 变量名 类型
```

我们写个例子 `var_01.go` ：
```go
package main

import "fmt"

func main() {
    var age int     // 定义一个变量age，类型为int类型
    fmt.Println("Var age is", age)
}
```

运行代码输出：
```sh
Var age is 0
```

上面代码`var age int` 声明一个名为 age 的 int 类型的变量，我们没有给变量分配任何值，但是运行之后输出 `0`。这是因为如果没有为变量赋值，则使用变量类型的零值自动初始化它。

