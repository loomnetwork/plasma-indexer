package client

import (
	"context"
	"encoding/json"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/client"
	loomcommon "github.com/loomnetwork/go-loom/common"
	"github.com/pkg/errors"
)

type BlockHeight string

type EthBlockFilter struct {
	Addresses []loom.LocalAddress
	Topics    [][]string
}

type EthFilter struct {
	EthBlockFilter
	FromBlock BlockHeight
	ToBlock   BlockHeight
}

// LoomchainBackend implments a bare bone ethcontract filterer
type LoomchainBackend struct {
	*client.DAppChainRPCClient
}

var ErrNotImplemented = errors.New("not implememented")

var _ bind.ContractFilterer = &LoomchainBackend{}

func NewLoomchainFilter(dappClient *client.DAppChainRPCClient) bind.ContractFilterer {
	return &LoomchainBackend{DAppChainRPCClient: dappClient}
}

func (l *LoomchainBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	var ethFilter EthFilter
	if query.FromBlock != nil {
		ethFilter.FromBlock = BlockHeight(query.FromBlock.String())
	}
	if query.ToBlock != nil {
		ethFilter.ToBlock = BlockHeight(query.ToBlock.String())
	}
	if len(query.Addresses) > 0 {
		addrs, err := commonAddressToLoomAddress(query.Addresses...)
		if err != nil {
			return nil, err
		}
		ethFilter.Addresses = addrs
	}
	if len(query.Topics) > 0 {
		ethFilter.Topics = hashToStrings(query.Topics)
	}
	filter, err := json.Marshal(&ethFilter)
	if err != nil {
		return nil, err
	}
	logs, err := l.GetEvmLogs(string(filter))
	if err != nil {
		return nil, err
	}
	var tlogs []types.Log
	for _, log := range logs.EthBlockLogs {
		tlogs = append(tlogs, types.Log{
			Address:     common.BytesToAddress(log.Address),
			Topics:      byteArrayToHashes(log.Topics...),
			Data:        log.Data,
			BlockNumber: uint64(log.BlockNumber),
			TxHash:      common.BytesToHash(log.TransactionHash),
			TxIndex:     uint(log.TransactionIndex),
			BlockHash:   common.BytesToHash(log.BlockHash),
			Index:       uint(log.LogIndex),
			Removed:     log.Removed,
		})
	}
	return tlogs, nil
}

func (l *LoomchainBackend) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, ErrNotImplemented
}

// Utils

func byteArrayToHashes(bs ...[]byte) []common.Hash {
	var hashes []common.Hash
	for _, b := range bs {
		hashes = append(hashes, common.BytesToHash(b))
	}
	return hashes
}

func hashToStrings(hss [][]common.Hash) [][]string {
	var strs [][]string
	for _, hs := range hss {
		var newhs []string
		for _, h := range hs {
			newhs = append(newhs, h.String())
		}
		if len(newhs) > 0 {
			strs = append(strs, newhs)
		}
	}
	return strs
}

func commonAddressToLoomAddress(ca ...common.Address) ([]loomcommon.LocalAddress, error) {
	var addrs []loomcommon.LocalAddress
	for _, a := range ca {
		addr, err := loom.LocalAddressFromHexString(a.Hex())
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, addr)
	}
	return addrs, nil
}
