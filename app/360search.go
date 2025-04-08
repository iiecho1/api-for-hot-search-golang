package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type search360Item struct {
	URL       string `json:"url"`
	LongTitle string `json:"long_title"`
	Title     string `json:"title"`
	Score     string `json:"score"`
	Rank      string `json:"rank"`
}

func Search360() map[string]interface{} {
	url := "https://ranks.hao.360.com/mbsug-api/hotnewsquery?type=news&realhot_limit=50"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get error")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll error")

	var resultSlice []search360Item
	err = json.Unmarshal([]byte(string(pageBytes)), &resultSlice)
	utils.HandleError(err, "json.Unmarshal error")

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "360搜索"
	var obj []map[string]interface{}
	for _, item := range resultSlice {
		result := make(map[string]interface{})
		result["index"] = item.Rank

		if item.LongTitle == "" {
			result["title"] = item.Title
		} else {
			result["title"] = item.LongTitle
		}

		hot, err := strconv.ParseFloat(item.Score, 64)
		utils.HandleError(err, "strconv.ParseFloat")

		result["hotValue"] = fmt.Sprintf("%.1f", hot/10000) + "万"
		result["url"] = item.URL
		obj = append(obj, result)
	}
	api["obj"] = obj
	api["icon"] = "https://ss.360tres.com/static/121a1737750aa53d.ico" // 32 x 32
	return api
}
