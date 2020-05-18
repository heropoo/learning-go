# package

## 说明
- 每个能运行的 go 程序都必须包含 main 函数。此函数是程序的入口点。 main 函数应该在 main package 中。
- 属于 package 的源文件应放在自己的单独文件夹中。 Go 规定使用与 package 名字相同的名称命名此文件夹。

## 临时设置GOPATH
```sh
export GOPATH=$PWD
``` 

## 安装运行
```sh
go install hello
./bin/hello
```
