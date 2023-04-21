package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//resp, err := http.Get("https://baidu.com/")
	//resp, err = http.Post("http://example.com/upload","image/jpeg",&buf)
	//rsp, err = http.PostForm("http://example.com/form", url.Values{"key":{"Value"},"id":{"123"}})

	resp, err := http.Get("https://www.liwenzhou.com")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Print(string(body))
}
