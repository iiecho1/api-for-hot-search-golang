package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type sspResponse struct {
	Data []sspData `json:"data"`
}
type sspData struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
}

func Shaoshupai() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	url := "https://sspai.com/api/v1/article/tag/page/get?limit=100000&tag=%E7%83%AD%E9%97%A8%E6%96%87%E7%AB%A0"
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()
	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP请求失败，状态码: %d", resp.StatusCode)
	}
	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var resultMap sspResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}
	// 检查数据是否为空
	if len(resultMap.Data) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空",
			"icon":    "https://cdn-static.sspai.com/favicon/sspai.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	data := resultMap.Data
	var obj []map[string]interface{}
	for index, item := range data {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Title,
			"url":   "https://sspai.com/post/" + fmt.Sprint(item.ID),
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "少数派",
		"icon":    "https://cdn-static.sspai.com/favicon/sspai.ico", // 64 x 64
		"obj":     obj,
	}
	return api, nil
}
