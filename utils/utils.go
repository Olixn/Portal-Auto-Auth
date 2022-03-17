/**
 * @Author: Ne-21
 * @Description: tools
 * @File: utils.go
 * @Version: 1.0.0
 * @Date: 2022/3/13
 */

package utils

import (
	"errors"
	"github.com/Olixn/Potal-Auto-Auth/logger"
	"net"
	url2 "net/url"
	"reflect"
)

func Mac(interName string) (mac string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {
		if inter.Name == interName {
			mac := inter.HardwareAddr
			if mac.String() == "" {
				err = errors.New("wan MAC Address is empty")
				return "", err
			}
			return mac.String(), nil
		} else {
			continue
		}
	}
	return "", err
}

func Ip(interName string) (ip string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {
		if inter.Name == interName {
			addrs, _ := inter.Addrs()
			for _, v := range addrs {
				ipv4 := v.(*net.IPNet).IP.To4()
				if ipv4 != nil {
					return ipv4.String(), nil
				} else {
					err = errors.New("wan IP Address is empty")
					return "", err
				}
			}
		}
	}
	return "", err
}

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
