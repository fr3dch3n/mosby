.PHONY: build test clean

VERSION=$(shell git describe --tags)

build: mosby_amd64 mosby_alpine mosby_darwin mosby.exe

clean:
	rm -f *_amd64 *_darwin *.exe

test:
	dep ensure
	go test -v ./...

mosby_amd64:
	dep ensure
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $@ *.go

mosby_alpine:
	dep ensure
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $@ *.go

mosby_darwin:
	dep ensure
	GOOS=darwin go build -o $@ *.go

mosby.exe:
	dep ensure
	GOOS=windows GOARCH=amd64 go build -o $@ *.go
