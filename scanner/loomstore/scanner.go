package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
	loom "github.com/loomnetwork/go-loom"
	"github.com/loomnetwork/go-loom/client"
	pclient "github.com/loomnetwork/plasma-indexer/client"
	"github.com/loomnetwork/plasma-indexer/ethcontract"
	"github.com/loomnetwork/plasma-indexer/models"
	"github.com/pkg/errors"
)

type Config struct {
	ReadURI           string
	WriteURI          string
	ReconnectInterval time.Duration
	PollInterval      time.Duration
	ContractAddress   string
	ChainID           string
	Name              string
	BlockInterval     int
	BlockHeight       uint64
}

type Scanner struct {
	cfg   *Config
	db    *gorm.DB
	stopC chan interface{}
}

func NewScanner(db *gorm.DB, config *Config) *Scanner {
	return &Scanner{
		db:    db,
		cfg:   config,
		stopC: make(chan interface{}),
	}
}

func (s *Scanner) Start() {
	if err := s.setBlockHeight(s.cfg.BlockHeight); err != nil {
		panic(err)
	}
	for {
		err := s.scan()
		if err == nil {
			break
		}
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		// just delay before connecting again
		time.Sleep(s.cfg.ReconnectInterval)
	}
}

func (s *Scanner) Stop() {
	close(s.stopC)
}

func (s *Scanner) scan() error {
	ticker := time.NewTicker(s.cfg.PollInterval)

	// create new client
	dappClient := client.NewDAppChainRPCClient(s.cfg.ChainID, s.cfg.WriteURI, s.cfg.ReadURI)
	backend := pclient.NewLoomchainFilter(dappClient)
	contract, err := ethcontract.NewLoomStoreFilterer(common.HexToAddress(s.cfg.ContractAddress), backend)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ticker.C:
			var height models.Height
			err := s.db.Where(&models.Height{Name: s.cfg.Name}).
				Attrs(models.Height{Name: s.cfg.Name}).
				FirstOrCreate(&height).Error

			if err != nil {
				return err
			}
			fromBlock := height.LastBlockHeight + 1
			toBlock := fromBlock + uint64(s.cfg.BlockInterval) - 1

			lastBlockHeight, err := dappClient.GetBlockHeight()
			if err != nil {
				return err
			}

			if toBlock > lastBlockHeight {
				continue
			}

			filterOpts := &bind.FilterOpts{
				Start: fromBlock,
				End:   &toBlock,
			}

			events, err := fetchNewValueSet(dappClient, contract, filterOpts)
			if err != nil {
				return err
			}

			fmt.Printf("fetching block %d to %d , latest block: %d, %d events found\n", fromBlock, toBlock, lastBlockHeight, len(events))

			tx := s.db.Begin()
			if err := batchProcessNewValueSet(tx, events); err != nil {
				tx.Rollback()
				return err
			}

			height.LastBlockHeight = toBlock
			if err = tx.Save(&height).Error; err != nil {
				tx.Rollback()
				return err
			}
			tx.Commit()
		case <-s.stopC:
			ticker.Stop()
			return nil
		}

	}
}

func batchProcessNewValueSet(db *gorm.DB, events []*models.NewValueSet) error {
	if len(events) == 0 {
		return nil
	}
	for _, event := range events {
		err := db.Where(models.NewValueSet{TxHash: event.TxHash}).
			Assign(event).
			FirstOrInit(&event).
			Error
		if err != nil {
			return err
		}
		if err := db.Save(&event).Error; err != nil {
			return err
		}
	}
	return nil
}

func fetchNewValueSet(
	dappClient *client.DAppChainRPCClient, filterer *ethcontract.LoomStoreFilterer, filterOpts *bind.FilterOpts,
) ([]*models.NewValueSet, error) {
	it, err := filterer.FilterNewValueSet(filterOpts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get logs for GeneratedCard")
	}
	var chainID = dappClient.GetChainID()
	events := make([]*models.NewValueSet, 0)
	for {
		ok := it.Next()
		if ok {
			ev := it.Event
			receipt, err := dappClient.GetEvmTxReceipt(ev.Raw.TxHash.Bytes())
			if err != nil {
				return nil, err
			}
			contractAddr := loom.Address{ChainID: chainID, Local: receipt.ContractAddress}.MarshalPB()
			events = append(events, &models.NewValueSet{
				BlockHeight: ev.Raw.BlockNumber,
				TxIdx:       ev.Raw.TxIndex,
				Value:       ev.Value.String(),
				TxHash:      ev.Raw.TxHash.String(),
				Name:        ev.Name,
				Contract:    contractAddr.String(),
			})
		} else {
			err = it.Error()
			if err != nil {
				return nil, errors.Wrap(err, "Failed to get event data for GeneratedCard")
			}
			it.Close()
			break
		}
	}
	return events, nil
}

func (s *Scanner) setBlockHeight(h uint64) error {
	var height models.Height
	err := s.db.Where(&models.Height{Name: s.cfg.Name}).
		Attrs(models.Height{Name: s.cfg.Name}).
		FirstOrCreate(&height).Error
	height.LastBlockHeight = h
	if err = s.db.Save(&height).Error; err != nil {
		return err
	}
	return nil
}
