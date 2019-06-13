package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/getsentry/raven-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/loomnetwork/plasma-indexer/models"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmdExample = `
# Run with default configuration
loomstore-indexer
`

var rootCmd = &cobra.Command{
	Use:          "loomstore-indexer",
	Short:        "LoomStore indexer",
	Long:         `A DappChain indexer that captures events from LoomStore contract`,
	Example:      rootCmdExample,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().String("db-url", "", "MySQL Connection URL")
	rootCmd.PersistentFlags().String("db-host", "127.0.0.1", "MySQL host")
	rootCmd.PersistentFlags().String("db-port", "3306", "MySQL port")
	rootCmd.PersistentFlags().String("db-name", "plasma-indexer", "MySQL database name")
	rootCmd.PersistentFlags().String("db-user", "root", "MySQL database user")
	rootCmd.PersistentFlags().String("db-password", "", "MySQL database password")
	rootCmd.PersistentFlags().String("block-height", "", "Block height at which the scanner start query events from")
	rootCmd.PersistentFlags().String("chain-id", "default", "DAppChain Id")
	rootCmd.PersistentFlags().String("read-uri", "http://localhost:46658/query", "URI for quering app state")
	rootCmd.PersistentFlags().String("name", "cardfaucet", "Indexer name")
	rootCmd.PersistentFlags().Int("poll-interval", 5, "Poll interval in seconds")
	rootCmd.PersistentFlags().Int("block-interval", 20, "Amount of blocks to fetch")
	rootCmd.PersistentFlags().Int("reconnect-interval", 5, "Reconnect interval in seconds")
	rootCmd.PersistentFlags().String("contract-address", "", "Contract Address in hex format")

	viper.BindPFlag("db-url", rootCmd.PersistentFlags().Lookup("db-url"))
	viper.BindPFlag("db-host", rootCmd.PersistentFlags().Lookup("db-host"))
	viper.BindPFlag("db-port", rootCmd.PersistentFlags().Lookup("db-port"))
	viper.BindPFlag("db-name", rootCmd.PersistentFlags().Lookup("db-name"))
	viper.BindPFlag("db-user", rootCmd.PersistentFlags().Lookup("db-user"))
	viper.BindPFlag("db-password", rootCmd.PersistentFlags().Lookup("db-password"))
	viper.BindPFlag("chain-id", rootCmd.PersistentFlags().Lookup("chain-id"))
	viper.BindPFlag("read-uri", rootCmd.PersistentFlags().Lookup("read-uri"))
	viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("name"))
	viper.BindPFlag("poll-interval", rootCmd.PersistentFlags().Lookup("poll-interval"))
	viper.BindPFlag("block-interval", rootCmd.PersistentFlags().Lookup("block-interval"))
	viper.BindPFlag("block-height", rootCmd.PersistentFlags().Lookup("block-height"))
	viper.BindPFlag("reconnect-interval", rootCmd.PersistentFlags().Lookup("reconnect-interval"))
	viper.BindPFlag("contract-address", rootCmd.PersistentFlags().Lookup("contract-address"))

}

func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	var (
		dbURL             = viper.GetString("db-url")
		dbHost            = viper.GetString("db-host")
		dbPort            = viper.GetString("db-port")
		dbName            = viper.GetString("db-name")
		dbUser            = viper.GetString("db-user")
		dbPassword        = viper.GetString("db-password")
		readURI           = viper.GetString("read-uri")
		reconnectInterval = viper.GetInt("reconnect-interval")
		pollInterval      = viper.GetInt("poll-interval")
		contractAddress   = viper.GetString("contract-address")
		chainID           = viper.GetString("chain-id")
		name              = viper.GetString("name")
		blockInterval     = viper.GetInt("block-interval")
		blockHeight       = viper.GetInt64("block-height")
	)

	var parsedURL *url.URL
	var err error

	dbConnStr := dbURL
	if dbURL == "" {
		dbConnStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	}
	log.Printf("connecting to database host %s", dbHost)

	db, err := gorm.Open("mysql", dbConnStr)
	if err != nil {
		return errors.Wrapf(err, "fail to connect to database")
	}
	log.Printf("connected to database host %s", dbHost)
	defer db.Close()

	err = db.AutoMigrate(
		&models.Height{},
		&models.NewValueSet{},
	).Error
	if err != nil {
		return err
	}

	parsedURL, err = url.Parse(readURI)
	if err != nil {
		return errors.Wrapf(err, "Error parsing url %s", readURI)
	}

	// trap signals
	doneC := make(chan struct{})
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigC)

	config := &Config{
		ReadURI:           parsedURL.String(),
		PollInterval:      time.Duration(int64(pollInterval)) * time.Second,
		ReconnectInterval: time.Duration(int64(reconnectInterval)) * time.Second,
		ContractAddress:   contractAddress,
		ChainID:           chainID,
		Name:              name,
		BlockInterval:     blockInterval,
		BlockHeight:       uint64(blockHeight),
	}
	r := NewScanner(db, config)

	// run the scanner forever
	go r.Start()

	// watch for kill signals
	go func() {
		select {
		case <-sigC:
			log.Println("stopping scanner...")
			r.Stop()
			close(doneC)
		}
	}()

	<-doneC

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		raven.CaptureErrorAndWait(err, map[string]string{})
		log.Println(err)
		os.Exit(1)
	}
}
