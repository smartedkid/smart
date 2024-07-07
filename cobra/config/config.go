package config

import "github.com/spf13/viper"

type MysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       string
}
type EsConfig struct {
	Host string
	Port int
}
type AppConfig struct {
	MysqlConfig
	EsConfig
}

var appConfig AppConfig

func InitConfig() {
	viper.SetConfigFile("config.yaml")
	viper.ReadInConfig()
	viper.Unmarshal(&appConfig)
}
