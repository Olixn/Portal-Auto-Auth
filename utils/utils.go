/**
 * @Author: Ne-21
 * @Description: tools
 * @File: utils.go
 * @Version: 1.1
 * @Date: 2022/3/13
 */

package utils

import (
	"errors"
	"github.com/Olixn/Potal-Auto-Auth/logger"
	"net"
	url2 "net/url"
	"reflect"
	"time"
)

func Struct2Values(a interface{}) (values url2.Values, err error) {
	rType := reflect.TypeOf(a).Elem()
	rValue := reflect.ValueOf(a).Elem()
	if rType.Kind() != reflect.Struct {
		logger.Error.Println("Struct2Values Error! Not Struct")
		err = errors.New("Struct2Values Error! Not Struct")
		return nil, err
	}

	values = url2.Values{}
	for i := 0; i < rValue.NumField(); i++ {
		if rValue.Field(i).String() == "" {
			rValue.Field(i).SetString("none")
		}
		values.Add(rType.Field(i).Tag.Get("param"), rValue.Field(i).String())
	}

	return values, nil
}

func ParseUrl(url string) (urlParams url2.Values) {
	values, err := url2.ParseRequestURI(url)
	if err != nil {
		logger.Error.Println("url parse error!" + err.Error())
	}

	urlParams = values.Query()
	return
}

func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}
