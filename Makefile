PKG = github.com/loomnetwork/plasma-indexer

all: build

build: plasma-indexer loomstore-scanner

plasma-indexer:
	go build -o plasma-indexer $(PKG)

loomstore-scanner:
	go build -o loomstore-scanner $(PKG)/scanner/loomstore

test:
	go test -tags evm

clean:
	go clean
	rm -rf loomstore-scanner
	rm -rf plasma-indexer

loomstore-abigen: abigen
	abigen --abi abi/loom_store.abi --pkg ethcontract --type LoomStore --out ethcontract/loom_store.go

abigen:
	go get github.com/ethereum/go-ethereum/cmd/abigen

deps:
	dep ensure -vendor-only

.PHONY: all build clean test abigen cardfaucet-abigen
