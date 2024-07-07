package config

import (
	"fmt"
	"github.com/olivere/elastic/v7"
)

var (
	EsClient *elastic.Client
)

func InitEs() {
	addr := fmt.Sprintf("%s:%d", appConfig.EsConfig.Host, appConfig.EsConfig.Port)
	var err error
	EsClient, err = elastic.NewClient(elastic.SetURL(addr), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	fmt.Println("ok")
}
