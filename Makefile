BINARY := bin/wxipad-server
PKG    := ./...

.PHONY: help build run test vet fmt tidy clean

help:
	@echo "build  - 编译二进制到 $(BINARY)"
	@echo "run    - 本地运行服务"
	@echo "test   - 运行测试"
	@echo "vet    - 静态检查"
	@echo "fmt    - 格式化代码"
	@echo "tidy   - 整理依赖"
	@echo "clean  - 清理编译产物"

build:
	go build -trimpath -o $(BINARY) ./cmd/server

run:
	go run ./cmd/server -env dev

test:
	go test $(PKG)

vet:
	go vet $(PKG)

fmt:
	gofmt -w ./cmd ./internal

tidy:
	go mod tidy

clean:
	rm -rf bin
