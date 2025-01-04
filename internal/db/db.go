package db

import (
	"backend/internal/config"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB(config *config.Config) {
	var gdb *gorm.DB
	var err error
	if config.Gorm.DBType == "mysql" {
		config.Gorm.DSN = config.MySQL.DSN()
	} else if config.Gorm.DBType == "sqlite3" {
		config.Gorm.DSN = config.Sqlite3.DSN()
	}
	gdb, err = gorm.Open(config.Gorm.DBType, config.Gorm.DSN)
	if err != nil {
		panic(err)
	}
	gdb.SingularTable(true)
	if config.Gorm.Debug {
		gdb.LogMode(true)
		gdb.SetLogger(log.New(os.Stdout, "\r\n", 0))
	}
	gdb.DB().SetMaxIdleConns(config.Gorm.MaxIdleConns)
	gdb.DB().SetMaxOpenConns(config.Gorm.MaxOpenConns)
	gdb.DB().SetConnMaxLifetime(time.Duration(config.Gorm.MaxLifetime) * time.Second)
	DB = gdb
}
