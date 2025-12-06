package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Doc360() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "http://www.360doc.com/"
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	pattern := `<div class=" num\d* yzphlist hei"><a href="(.*?)".*?>(?:<span class="icon_yuan2"></span>)?(.*?)</a></div>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	var obj []map[string]interface{}
	for index, item := range matched {
		// 添加边界检查，确保有足够的匹配项
		if len(item) >= 3 {
			obj = append(obj, map[string]interface{}{
				"index": index + 1,
				"title": item[2],
				"url":   item[1],
			})
		}
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "360doc",
		"icon":    "http://www.360doc.com/favicon.ico", // 16 x 16
		"obj":     obj,
	}
	return api, nil
}
