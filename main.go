/**
 * @Author: Ne-21
 * @Description: main
 * @File: main.go
 * @Version: 1.1
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
	cmd.InitCmd()
}

func main() {
	cmd.Run()
}
