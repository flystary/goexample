package main

import "fmt"

func main() {
	mySlice := make([]string, 5, 8)
	mySlice[0] = "Apple"
	mySlice[1] = "Orange"
	mySlice[2] = "Banana"
	mySlice[3] = "Grape"
	mySlice[4] = "Plum"
	fmt.Println(mySlice)

	//mySlice := make([]string,5)
	//mySlice := []string{"Apple", "Orange", "Banana", "Grape", "Plum"}
	newSlice := mySlice[1:4]
	fmt.Println(newSlice)

}
