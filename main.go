/**
 * @Author: Ne-21
 * @Description: main
 * @File: main.go
 * @Version: 1.3
 * @Date: 2022/3/5
 */

package main

import (
	"context"
	"github.com/Olixn/Potal-Auto-Auth/cmd"
	"github.com/Olixn/Potal-Auto-Auth/config"
	"github.com/Olixn/Potal-Auto-Auth/logger"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	logger.InitLog("/tmp")
	config.InitConfig(os.Args[1])
	cmd.InitCmd()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go cmd.Run(ctx)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
