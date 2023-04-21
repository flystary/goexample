package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTimer(2 * time.Second)
	i := 0

	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				ticker.Stop()
			}
		}
	}()
	for {

	}
}
