package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func WangyiNews() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	url := "https://news.163.com/"
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP请求失败,状态码: %d", resp.StatusCode)
	}

	pageBytes, _ := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	pattern := `<em>\d*</em>\s*<a href="([^"]+)"[^>]*>(.*?)</a>\s*<span>(\d*)</span>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)
	// 检查是否匹配到数据
	if len(matched) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "未匹配到数据，可能页面结构已变更",
			"icon":    "https://news.163.com/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	var obj []map[string]interface{}
	for index, item := range matched {
		hot, err := strconv.ParseFloat(item[3], 64)
		if err != nil {
			// 处理转换错误，可以给默认值或者跳过该项
			hot = 0
			// 或者使用日志记录错误但不中断程序
			// log.Printf("parse hot value error for item %s: %v", title, err)
		}
		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item[2],
			"url":      item[1],
			"hotValue": fmt.Sprintf("%.1f万", hot/10000),
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "网易新闻",
		"icon":    "https://news.163.com/favicon.ico", // 16 x 16
		"obj":     obj,
	}
	return api, nil
}
