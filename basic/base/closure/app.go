package main

import (
	"fmt"
	"net/http"
	"time"
)

func myFunc() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Microsecond)
}

func coolFunc(a func()) {
	start := time.Now()
	a()
	end := time.Since(start)
	fmt.Printf("该函数执行完成耗时: %s\n", end)
}
func isAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Checking to see if Authorized header set ...")

		if val, ok := r.Header["Athorized"]; ok {
			fmt.Println(val)
			if val[0] == "true" {
				fmt.Println("Header is set! We can server content!")
				endpoint(w, r)
			}
		} else {
			fmt.Println("Not Authorized!!")
			fmt.Fprintf(w, "Not Authorized!!")
		}
	})
}
