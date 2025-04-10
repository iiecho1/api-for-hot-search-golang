package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type doubanItem struct {
	Score float64 `json:"score"`
	Name  string  `json:"name"`
	URI   string  `json:"uri"`
}

func Douban() map[string]interface{} {
	url := "https://m.douban.com/rexxar/api/v2/chart/hot_search_board?count=10&start=0"

	req, err := http.NewRequest("GET", url, nil)
	utils.HandleError(err, "io.ReadAll")
	// 设置 User-Agent（模拟浏览器）
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")

	refer := "https://www.douban.com/gallery/"
	req.Header.Set("Referer", refer)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	utils.HandleError(err, "io.ReadAll")
	defer resp.Body.Close()

	// 读取响应
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")

	var items []doubanItem
	_ = json.Unmarshal([]byte(string(pageBytes)), &items)

	var obj []map[string]interface{}
	for index, item := range items {
		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item.Name,
			"url":      item.URI,
			"hotValue": fmt.Sprintf("%.2f万", item.Score/10000),
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "豆瓣",
		"icom":    "https://www.douban.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api
}
