package cmd

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCheckUrlStatus(t *testing.T) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	res, err := client.Get(CheckUrl)
	if err == nil && res.StatusCode == 204 {
		// 网络连接正常
		fmt.Println(res.StatusCode)
	}

}
