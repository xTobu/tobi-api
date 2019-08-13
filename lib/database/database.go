package database

import (
	"fmt"
	"tobi-api/lib/config"
	"tobi-api/lib/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Init : Initial Db connection
func Init(c config.DatabaseStruct) (err error) {
	connArgs := fmt.Sprintf("sslmode=%s host=%s dbname=%s user=%s password=%s", c.SSLMode, c.Host, c.DBName, c.User, c.Password)
	db, err = gorm.Open("postgres", connArgs)
	if err != nil {
		log.Error("database.Init", err)
	}
	// 全局使用單數表名
	db.SingularTable(true)
	return
}

// GetDB : 取得 db
func GetDB() *gorm.DB {
	return db
}
