package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Xinjingbao() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://www.bjnews.com.cn/"
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
	pattern := `<h3>\s*<a class="link" href="([^"]+)"[^>]*>\s*<span[^>]*>\d*</span>\s*(.*?)</a>\s*</h3>[\s\S]*?</i>(.*?)</span>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)
	// 检查是否匹配到数据
	if len(matched) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "未匹配到数据，可能页面结构已变更",
			"icon":    "https://www.bjnews.com.cn/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	var obj []map[string]interface{}
	for index, item := range matched {
		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item[2],
			"url":      item[1],
			"hotValue": item[3],
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "新京报",
		"icon":    "https://www.bjnews.com.cn/favicon.ico", // 20 x 20
		"obj":     obj,
	}
	return api, nil
}
