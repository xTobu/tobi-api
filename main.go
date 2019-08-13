package main

import (
	"tobi-api/lib/config"
	"tobi-api/lib/database"
	"tobi-api/lib/utils"

	"time"
	"tobi-api/lib/database/repositories"
	"tobi-api/server"
)

var (
	// ENV 編譯的環境變數
	ENV = "development"
)

func main() {
	// 初始化 config
	config.Init(ENV)
	conf := config.GetConfig()
	// 初始化 database
	err := utils.Retry(3, 5*time.Second, func() (err error) {
		err = database.Init(conf.Database)
		return
	})
	if err != nil {
		panic("database.Init Fail")
	}
	repositories.Init()

	// 初始化 server
	server.Init(conf.Server)
}
