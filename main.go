package main

import (
	"tobi-api/lib/config"
	"tobi-api/lib/database"
	"tobi-api/lib/log"
	"tobi-api/lib/database/repositories"
	"tobi-api/server"
	"time"
)

func main() {
	// 初始化 config
	config.Init("development")
	conf := config.GetConfig()
	log.Info("wait")
	time.Sleep(time.Duration(10) * time.Second)
	// 初始化 database
	database.Init(conf.Database)
	repositories.Init()

	// 初始化 server
	server.Init(conf.Server)
}
