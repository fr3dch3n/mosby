.PHONY: build test clean

VERSION=$(shell git describe --tags)

build: mosby_amd64 mosby_alpine mosby_darwin mosby.exe

clean:
	rm -f *_amd64 *_darwin *.exe

test:
	go test -v ./...

mosby_amd64:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $@ *.go

mosby_alpine:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $@ *.go

mosby_darwin:
	GOOS=darwin go build -o $@ *.go

mosby.exe:
	GOOS=windows GOARCH=amd64 go build -o $@ *.go
