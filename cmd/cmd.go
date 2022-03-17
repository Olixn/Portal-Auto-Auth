/**
 * @Author: Ne-21
 * @Description: 控制器
 * @File: cmd.go
 * @Version: 1.0.0
 * @Date: 2022/3/17
 */

package cmd

import (
	"fmt"
	"github.com/Olixn/Potal-Auto-Auth/config"
	"github.com/Olixn/Potal-Auto-Auth/logger"
	"github.com/Olixn/Potal-Auto-Auth/model"
	"github.com/Olixn/Potal-Auto-Auth/utils"
	"os"
)

func Run() {
	mobile := config.AppConfig.Setting.Mobile
	password := config.AppConfig.Setting.Password
	interName := config.AppConfig.Setting.InterName

	mac, err := utils.Mac(interName)
	fmt.Println("wan MAC : ", mac)
	if err != nil {
		logger.Error.Println("wan mac is empty!" + err.Error())
		os.Exit(0)
	}

	ip, err := utils.Ip(interName)
	fmt.Println("wan IP : ", ip)
	if err != nil {
		logger.Error.Println("wan ip is empty!" + err.Error())
		os.Exit(0)
	}

	portal := model.NewPortal(mobile, password, ip, mac)
	values, err := utils.Struct2Values(portal)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}
	portal.Run(values)
}
