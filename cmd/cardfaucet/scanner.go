package cardfaucet

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
	"github.com/loomnetwork/plasma-indexer/model"
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
	contract, err := ethcontract.NewCardFaucetFilterer(common.HexToAddress(s.cfg.ContractAddress), backend)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ticker.C:
			var height model.Height
			err := s.db.Where(&model.Height{Name: s.cfg.Name}).
				Attrs(model.Height{Name: s.cfg.Name}).
				FirstOrCreate(&height).Error

			if err != nil {
				return err
			}
			fromBlock := height.LastBlockHeight + 1
			toBlock := fromBlock + uint64(s.cfg.BlockInterval) - 1

			lastBlockHeight, err := queryBlockHeight(dappClient)
			if err != nil {
				return err
			}

			if toBlock > lastBlockHeight {
				continue
			}

			fmt.Printf("fetching block %d to %d , latest block: %d\n", fromBlock, toBlock, lastBlockHeight)

			filterOpts := &bind.FilterOpts{
				Start: fromBlock,
				End:   &toBlock,
			}

			events, err := fetchGeneratedCard(dappClient, contract, filterOpts)
			if err != nil {
				return err
			}

			tx := s.db.Begin()
			if err := batchProcessCardGenerated(tx, events); err != nil {
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

func batchProcessCardGenerated(db *gorm.DB, events []*model.GeneratedCard) error {
	if len(events) == 0 {
		return nil
	}

	return nil
}

func fetchGeneratedCard(dappClient *client.DAppChainRPCClient, filterer *ethcontract.CardFaucetFilterer, filterOpts *bind.FilterOpts) ([]*model.GeneratedCard, error) {
	it, err := filterer.FilterGeneratedCard(filterOpts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get logs for GeneratedCard")
	}
	var chainID = dappClient.GetChainID()
	events := make([]*model.GeneratedCard, 0)
	for {
		ok := it.Next()
		if ok {
			ev := it.Event
			receipt, err := dappClient.GetEvmTxReceipt(ev.Raw.TxHash.Bytes())
			if err != nil {
				return nil, err
			}
			contractAddr := loom.Address{ChainID: chainID, Local: receipt.ContractAddress}.MarshalPB()
			events = append(events, &model.GeneratedCard{
				BlockHeight: ev.Raw.BlockNumber,
				TxIdx:       ev.Raw.TxIndex,
				Owner:       receipt.CallerAddress.String(),
				CardID:      ev.CardId.String(),
				BoosterType: ev.BoosterType,
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

func queryBlockHeight(c *client.DAppChainRPCClient) (uint64, error) {
	block, err := c.GetEvmBlockByNumber("latest", false)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get latest block number")
	}
	return uint64(block.Number), nil
}
