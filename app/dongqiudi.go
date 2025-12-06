package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type dqdresponse struct {
	Data dqdList `json:"data"`
}
type dqdList struct {
	NewList []dqddata `json:"new_list"`
}
type dqddata struct {
	Title string `json:"title"`
	URL   string `json:"share"`
}

func Dongqiudi() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://dongqiudi.com/api/v3/archive/pc/index/getIndex"
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var resultMap dqdresponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	// 检查数据是否为空
	if len(resultMap.Data.NewList) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空",
			"icon":    "https://www.dongqiudi.com/images/dqd-logo.png",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	var obj []map[string]interface{}
	for index, item := range resultMap.Data.NewList {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Title,
			"url":   item.URL,
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "懂球帝",
		"icon":    "https://www.dongqiudi.com/images/dqd-logo.png", // 800 x 206
		"obj":     obj,
	}
	return api, nil
}
