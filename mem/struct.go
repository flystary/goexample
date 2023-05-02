package main

import (
	"fmt"
	"unsafe"
)

type C struct {
	a struct{}
	b int64
	c int64
}

type D struct {
	a int64
	b struct{}
	c int64
}

type E struct {
	a int64
	b int64
	c struct{}
}

type F struct {
	a int32
	b int32
	c struct{}
}

/*
 *如果空结构体作为结构体的内置字段：当变量位于结构体的前面和中间
 *时，不会占用内存；当该变量位于结构体的末尾位置时，需要进行内存
 *对齐，内存占用大小和前一个变量的大小保持一致。
*/
func main() {
	fmt.Println(unsafe.Sizeof(C{})) // 16 0+8+8
	fmt.Println(unsafe.Sizeof(D{})) // 16 8+0+8
	fmt.Println(unsafe.Sizeof(E{})) // 24 8+8+8
	fmt.Println(unsafe.Sizeof(F{})) // 12 4+4+4
}
