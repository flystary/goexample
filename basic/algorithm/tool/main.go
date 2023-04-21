package main

import "fmt"

func main() {
	//fmt.Println(reserve(123))  //整数反转 例子： 123 > 321
	//
	//
	//fmt.Println(toLowerCase("SUM")) //大写转小写
	//fmt.Println(toLowerCase("SuM")) //
	//fmt.Println(toLowerCase("MMMMMMMMMMMMMMMMMmmM")) //

	arr := []int{1, 3, 2, 8, 4, 7}
	result := mergeSort(arr)
	fmt.Println(result)
}
