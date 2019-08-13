package config

import (
	"log"

	"github.com/spf13/viper"
)

/**
 * 因為會被 UnmarshalKey，
 * 故 struct 內的的變量都要求首字母大寫以保持可被外部引用
 */

// ServerStruct : struct of Server
type ServerStruct struct {
	Port string
}

// DatabaseStruct : struct of Database
type DatabaseStruct struct {
	SSLMode  string
	Host     string
	DBName   string
	User     string
	Password string
}

// Conf : struct of Configs toml
type Conf struct {
	Server   ServerStruct
	Database DatabaseStruct
}

var config Conf

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	v := viper.New()
	v.SetConfigType("toml")
	v.SetConfigName(env)
	v.AddConfigPath("../configs/")
	v.AddConfigPath("configs/")
	err = v.ReadInConfig() // viper 讀取 configs 內的 toml 檔
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}

	err = v.Unmarshal(&config) // 將 viper 讀取到的 configs 解析到 config
	if err != nil {
		log.Fatal("unable to decode viper into struct")
	}
}

// GetConfig : 取得 config
func GetConfig() *Conf {
	return &config
}
