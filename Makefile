PKG = github.com/loomnetwork/plasma-indexer

all: build

build:
	go build -o cardfaucet-indexer $(PKG)

test:
	go test -tags evm

clean:
	go clean
	rm -rf cardfaucet-indexer

cardfaucet-abigen: abigen
	abigen --abi abi/card_faucet.abi --pkg ethcontract --type CardFaucet --out ethcontract/card_faucet.go

abigen:
	go get github.com/ethereum/go-ethereum/cmd/abigen

.PHONY: all build clean test abigen cardfaucet-abigen