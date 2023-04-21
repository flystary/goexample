package main

import (
	"fmt"
)

func binarySearch(s []int, number int) int {
	if len(s) <= 2 {
		if s[0] == number {
			return 0
		} else if s[1] == number {
			return 1
		} else {
			return -1
		}
	}
	low := 0
	hight := len(s) - 1
	for {
		if low > hight {
			return -2
		}

		mid := (low + hight) / 2

		if s[mid] > number {
			hight = mid - 1
		} else if s[mid] < number {
			low = mid + 1
		} else if s[mid] == number {
			return mid
		} else {
			return -3
		}
	}
}

func main() {
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("len() %d \n", len((s)))
	index := binarySearch(s, 9)
	fmt.Printf("====> %d \n", index)

	if index >= 0 {
		fmt.Printf("BinarySearch 2 =? %d\n", s[index])
	} else {
		fmt.Printf("没找到！\n")
	}
}
