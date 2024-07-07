/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/cobra"
	"sync"
	"time"
	"work/cobra/config"
	"work/cobra/types"
)

// jobsCmd represents the jobs command
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: AddElasticsearch,
}

func AddElasticsearch(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup
	Str := make(chan struct{}, 5)
	offset := 0
	for {
		var goods []types.Goods
		config.MysqlDB.Table("goods").Limit(100).Offset(offset).Find(&goods)
		if len(goods) == 0 {
			break
		}
		wg.Add(1)
		Str <- struct{}{}
		go func(Allgoods []types.Goods) {
			defer wg.Done()
			defer func() {
				<-Str
			}()

			bulk := config.EsClient.Bulk()
			for _, v := range Allgoods {
				marshal, _ := json.Marshal(v)
				doc := elastic.NewBulkIndexRequest().Index("goods").Doc(string(marshal))
				bulk.Add(doc)
			}
			timeout, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancelFunc()
			_, err := bulk.Do(timeout)
			if err != nil {
				fmt.Println("添加失败")
			} else {
				fmt.Printf("添加成功%d条数据\n", len(Allgoods))
			}

		}(goods)
		offset += 5

	}
	wg.Wait()
	fmt.Println("添加成功")
}
func init() {
	rootCmd.AddCommand(jobsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jobsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jobsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
