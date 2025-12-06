package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type bilibiliResponse struct {
	Data bilibiliList `json:"data"`
}

type bilibiliList struct {
	List []bilibiliData `json:"list"`
}
type bilibiliData struct {
	Title string `json:"title"`
	Bvid  string `json:"bvid"`
}

func Bilibili() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://api.bilibili.com/x/web-interface/ranking/v2?rid=0&type=all"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest error: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.Client.Do error: %w", err)
	}
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var resultMap bilibiliResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	// 检查数据是否为空
	if len(resultMap.Data.List) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空或格式不正确，实际返回数据：" + fmt.Sprintf("%+v", resultMap.Data),
		}, nil // 这里返回 nil error，因为这是业务逻辑错误，不是程序错误
	}

	var obj []map[string]interface{}
	for index, item := range resultMap.Data.List {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Title,
			"url":   "https://www.bilibili.com/video/" + item.Bvid,
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "哔哩哔哩",
		"icon":    "https://www.bilibili.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api, nil
}
