# 通过tag标签获取版本号
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "1.0.0")
# 动态信息注入 版本号
LDFLAGS := -ldflags "-X config.Version=${VERSION}"


.PHONY: build
build:
	GOARCH=amd64 GOOS=linux go build -ldflags="-w -s"  -o  test main.go
