/**
 * @Author: Ne-21
 * @Description:
 * @File: portal.go
 * @Version: 1.1
 * @Date: 2022/3/17
 */

package model

import (
	"encoding/json"
	"github.com/Olixn/Potal-Auto-Auth/logger"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Portal struct {
	Mobile          string `json:"mobile"`
	MobileEnglish   string `json:"mobile_english"`
	Password        string `json:"password"`
	PasswordEnglish string `json:"password_english"`
	AuthType        string `json:"auth_type"`
	EnterpriseId    string `json:"enterprise_id"`
	EnterpriseUrl   string `json:"enterprise_url"`
	SiteId          string `json:"site_id"`
	ClientMac       string `json:"client_mac"`
	NasIp           string `json:"nas_ip"`
	Wlanacname      string `json:"wlanacname"`
	UserIp          string `json:"user_ip"`
	SrdIp           string `json:"3rd_ip"`
	ApMac           string `json:"ap_mac"`
	Vlan            string `json:"vlan"`
	Ssid            string `json:"ssid"`
	VlanId          string `json:"vlan_id"`
	Ip              string `json:"ip"`
	AcIp            string `json:"ac_ip"`
	From            string `json:"from"`
	Sn              string `json:"sn"`
	GwId            string `json:"gw_id"`
	GwAddress       string `json:"gw_address"`
	GwPort          string `json:"gw_port"`
	Url             string `json:"url"`
	LanguageTag     string `json:"language_tag"`
}

func (p *Portal) Run(values url.Values) {
	logger.Trace.Println("ip : " + p.UserIp + "	mac : " + p.ClientMac)
	logger.Trace.Println("account : " + p.Mobile + "	password : " + p.Password)

	reqBody := values.Encode()

	referer := "http://61.240.137.242:8888/hw/HBHUAWEI/login?apmac=" + p.Vlan + "&userip=" + p.UserIp + "&nasip=" + p.NasIp + "&user-mac=" + p.ClientMac

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
	req.Header.Add("Referer", referer)
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")

	resp, err := client.Do(req)
	if err != nil {
		logger.Warning.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Warning.Println(err)
	}

	var respond map[string]interface{}
	err = json.Unmarshal(body, &respond)
	if err != nil {
		logger.Warning.Println(err)
		return
	}

	if respond["op"].(string) == "ok" {
		logger.Info.Println("Login success!")
	} else {
		logger.Info.Println("Login error!" + respond["message"].(string))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error.Println(err.Error())
		}
	}(resp.Body)
}

func NewPortal() (p *Portal) {
	return &Portal{
		Vlan:          "11-11-11-11-11-11",
		AuthType:      "account",
		EnterpriseId:  "51",
		EnterpriseUrl: "HBHUAWEI",
		SiteId:        "4662",
		LanguageTag:   "0",
	}
}
