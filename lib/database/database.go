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

	// // 启用Logger，显示详细日志
	// db.LogMode(true)

	// // 禁用日志记录器，不显示任何日志
	// db.LogMode(false)

	// // 调试单个操作，显示此操作的详细日志
	// db.Debug().Where("name = ?", "jinzhu").First(&User{})

}

// GetDB : 取得 db
func GetDB() *gorm.DB {
	return db
}
