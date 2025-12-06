package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type acfunResponse struct {
	RankList []acfunData `json:"rankList"`
}
type acfunData struct {
	Title string `json:"contentTitle"`
	URL   string `json:"shareUrl"`
}

func Acfun() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://www.acfun.cn/rest/pc-direct/rank/channel?channelId=&subChannelId=&rankLimit=30&rankPeriod=DAY"
	// 创建一个自定义请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest error: %w", err)
	}

	// 设置 Headers
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

	var resultMap acfunResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	var obj []map[string]interface{}
	for index, item := range resultMap.RankList {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Title,
			"url":   item.URL,
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "AcFun",
		"icon":    "https://cdn.aixifan.com/ico/favicon.ico", // 32 x 32
		"obj":     obj,
	}

	return api, nil
}
