package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)


func main() {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:    30 * time.Second,
			KeepAlive:  30 * time.Second,
		}).DialContext,
		MaxIdleConns:     		100,
		IdleConnTimeout:  		90 * time.Second,
		TLSHandshakeTimeout:   	10 * time.Second,
		ExpectContinueTimeout: 	1 * time.Second,
	}

	client := &http.Client{
		Transport:  transport,
		Timeout:    30 * time.Second,
	}

	resp, err := client.Get("http://127.0.0.1:1210/hello")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bds, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bds))
}