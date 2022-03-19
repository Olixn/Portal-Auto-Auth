/**
 * @Author: Ne-21
 * @Description: 日志部分
 * @File: logger.go
 * @Version: 1.1
 * @Date: 2022/3/12
 */

package logger

import (
	"io"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func InitLog(path string) {
	file, err := os.OpenFile(path+"/campus_run.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error logger file:", err)
	}

	Trace = log.New(io.MultiWriter(file, os.Stderr),
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(io.MultiWriter(file, os.Stderr),
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(io.MultiWriter(file, os.Stderr),
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info.Println("----------------------------------------------")
	Info.Println("-----welcome to use Portal Auto Auth")
	Info.Println("-----https://github.com/Olixn/Portal-Auto-Auth")
	Info.Println("-----Author:Ne-21 QQ:865194400 Version:1.1")
	Info.Println("----------------------------------------------")
}
