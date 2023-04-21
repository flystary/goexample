package main

import (
	"crypto/md5"
	"encoding/hex"
	"net"
	"net/http"
	"net/rpc"
)

type EncryptionUtil struct {
}

func (eu *EncryptionUtil) Encryption(req string, resp *string) error {
	*resp = ToMd5(req)
	return nil
}

func ToMd5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
	encryption := new(EncryptionUtil)
	err := rpc.Register(encryption)
	if err != nil {
		panic(err.Error())
	}
	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	_ = http.Serve(listen, nil)
}
