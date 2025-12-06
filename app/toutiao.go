package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type ttResponse struct {
	Data []ttData `json:"data"`
}
type ttData struct {
	Title    string `json:"Title"`
	URL      string `json:"Url"`
	HotValue string `json:"HotValue"`
}

func Toutiao() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	url := "https://www.toutiao.com/hot-event/hot-board/?origin=toutiao_pc"
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

	var resultMap ttResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}
	// 检查数据是否为空
	if len(resultMap.Data) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空",
			"icon":    "https://lf3-static.bytednsdoc.com/obj/eden-cn/pipieh7nupabozups/toutiao_web_pc/tt-icon.png",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	data := resultMap.Data
	var obj []map[string]interface{}
	for index, item := range data {
		parsedHot, err := strconv.ParseFloat(item.HotValue, 64)
		hot := 0.0
		if err != nil {
			// 解析失败时使用默认值
			hot = 0
		} else {
			hot = parsedHot
		}
		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item.Title,
			"url":      item.URL,
			"hotValue": fmt.Sprintf("%.1f万", hot/10000),
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "今日头条",
		"icon":    "https://lf3-static.bytednsdoc.com/obj/eden-cn/pipieh7nupabozups/toutiao_web_pc/tt-icon.png", // 144 x 144
		"obj":     obj,
	}
	return api, nil
}
