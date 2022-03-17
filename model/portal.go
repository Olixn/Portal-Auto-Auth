/**
 * @Author: Ne-21
 * @Description:
 * @File: portal.go
 * @Version: 1.0.0
 * @Date: 2022/3/17
 */

package model

import (
	"github.com/Olixn/Potal-Auto-Auth/logger"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Portal struct {
	Mobile          string `param:"mobile"`
	MobileEnglish   string `param:"mobile_english"`
	Password        string `param:"password"`
	PasswordEnglish string `param:"password_english"`
	AuthType        string `param:"auth_type"`
	EnterpriseId    string `param:"enterprise_id"`
	EnterpriseUrl   string `param:"enterprise_url"`
	SiteId          string `param:"site_id"`
	ClientMac       string `param:"client_mac"`
	NasIp           string `param:"nas_ip"`
	Wlanacname      string `param:"wlanacname"`
	UserIp          string `param:"user_ip"`
	SrdIp           string `param:"3rd_ip"`
	ApMac           string `param:"ap_mac"`
	Vlan            string `param:"vlan"`
	Ssid            string `param:"ssid"`
	VlanId          string `param:"vlan_id"`
	Ip              string `param:"ip"`
	AcIp            string `param:"ac_ip"`
	From            string `param:"from"`
	Sn              string `param:"sn"`
	GwId            string `param:"gw_id"`
	GwAddress       string `param:"gw_address"`
	GwPort          string `param:"gw_port"`
	Url             string `param:"url"`
	LanguageTag     string `param:"language_tag"`
}

func (p *Portal) Run(values url.Values) {
	logger.Trace.Println("ip : " + p.UserIp + "	mac : " + p.ClientMac)
	logger.Trace.Println("account : " + p.Mobile + "	password : " + p.Password)

	reqBody := values.Encode()

	refer := "http://61.240.137.242:8888/hw/HBHUAWEI/login?apmac=11-11-11-11-11-11&userip=" + p.UserIp + "&nasip=221.192.23.190&user-mac=" + p.ClientMac

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://61.240.137.242:8888/hw/internal_auth", strings.NewReader(reqBody))
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

func NewPortal(mobile string, password string, ip string, mac string) (p *Portal) {
	return &Portal{
		Mobile:        mobile,
		Password:      password,
		UserIp:        ip,
		ClientMac:     mac,
		AuthType:      "account",
		EnterpriseId:  "51",
		EnterpriseUrl: "HBHUAWEI",
		SiteId:        "4662",
		NasIp:         "221.192.23.190",
		Vlan:          "11-11-11-11-11-11",
		LanguageTag:   "0",
	}
}
