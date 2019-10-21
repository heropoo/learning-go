package main

import "fmt"

func main() {
	//byteSlice := []byte{0x43, 0x61, 0x66, 0xc3, 0xa9}
	byteSlice := []byte{67, 97, 102, 195, 169}
	str := string(byteSlice)
	fmt.Println(str)
}
