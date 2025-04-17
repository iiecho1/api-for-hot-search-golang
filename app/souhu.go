package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type souhuResponse struct {
	Data []newsArticles `json:"newsArticles"`
}
type newsArticles struct {
	Title string  `json:"title"`
	ID    float64 `json:"newsId"`
	Hot   string  `json:"score"`
}

func Souhu() map[string]interface{} {
	url := "https://3g.k.sohu.com/api/channel/hotchart/hotnews.go?page=1"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap souhuResponse
	_ = json.Unmarshal(pageBytes, &resultMap)

	wordList := resultMap.Data

	var obj []map[string]interface{}
	for index, item := range wordList {
		hotValue, err := strconv.ParseFloat(item.Hot, 64)
		utils.HandleError(err, "strconv.ParseFloat")
		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item.Title,
			"url":      "https://search.sohu.com/?queryType=edit&keyword=" + item.Title,
			"hotValue": fmt.Sprintf("%.2f万", hotValue),
		})

	}
	api := map[string]interface{}{
		"code":    200,
		"message": "搜狐新闻",
		"icon":    "https://3g.k.sohu.com/favicon.ico", // 48 x 48
		"obj":     obj,
	}
	return api
}
