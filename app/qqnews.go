package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type qqResponse struct {
	IdList []idListItem `json:"idlist"`
}

type idListItem struct {
	IdsHash  string     `json:"ids_hash"`
	NewsList []newsItem `json:"newslist"`
}

type newsItem struct {
	Title    string   `json:"title"`
	Url      string   `json:"url"`
	Time     string   `json:"time"`
	HotEvent hotEvent `json:"hotEvent"`
}

type hotEvent struct {
	HotScore float64 `json:"hotScore"`
}

func Qqnews() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://r.inews.qq.com/gw/event/hot_ranking_list?page_size=51"
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

	// 使用结构体解析响应
	var result qqResponse
	err = json.Unmarshal(pageBytes, &result)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	// 检查数据是否为空或格式不正确
	if len(result.IdList) == 0 || len(result.IdList[0].NewsList) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空",
			"icon":    "https://mat1.gtimg.com/qqcdn/qqindex2021/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	// 获取新闻列表数据
	newsListData := result.IdList[0].NewsList

	var obj []map[string]interface{}
	for index, item := range newsListData {
		if index == 0 {
			continue
		}
		hot := item.HotEvent.HotScore / 10000
		hotValue := fmt.Sprintf("%.1f万", hot)

		obj = append(obj, map[string]interface{}{
			"index":    index,
			"title":    item.Title,
			"url":      item.Url,
			"time":     item.Time,
			"hotValue": hotValue,
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "腾讯新闻",
		"icon":    "https://mat1.gtimg.com/qqcdn/qqindex2021/favicon.ico", // 96 x 96
		"obj":     obj,
	}
	return api, nil
}
