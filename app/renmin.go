package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Renminwang() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	url := "http://www.people.com.cn/GB/59476/index.html"
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

	pattern := `<li><a href="(.*?)" target="_blank">(.*?)</a></li>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)
	// 检查是否匹配到数据
	if len(matched) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "未匹配到数据，可能页面结构已变更",
			"icon":    "http://www.people.com.cn/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}
	var obj []map[string]interface{}
	for index, item := range matched {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item[2]
		result["url"] = item[1]
		obj = append(obj, result)
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "人民网",
		"icon":    "http://www.people.com.cn/favicon.ico", // 16 x 16
		"obj":     obj,
	}
	return api, nil
}
