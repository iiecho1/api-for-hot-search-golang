package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type nfResponse struct {
	NfzmData nfData `json:"data"`
}
type nfData struct {
	HotContents []contents `json:"hot_contents"`
}
type contents struct {
	Title string  `json:"subject"`
	ID    float64 `json:"id"`
}

func Nanfangzhoumo() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://www.infzm.com/hot_contents?format=json"
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
	var resultMap nfResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}
	wordList := resultMap.NfzmData.HotContents
	// 检查数据是否为空
	if len(wordList) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空",
			"icon":    "https://www.infzm.com/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}
	var obj []map[string]interface{}
	for index, item := range wordList {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Title,
			"url":   "https://www.infzm.com/contents/" + strconv.FormatFloat(item.ID, 'f', -1, 64),
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "南方周末",
		"icon":    "https://www.infzm.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api, nil
}
