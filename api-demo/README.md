在编译的时候把一些debug信息给去掉:
```
go build -ldflags "-s"
```
-s:omit symbol table and debug info(忽略符号表和debug信息)
