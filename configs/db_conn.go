package configs

import (
	"fmt"

	"github.com/gabriel-valin/mmr/pkg/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DbConn() (*gorm.DB, error) {
	conf, err := LoadConfig(".")
	if err != nil {
		return nil, err
	}
	port := conf.DBPort
	if port == "" {
		port = "3306"
	}
	dbName := conf.DBName
	if dbName == "" {
		dbName = "mmr_db"
	}
	// dsn := "user:pass@tcp(127.0.0.1:" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(sqlite.Open("gorm.sqlite"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Debug().AutoMigrate(&schemas.Player{}, &schemas.Perfomance{}, &schemas.Match{}, &schemas.RankedsStatus{})
	if err != nil {
		fmt.Print(err)
	}
	return db, nil
}
