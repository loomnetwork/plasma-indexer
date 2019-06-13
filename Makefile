PKG = github.com/loomnetwork/plasma-indexer

all: build

build: plasma-api loomstore-indexer

plasma-api:
	go build -o plasma-api $(PKG)

loomstore-indexer:
	go build -o loomstore-indexer $(PKG)/indexer/loomstore

test:
	go test -tags evm

clean:
	go clean
	rm -rf loomstore-indexer
	rm -rf plasma-api

loomstore-abigen: abigen
	abigen --abi abi/loom_store.abi --pkg ethcontract --type LoomStore --out ethcontract/loom_store.go

abigen:
	go get github.com/ethereum/go-ethereum/cmd/abigen

deps:
	dep ensure -vendor-only

.PHONY: all build clean test abigen loomstore-abigen
