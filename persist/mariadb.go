package persist

import (
	"github.com/jinzhu/gorm"

	"github.com/Eric-GreenComb/travelrely-server/config"
)

// InitDb Init Db
func InitDb() *gorm.DB {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)

	if config.ServerConfig.GormLogMode == "false" {
		db.LogMode(false)
	}

	if err != nil {
		panic(err)
	}

	return db
}

// ConnectDb connect Db
func ConnectDb() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)

	if config.ServerConfig.GormLogMode == "false" {
		db.LogMode(false)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

// InitDatabase Init Db
func InitDatabase() {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	defer db.Close()

	if config.ServerConfig.GormLogMode == "false" {
		db.LogMode(false)
	}

	if err != nil {
		panic(err)
	}

	return
}
