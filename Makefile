test: deps
	go test -race -v ./...

export IPFS_API ?= v04x.ipfs.io

deps:
	go get -t ./...
