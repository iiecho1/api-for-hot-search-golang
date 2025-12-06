package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Sougou() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	url := "https://www.sogou.com/web?query=%E6%90%9C%E7%8B%97%E7%83%AD%E6%90%9C"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest error: %w", err)
	}
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.Client.Do error: %w", err)
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

	pattern := `<span [^>]*>[\s\S]*?<p>\s*<a href="([^"]+)"[^>]*>(.*?)</a>\s*</p>[\s\S]*?</span>\s*<span class="hot-rank-right">(.*?)</span>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)
	// 检查是否匹配到数据
	if len(matched) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "未匹配到数据，可能页面结构已变更",
			"icon":    "https://www.sogou.com/favicon.ico",
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
		"message": "搜狗",
		"icon":    "https://www.sogou.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api, nil
}
