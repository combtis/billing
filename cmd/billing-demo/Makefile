DT:=$(shell date +%Y-%m-%dT%H:%M:%SZ%z)
PATH:=${PATH}:/usr/local/go/bin:/root/go/bin
GOROOT:=/usr/local/go

.PHONY:	all
all:
	@echo -n "${DT}\n"
	@cd ../..;go mod tidy;go fmt ./...
	@go run -trimpath -ldflags='-s -w' -gcflags='all=-N -l -c 2' .
