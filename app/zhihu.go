package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type zhResponse struct {
	Response response `json:"recommend_queries"`
}
type response struct {
	Data []zhData `json:"queries"`
}
type zhData struct {
	Title string `json:"query"`
}

func Zhihu() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	urlStr := "https://www.zhihu.com/api/v4/search/recommend_query/v2"
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest error: %w", err)
	}

	// 设置请求头，模拟正常浏览器访问
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Referer", "https://www.zhihu.com/")

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

	var resultMap zhResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}
	// 检查数据是否为空
	if len(resultMap.Response.Data) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空",
			"icon":    "https://static.zhihu.com/static/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	data := resultMap.Response.Data
	var obj []map[string]interface{}
	for index, item := range data {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Title,
			"url":   "https://www.zhihu.com/search?q=" + item.Title,
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "知乎",
		"icon":    "https://static.zhihu.com/static/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api, nil
}
