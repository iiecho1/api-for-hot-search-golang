package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DoubanItem struct {
	Score     float64 `json:"score"`
	TrendFlag int     `json:"trend_flag"`
	Name      string  `json:"name"`
	URI       string  `json:"uri"`
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

	var items []DoubanItem
	_ = json.Unmarshal([]byte(string(pageBytes)), &items)

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "豆瓣"

	var obj []map[string]interface{}
	for index, item := range items {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.Name
		hot := item.Score
		result["hotValue"] = fmt.Sprint(hot/10000) + "万"
		result["url"] = item.URI
		obj = append(obj, result)
	}
	api["obj"] = obj
	api["icon"] = "https://www.douban.com/favicon.ico" // 32 x 32
	return api
}
