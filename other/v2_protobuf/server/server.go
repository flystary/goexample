package main

import (
	"test/PRC/v2_protobuf/message"
	"crypto/md5"
	"encoding/hex"
	"net"
	"net/http"
	"net/rpc"
)

type EncryptionService struct {
}

func (es *EncryptionService) Encryption(request *message.EncryptionRequest, response *message.EncryptionResult) error {
	md5Str := ToMd5(request.GetStr())
	response.Result = md5Str
	return nil
}

// ToMd5 简单封装的 md5
func ToMd5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
	// 功能对象注册
	encryption := new(EncryptionService)
	err := rpc.Register(encryption) //rpc.RegisterName("自定义服务名",encryption)
	if err != nil {
		panic(err.Error())
	}
	// HTTP注册
	rpc.HandleHTTP()

	// 端口监听
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	_ = http.Serve(listen, nil)
}
