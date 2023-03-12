/**
 * @Author: Ne-21
 * @Description: 控制器
 * @File: cmd.go
 * @Version: 1.3
 * @Date: 2022/3/17
 */

package cmd

import (
	"context"
	"github.com/Olixn/Potal-Auto-Auth/config"
	"github.com/Olixn/Potal-Auto-Auth/logger"
	"github.com/Olixn/Potal-Auto-Auth/model"
	"github.com/Olixn/Potal-Auto-Auth/utils"
	"net/http"
	"os"
	"time"
)

var Portal *model.Portal
var Client *http.Client
var redirect bool
var portalUrl string

const CheckUrl = "http://connect.rom.miui.com/generate_204"
const CheckTime = 10

func InitCmd() {
	mobile := config.AppConfig.Setting.Mobile
	password := config.AppConfig.Setting.Password
	Portal = model.NewPortal()
	Portal.Mobile = mobile
	Portal.Password = password

	Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

func Login() {

	if !redirect {
		res, err := Client.Get("http://1.1.1.1")
		if err != nil {
			logger.Error.Println("There is an error in forcibly jumping to the authentication interface. Please restart the router or go to GitHub for help.")
			return
		}

		if redirectUrl := res.Header.Get("Location"); redirectUrl != "" {
			portalUrl = redirectUrl
		} else {
			logger.Error.Println("There is no redirect URL")
			return
		}
	}

	p := utils.ParseUrl(portalUrl)
	Portal.UserIp = p["userip"][0]
	Portal.NasIp = p["nasip"][0]
	Portal.ClientMac = p["user-mac"][0]

	values, err := utils.Struct2Values(Portal)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}
	Portal.Run(values)
}

func Check() (b bool) {

	res, err := Client.Get(CheckUrl)
	if err == nil && res.StatusCode == 204 {
		// 网络连接正常
		return true
	}

	if redirectUrl := res.Header.Get("Location"); redirectUrl != "" {
		redirect = true
		portalUrl = redirectUrl
		// 网络连接错误，但获取到了认证网址，执行登录
		return false
	} else {
		redirect = false
		// 网络连接错误，没有获取到了认证网址，执行跳转登录
		return false
	}

}

func Run(ctx context.Context) {
	for {
		_, err := os.Stat("/tmp/campus_run.log")
		if err != nil {
			if os.IsNotExist(err) {
				logger.InitLog("/tmp")
			}
		}
		if Check() {
			time.Sleep(time.Second * CheckTime)
		} else {
			Login()
			time.Sleep(time.Second * CheckTime * 2)
		}
		select {
		case <-ctx.Done():
			logger.Info.Println("程序退出啦~")
			os.Exit(0)
			return
		}
	}
}
