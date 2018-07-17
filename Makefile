GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')

default: build

workdir:
	mkdir -p workdir

build: workdir/easyasync-cli

workdir/easyasync-cli: $(GOFILES)
	go build -o workdir/easyasync-cli .

dependencies: 
	@go get gopkg.in/urfave/cli.v1
test: test-all

test-all:
	@go test -v ./...

test-min:
	@go test ./...

release:
	echo "not implemented"