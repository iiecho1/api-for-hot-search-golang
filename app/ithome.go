package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Ithome() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://m.ithome.com/rankm/"
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP请求失败,状态码: %d", resp.StatusCode)
	}

	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	pattern := `<a href="(https://m\.ithome\.com/html/\d+\.htm)"[^>]*>[\s\S]*?<p class="plc-title">([^<]+)</p>`
	matches := utils.ExtractMatches(string(pageBytes), pattern)

	// 检查是否匹配到数据
	if len(matches) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "未匹配到数据，可能页面结构已变更",
			"icon":    "https://www.ithome.com/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	// 确定要取多少条数据（最多12条）
	count := len(matches)
	if count > 12 {
		count = 12
	}

	var obj []map[string]interface{}
	for index := 0; index < count; index++ {
		item := matches[index]
		// 添加边界检查
		if len(item) >= 3 {
			obj = append(obj, map[string]interface{}{
				"index": index + 1,
				"title": item[2],
				"url":   item[1],
			})
		}
	}

	// 确保有有效数据
	if len(obj) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "处理后的数据为空",
			"icon":    "https://www.ithome.com/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "IT之家",
		"icon":    "https://www.ithome.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api, nil
}
