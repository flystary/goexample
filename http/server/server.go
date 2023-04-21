package main

import (
	"log"
	"net/http"
	"time"
)


func main() {

	// 创建路由器
	mux := http.NewServeMux()
	// 设置路由规则
	mux.HandleFunc("/hello", sayHello)

	// 创建服务器
	server := &http.Server{
		Addr: ":1210",
		WriteTimeout: 3 * time.Second,
		Handler: mux,
	}

	log.Println("starting httpserver at localhost:1210")
	log.Fatal(server.ListenAndServe())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.Write([]byte("hello, this is httpserver!"))
}
