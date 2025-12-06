package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type search360Item struct {
	LongTitle string `json:"long_title"`
	Title     string `json:"title"`
	Score     string `json:"score"`
	Rank      string `json:"rank"`
}

func Search360() (map[string]interface{}, error) {
	url := "https://ranks.hao.360.com/mbsug-api/hotnewsquery?type=news&realhot_limit=50"
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var resultSlice []search360Item
	err = json.Unmarshal(pageBytes, &resultSlice)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	var obj []map[string]interface{}
	for _, item := range resultSlice {
		title := item.Title
		if item.LongTitle != "" {
			title = item.LongTitle
		}

		hot, err := strconv.ParseFloat(item.Score, 64)
		if err != nil {
			// 处理转换错误，可以给默认值或者跳过该项
			hot = 0
			// 或者使用日志记录错误但不中断程序
			// log.Printf("parse hot value error for item %s: %v", title, err)
		}

		obj = append(obj, map[string]interface{}{
			"index":    item.Rank,
			"title":    title,
			"hotValue": fmt.Sprintf("%.1f万", hot/10000),
			"url":      "https://www.so.com/s?q=" + title,
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "360搜索",
		"icon":    "https://ss.360tres.com/static/121a1737750aa53d.ico",
		"obj":     obj,
	}
	return api, nil
}
