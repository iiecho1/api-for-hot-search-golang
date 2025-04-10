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

	var obj []map[string]interface{}
	for _, item := range resultSlice {
		title := item.Title
		if item.LongTitle != "" {
			title = item.LongTitle
		}
		hot, err := strconv.ParseFloat(item.Score, 64)
		utils.HandleError(err, "strconv.ParseFloat")

		obj = append(obj, map[string]interface{}{
			"index":    item.Rank,
			"title":    title,
			"hotValue": fmt.Sprintf("%.1f万", hot/10000),
			"url":      item.URL,
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "夸克",
		"icon":    "https://ss.360tres.com/static/121a1737750aa53d.ico",
		"obj":     obj,
	}
	return api
}
