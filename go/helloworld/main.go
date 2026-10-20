package main

import (
	"fmt"
	"unsafe"
)

func bool2Int(b bool) int {
	return int(*(*byte)(unsafe.Pointer(&b)))
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(bool2Int(5 >= 4))
}
