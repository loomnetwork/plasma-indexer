PKG = github.com/loomnetwork/plasma-indexer
GO_ETHEREUM_DIR = $(GOPATH)/src/github.com/ethereum/go-ethereum

# loomnetwork/go-ethereum
ETHEREUM_GIT_REV = 1fb6138d017a4309105d91f187c126cf979c93f9

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

deps: $(GO_ETHEREUM_DIR)
	go get \
		github.com/loomnetwork/go-ethereum \
		github.com/grpc-ecosystem/go-grpc-prometheus \
		github.com/prometheus/client_golang/prometheus
	dep ensure -vendor-only

$(GO_ETHEREUM_DIR):
	git clone -q git@github.com:loomnetwork/go-ethereum.git 

.PHONY: all build clean test abigen loomstore-abigen
