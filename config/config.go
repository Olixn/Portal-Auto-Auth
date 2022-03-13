/**
 * @Author: Ne-21
 * @Description: 读取配置
 * @File: config.go
 * @Version: 1.0.0
 * @Date: 2022/3/13
 */

package config

import (
	"fmt"
	"github.com/Olixn/Potal-Auto-Auth/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var AppConfig *Config

type Config struct {
	Setting *Setting
}

type Setting struct {
	Mobile    string `yaml:"mobile"`
	Password  string `yaml:"password"`
	InterName string `yaml:"interName"`
}

func InitConfig() {
	AppConfig = &Config{}
	content, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		logger.Error.Println("解析config.yaml读取错误: " + err.Error())
		os.Exit(0)
	}

	fmt.Println(string(content))
	if yaml.Unmarshal(content, &AppConfig) != nil {
		logger.Error.Println("解析config.yaml读取错误: " + err.Error())
		os.Exit(0)
	}
}
