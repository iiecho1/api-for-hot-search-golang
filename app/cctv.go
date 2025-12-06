package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type cctvResponse struct {
	Data cctvData `json:"data"`
}
type cctvData struct {
	List []cctvList `json:"list"`
}

type cctvList struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func CCTV() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://news.cctv.com/2019/07/gaiban/cmsdatainterface/page/world_1.jsonp"
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	// 检查响应长度是否足够
	if len(pageBytes) <= 6 {
		return nil, fmt.Errorf("API返回数据长度不足")
	}

	var resultMap cctvResponse
	// 删除 JSONP 回调函数包裹，解析实际 JSON 数据
	err = json.Unmarshal(pageBytes[6:len(pageBytes)-1], &resultMap)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	// 检查数据是否为空
	if len(resultMap.Data.List) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空",
			"icon":    "https://news.cctv.com/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	var obj []map[string]interface{}
	for index, item := range resultMap.Data.List {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Title,
			"url":   item.URL,
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "CCTV新闻",
		"icon":    "https://news.cctv.com/favicon.ico", // 16 x 16
		"obj":     obj,
	}
	return api, nil
}
