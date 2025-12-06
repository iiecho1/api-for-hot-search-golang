package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Douyinresponse struct {
	WordList []Douyindata `json:"word_list"`
}

type Douyindata struct {
	Title    string  `json:"word"`
	HotVaule float64 `json:"hot_value"`
}

func Douyin() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	urlStr := "https://www.iesdouyin.com/web/api/v2/hotsearch/billboard/word/"
	resp, err := client.Get(urlStr)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var resultMap Douyinresponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	// 检查数据是否为空
	if len(resultMap.WordList) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空",
			"icon":    "https://lf1-cdn-tos.bytegoofy.com/goofy/ies/douyin_web/public/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	var obj []map[string]interface{}
	for index, item := range resultMap.WordList {
		// URL 编码标题，确保特殊字符正确处理
		encodedTitle := url.QueryEscape(item.Title)

		hotValue := ""
		if item.HotVaule > 0 {
			hotValue = fmt.Sprintf("%.2f万", item.HotVaule/10000)
		}

		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item.Title,
			"url":      "https://www.douyin.com/search/" + encodedTitle,
			"hotValue": hotValue,
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "抖音",
		"icon":    "https://lf1-cdn-tos.bytegoofy.com/goofy/ies/douyin_web/public/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api, nil
}
