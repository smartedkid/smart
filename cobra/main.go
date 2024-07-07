/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"work/cobra/cmd"
	"work/cobra/config"
)

func main() {
	config.InitConfig()
	config.InitMysql()
	config.InitEs()
	cmd.Execute()
}
