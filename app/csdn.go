package app

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type csdbResponse struct {
	Data []csdnData `json:"data"`
}
type csdnData struct {
	Title    string `json:"articleTitle"`
	URL      string `json:"articleDetailUrl"`
	HotValue string `json:"pcHotRankScore"`
}

func CSDN() (map[string]interface{}, error) {
	// 创建自定义 Transport，跳过 TLS 验证（仅用于测试）
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 跳过证书验证
		},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	url := "https://blog.csdn.net/phoenix/web/blog/hotRank?&pageSize=100"
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP请求失败，状态码: %d", resp.StatusCode)
	}
	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var resultMap csdbResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, err
	}

	data := resultMap.Data
	var obj []map[string]interface{}
	for index, item := range data {
		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item.Title,
			"url":      item.URL,
			"hotValue": item.HotValue,
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "CSDN",
		"icon":    "https://csdnimg.cn/public/favicon.ico",
		"obj":     obj,
	}
	return api, nil
}
