/**
 * @Author: Ne-21
 * @Description: tools
 * @File: utils.go
 * @Version: 1.0.0
 * @Date: 2022/3/13
 */

package utils

import (
	"github.com/Olixn/Potal-Auto-Auth/logger"
	"io/ioutil"
	"net"
	"net/http"
	url2 "net/url"
	"strings"
)

func Mac(interName string) (mac string) {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {
		if inter.Name == interName {
			mac := inter.HardwareAddr
			return mac.String()
		} else {
			continue
		}
	}
	return ""
}

func Ip(interName string) (ip string) {
	// 获取本机的IP地址
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
					return ipv4.String()
				}
			}
		}
	}
	return ""
}

func PostData(url string, ip string, mac string, mobile string, password string) {
	logger.Trace.Println("ip : " + ip + "	mac : " + mac)
	logger.Trace.Println("account : " + mobile + "	password : " + password)

	urlValues := url2.Values{}
	urlValues.Add("mobile", mobile)
	urlValues.Add("mobile_english", "")
	urlValues.Add("password", password)
	urlValues.Add("password_english", "")
	urlValues.Add("auth_type", "account")
	urlValues.Add("enterprise_id", "51")
	urlValues.Add("enterprise_url", "HBHUAWEI")
	urlValues.Add("site_id", "4662")
	urlValues.Add("client_mac", mac)
	urlValues.Add("nas_ip", "221.192.23.190")
	urlValues.Add("wlanacname", "None")
	urlValues.Add("user_ip", ip)
	urlValues.Add("3rd_ip", "None")
	urlValues.Add("ap_mac", "None")
	urlValues.Add("vlan", "11-11-11-11-11-11")
	urlValues.Add("ssid", "None")
	urlValues.Add("vlan_id", "None")
	urlValues.Add("ip", "None")
	urlValues.Add("ac_ip", "None")
	urlValues.Add("from", "None")
	urlValues.Add("sn", "None")
	urlValues.Add("gw_id", "None")
	urlValues.Add("gw_address", "None")
	urlValues.Add("gw_port", "None")
	urlValues.Add("url", "None")
	urlValues.Add("language_tag", "0")
	reqBody := urlValues.Encode()

	refer := "http://61.240.137.242:8888/hw/HBHUAWEI/login?apmac=11-11-11-11-11-11&userip=" + ip + "&nasip=221.192.23.190&user-mac=" + mac

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(reqBody))
	if err != nil {
		logger.Warning.Println(err)
	}
	req.Header.Add("Proxy-Connection", "keep-alive")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 Edg/98.0.1108.62")
	req.Header.Add("Origin", "http://61.240.137.242:8888")
	req.Header.Add("Referer", refer)
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")

	resp, err := client.Do(req)
	if err != nil {
		logger.Warning.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Warning.Println(err)
	}

	logger.Info.Println(string(body))
	logger.Info.Println("Login success!!!!")
	defer resp.Body.Close()
}
