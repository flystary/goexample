package main

import (
	"test/PRC/v3_GRPC/message"
	"context"
	"crypto/md5"
	"encoding/hex"
	"google.golang.org/grpc"
	"net"
)

type EncryptionServiceImpl struct {
}

func (es *EncryptionServiceImpl) Encryption(ctx context.Context, request *message.EncryptionRequest) (*message.EncryptionResult, error) {
	mdtStr := ToMd5(request.GetStr())
	res := message.EncryptionResult{Result: mdtStr}
	return &res, nil
}

func ToMd5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
	server := grpc.NewServer()

	message.RegisterEncryptionServiceServer(server, new(EncryptionServiceImpl))

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}
	_ = server.Serve(lis)
}
