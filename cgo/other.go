package main

/*
#include "other.c"
*/
import "C"
import "fmt"

func other() {

	sum := C.add(1, 4)

	fmt.Println("hello world : ", sum)

	var p1 = C.struct_Point{3, 6}
	var p2 = C.struct_Point{6, 9}

	arean := C.arean(p1, p2)

	p3 := new(C.struct_Point)
	p3.x = 5
	p3.y = 10

	p4 := new(C.struct_Point)
	p4.x = 8
	p4.y = 16

	arean1 := C.arean_point(p3, p4)

	fmt.Println("arean: ", arean)
	fmt.Println("arean1: ", arean1)
}

func main() {
	other()
}