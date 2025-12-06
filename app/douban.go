package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type doubanItem struct {
	Score float64 `json:"score"`
	Name  string  `json:"name"`
	URI   string  `json:"uri"`
}

func Douban() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://m.douban.com/rexxar/api/v2/chart/hot_search_board?count=10&start=0"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest error: %w", err)
	}

	// 设置 Headers（模拟浏览器）
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Referer", "https://www.douban.com/gallery/")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.Client.Do error: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var items []doubanItem
	err = json.Unmarshal(pageBytes, &items)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	// 检查数据是否为空
	if len(items) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空",
			"icon":    "https://www.douban.com/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	var obj []map[string]interface{}
	for index, item := range items {
		hotValue := ""
		if item.Score > 0 {
			hotValue = fmt.Sprintf("%.2f万", item.Score/10000)
		}

		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item.Name,
			"url":      item.URI,
			"hotValue": hotValue,
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "豆瓣",
		"icon":    "https://www.douban.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api, nil
}
