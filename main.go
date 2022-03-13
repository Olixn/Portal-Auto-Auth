/**
 * @Author: Ne-21
 * @Description: 配合crontab进行portal认证
 * @File: main.go
 * @Version: 1.0.0
 * @Date: 2022/3/5
 */

package main

import (
	"fmt"
	"github.com/Olixn/Potal-Auto-Auth/config"
	"github.com/Olixn/Potal-Auto-Auth/logger"
	"github.com/Olixn/Potal-Auto-Auth/utils"
	"os"
)

func init() {
	logger.InitLog("/tmp/tmp")
	config.InitConfig()
}

func main() {
	// http://61.240.137.242:8888/hw/HBHUAWEI/login?apmac=11-11-11-11-11-11&userip=10.255.202.150&nasip=221.192.23.190&user-mac=20:76:93:43:84:57
	logger.Info.Println("----------------------------------------------")
	logger.Info.Println("-----welcome to use Portal Auto Auth")
	logger.Info.Println("-----https://github.com/Olixn/Portal-Auto-Auth")
	logger.Info.Println("-----Author Ne-21 QQ 865194400")
	logger.Info.Println("----------------------------------------------")

	mobile := config.AppConfig.Setting.Mobile
	password := config.AppConfig.Setting.Password
	interName := config.AppConfig.Setting.InterName

	mac := utils.Mac(interName)
	fmt.Println("wan MAC : ", mac)
	if mac == "" {
		logger.Error.Println("wan mac is empty!")
		os.Exit(0)
	}
	ip := utils.Ip(interName)
	fmt.Println("wan IP : ", ip)
	if ip == "" {
		logger.Error.Println("wan ip is empty!")
		os.Exit(0)
	}
	utils.PostData("http://61.240.137.242:8888/hw/internal_auth", ip, mac, mobile, password)
	logger.Info.Println("----------------------------------------------")
	logger.Info.Println("-----Portal Auto Auth Finished")
	logger.Info.Println("----------------------------------------------")
	os.Exit(0)
}
