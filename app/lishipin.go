package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Lishipin() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://www.pearvideo.com/popular"
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

	pattern := `<a\shref="(.*?)".*?>\s*<h2\sclass="popularem-title">(.*?)</h2>\s*<p\sclass="popularem-abs padshow">(.*?)</p>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)
	// 检查是否匹配到数据
	if len(matched) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "未匹配到数据，可能页面结构已变更",
			"icon":    "https://page.pearvideo.com/webres/img/logo.png",
			"obj":     []map[string]interface{}{},
		}, nil
	}
	var obj []map[string]interface{}
	for index, item := range matched {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item[2],
			"url":   "https://www.pearvideo.com/" + fmt.Sprint(item[1]),
			"desc":  item[3],
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "梨视频",
		"icon":    "https://page.pearvideo.com/webres/img/logo.png", // 76 x 98
		"obj":     obj,
	}
	return api, nil
}
