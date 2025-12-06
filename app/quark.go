package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type quarkResponse struct {
	Data quarkData `json:"data"`
}
type quarkData struct {
	HotNews hotNews `json:"hotNews"`
}
type hotNews struct {
	Item []quarkItem `json:"item"`
}
type quarkItem struct {
	URL      string `json:"url"`
	Title    string `json:"title"`
	HotValue string `json:"hot"`
}

func Quark() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://biz.quark.cn/api/trending/ranking/getNewsRanking?modules=hotNews&uc_param_str=dnfrpfbivessbtbmnilauputogpintnwmtsvcppcprsnnnchmicckpgixsnx"
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
	var resultMap quarkResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}
	// 检查数据是否为空
	if len(resultMap.Data.HotNews.Item) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空",
			"icon":    "https://gw.alicdn.com/imgextra/i3/O1CN018r2tKf28YP7ev0fPF_!!6000000007944-2-tps-48-48.png",
			"obj":     []map[string]interface{}{},
		}, nil
	}
	data := resultMap.Data.HotNews.Item
	obj := make([]map[string]interface{}, 0, len(data))

	for i, item := range data {
		hot, err := strconv.ParseFloat(item.HotValue, 64)
		if err != nil {
			hot = 0
		}

		obj = append(obj, map[string]interface{}{
			"index":    i + 1,
			"title":    item.Title,
			"url":      item.URL,
			"hotValue": fmt.Sprintf("%.1f万", hot/10000),
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "夸克",
		"obj":     obj,
		"icon":    "https://gw.alicdn.com/imgextra/i3/O1CN018r2tKf28YP7ev0fPF_!!6000000007944-2-tps-48-48.png",
	}
	return api, nil
}
