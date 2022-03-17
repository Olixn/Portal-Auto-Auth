/**
 * @Author: Ne-21
 * @Description: 配合crontab进行portal认证
 * @File: main.go
 * @Version: 1.0.0
 * @Date: 2022/3/5
 */

package main

import (
	"github.com/Olixn/Potal-Auto-Auth/cmd"
	"github.com/Olixn/Potal-Auto-Auth/config"
	"github.com/Olixn/Potal-Auto-Auth/logger"
	"os"
)

func init() {
	logger.InitLog("/tmp/tmp")
	config.InitConfig(os.Args[1])
}

func main() {
	logger.Info.Println("--------------------START---------------------")
	logger.Info.Println("-----welcome to use Portal Auto Auth")
	logger.Info.Println("-----https://github.com/Olixn/Portal-Auto-Auth")
	logger.Info.Println("-----Author Ne-21 QQ 865194400")
	logger.Info.Println("----------------------------------------------")
	cmd.Run()
	logger.Info.Println("---------------------END----------------------")
	os.Exit(0)
}
