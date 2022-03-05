/**
 * @Author: Ne-21
 * @Description: 配合crontab进行portal认证
 * @File: main.go
 * @Version: 1.0.0
 * @Date: 2022/3/5
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	url2 "net/url"
	"os"
	"strings"
)

func init() {
	// 初始化日志
	logFile, err := os.OpenFile(`/tmp/tmp/日志文件.log`, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// 设置存储位置
	log.SetOutput(logFile)
}

func PostData(url string, ip string, mac string) {
	refer := "http://61.240.137.242:8888/hw/HBHUAWEI/login?apmac=11-11-11-11-11-11&userip=" + ip + "&nasip=221.192.23.190&user-mac=" + mac

	urlValues := url2.Values{}
	urlValues.Add("mobile", "")
	urlValues.Add("mobile_english", "")
	urlValues.Add("password", "")
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
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(reqBody))
	req.Header.Add("Proxy-Connection", "keep-alive")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 Edg/98.0.1108.62")
	req.Header.Add("Origin", "http://61.240.137.242:8888")
	req.Header.Add("Referer", refer)
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	resp, err := client.Do(req)

	var result map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}
	log.Println(result)
	for k, v := range result {
		fmt.Println(k, ":", v)
	}
	defer resp.Body.Close()
}

func Mac() (mac string) {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {
		if inter.Name == "wan" {
			mac := inter.HardwareAddr
			return mac.String()
		} else {
			continue
		}
	}
	return ""
}

func Ip() (ip string) {
	// 获取本机的IP地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	for _, inter := range interfaces {
		if inter.Name == "wan" {
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

func main() {
	// http://61.240.137.242:8888/hw/HBHUAWEI/login?apmac=11-11-11-11-11-11&userip=10.255.202.150&nasip=221.192.23.190&user-mac=20:76:93:43:84:57
	mac := Mac()
	fmt.Println("wan MAC = ", mac)
	ip := Ip()
	fmt.Println("wan IP = ", ip)
	PostData("http://61.240.137.242:8888/hw/internal_auth", ip, mac)
}
