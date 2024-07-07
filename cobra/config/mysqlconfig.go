package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"sync"
)

var (
	MysqlDB *gorm.DB

	once sync.Once
)

func InitMysql() {

	host := appConfig.MysqlConfig.Host
	port := appConfig.MysqlConfig.Port
	user := appConfig.MysqlConfig.User
	password := appConfig.MysqlConfig.Password
	dbname := appConfig.MysqlConfig.Db

	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
		var err error
		MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("no")
			return
		}
		fmt.Println("ok")
	})

}
