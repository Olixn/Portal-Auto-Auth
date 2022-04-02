/**
 * @Author: Ne-21
 * @Description: 控制器
 * @File: cmd.go
 * @Version: 1.1
 * @Date: 2022/3/17
 */

package cmd

import (
	"github.com/Olixn/Potal-Auto-Auth/config"
	"github.com/Olixn/Potal-Auto-Auth/logger"
	"github.com/Olixn/Potal-Auto-Auth/model"
	"github.com/Olixn/Potal-Auto-Auth/utils"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var Portal *model.Portal

func InitCmd() {
	mobile := config.AppConfig.Setting.Mobile
	password := config.AppConfig.Setting.Password
	Portal = model.NewPortal()
	Portal.Mobile = mobile
	Portal.Password = password
}

func Login() {
	values, err := utils.Struct2Values(Portal)
	if err != nil {
		logger.Error.Println(err.Error())
		return
	}
	Portal.Run(values)
}

func Check() (b bool) {
	checkUrl := "http://connect.rom.miui.com/generate_204"

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	res, err := client.Get(checkUrl)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error.Println(err.Error())
		}
	}(res.Body)

	if err != nil {
		logger.Error.Println(err.Error() + ",Forced jump to authentication interface.")
	} else if res.StatusCode == 204 {
		logger.Info.Println("The network connection is normal.")
		return true
	} else {
		body, _ := ioutil.ReadAll(res.Body)
		logger.Trace.Println(string(body))
		logger.Error.Println("There is an error. Please restart the router or go to GitHub for help.")
		return true
	}

	res, err = client.Get("http://1.1.1.1")
	if err != nil {
		logger.Error.Println("There is an error in forcibly jumping to the authentication interface. Please restart the router or go to GitHub for help.")
		return true
	}

	if res.Header.Get("Location") != "" {
		logger.Info.Println("redirect portal url : " + res.Header.Get("Location"))
		p := utils.ParseUrl(res.Header.Get("Location"))
		Portal.UserIp = p["userip"][0]
		Portal.NasIp = p["nasip"][0]
		Portal.ClientMac = p["user-mac"][0]
		return false
	} else {
		logger.Info.Println("The network connection is normal")
		return true
	}
}

func Run() {
	for {
		_, err := os.Stat("/tmp/campus_run.log")
		if err != nil {
			if os.IsNotExist(err) {
				logger.InitLog("/tmp")
			}
		}
		if Check() {
			time.Sleep(time.Second * 60)
			continue
		} else {
			Login()
		}
		time.Sleep(time.Second * 60)
	}
}
