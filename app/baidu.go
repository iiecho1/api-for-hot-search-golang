package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func Baidu() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://top.baidu.com/board?tab=realtime"
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	pattern := `<div\sclass="c-single-text-ellipsis">(.*?)</div?`
	// 注意：这里假设 utils.ExtractMatches 没有改变，如果它也返回 error，需要修改
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	var obj []map[string]interface{}
	for index, item := range matched {
		// 添加边界检查
		if len(item) >= 2 {
			title := strings.TrimSpace(item[1])
			obj = append(obj, map[string]interface{}{
				"index": index + 1,
				"title": title,
				"url":   "https://www.baidu.com/s?wd=" + title,
			})
		}
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "百度",
		"icon":    "https://www.baidu.com/favicon.ico", // 64 x 64
		"obj":     obj,
	}
	return api, nil
}
