package main

import (
	"fmt"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql" // load gorm mysql driver

	"github.com/getsentry/raven-go"
	"github.com/gorilla/pat"
	"github.com/jinzhu/gorm"
	"github.com/loomnetwork/plasma-indexer/controllers"
	"github.com/loomnetwork/plasma-indexer/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	db *gorm.DB
)

func main() {
	if err := loadConfig(); err != nil {
		panic(err)
	}
	if err := loadDB(); err != nil {
		panic(err)
	}
	defer db.Close()
	serverPort := viper.GetString("SERVER_PORT")
	log.Info("Server running on port ", serverPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", serverPort), loadRouter()))
}

func loadRouter() *pat.Router {
	plasmaController := controllers.PlasmaController{DB: db}
	router := pat.New()
	router.Get("/loomstore_events", raven.RecoveryHandler(plasmaController.ListLoomStoreEvents))
	return router
}

func loadDB() error {
	// connect to db
	log.Info("Connecting to DB")
	var err error
	db, err = connectToDb()
	if err != nil {
		return err
	}

	// migrate schemas
	err = autoMigrate()
	if err != nil {
		log.WithError(err).Error("error migrating schemas")
		raven.CaptureErrorAndWait(err, map[string]string{})
		return err
	}
	return nil
}

// connectToDB connects to db and return a new db struct
func connectToDb() (*gorm.DB, error) {
	dbURL := viper.GetString("DATABASE_URL")
	if dbURL == "" {
		dbUserName := viper.GetString("DATABASE_USERNAME")
		dbName := viper.GetString("DATABASE_NAME")
		dbPass := viper.GetString("DATABASE_PASS")
		dbHost := viper.GetString("DATABASE_HOST")
		dbPort := viper.GetString("DATABASE_PORT")
		dbURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", dbUserName, dbPass, dbHost, dbPort, dbName)
	}
	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func autoMigrate() error {
	if viper.GetBool("AUTOMIGRATE") {
		// Migrate the schema
		log.Info("Auto migrating schemas")
		if err := db.AutoMigrate(&models.NewValueSet{}).Error; err != nil {
			return err
		}
	}
	return nil
}

func loadConfig() error {
	viper.SetConfigName("plasma")
	viper.AddConfigPath(".")
	viper.SetDefault("DATABASE_USERNAME", "root")
	viper.SetDefault("DATABASE_NAME", "plasma-indexer")
	viper.SetDefault("DATABASE_HOST", "localhost")
	viper.SetDefault("DATABASE_PORT", "3306")
	viper.SetDefault("SERVER_PORT", "3333")
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to load config file: %s", err)
	}
	return nil
}
