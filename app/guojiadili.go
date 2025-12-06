package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Guojiadili() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://www.dili360.com/"
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	pattern := `<li>\s*<span>\d*</span>\s*<h3><a href="(.*?)" target="_blank">(.*?)</a>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	// 检查是否匹配到数据
	if len(matched) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "未匹配到数据，可能页面结构已变更",
			"icon":    "https://www.dili360.com/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	var obj []map[string]interface{}
	for index, item := range matched {
		// 添加边界检查
		if len(item) >= 3 {
			urlPath := item[1]
			// 确保 URL 路径正确拼接
			fullURL := "http://www.dili360.com" + urlPath
			// 如果已经是完整 URL，不需要拼接
			if len(urlPath) > 4 && urlPath[:4] == "http" {
				fullURL = urlPath
			}

			obj = append(obj, map[string]interface{}{
				"index": index + 1,
				"title": item[2],
				"url":   fullURL,
			})
		}
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "国家地理",
		"icon":    "http://www.dili360.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api, nil
}
