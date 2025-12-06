package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ppResponse struct {
	Data ppData `json:"data"`
}
type ppData struct {
	HotNews []news `json:"hotNews"`
}
type news struct {
	Title  string `json:"name"`
	ContId string `json:"contId"`
}

func Pengpai() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://cache.thepaper.cn/contentapi/wwwIndex/rightSidebar"
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
	var resultMap ppResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}
	// 检查数据是否为空
	if len(resultMap.Data.HotNews) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空",
			"icon":    "https://www.thepaper.cn/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	data := resultMap.Data.HotNews
	var obj []map[string]interface{}
	for index, item := range data {
		// 确保 ContId 不为空
		if item.ContId == "" {
			continue // 跳过没有 ContId 的新闻
		}

		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Title,
			"url":   "https://www.thepaper.cn/newsDetail_forward_" + item.ContId,
		})
	}

	// 如果所有数据都因为没有 ContId 被跳过
	if len(obj) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据格式异常，缺少必要字段",
			"icon":    "https://www.thepaper.cn/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "澎湃新闻",
		"icon":    "https://www.thepaper.cn/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api, nil
}
