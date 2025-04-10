package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Douyinresponse struct {
	WordList []Douyindata `json:"word_list"`
}

type Douyindata struct {
	Title    string  `json:"word"`
	HotVaule float64 `json:"hot_value"`
}

func Douyin() map[string]interface{} {
	url := "https://www.iesdouyin.com/web/api/v2/hotsearch/billboard/word/"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")

	var resultMap Douyinresponse
	_ = json.Unmarshal(pageBytes, &resultMap)

	wordList := resultMap.WordList
	var obj []map[string]interface{}
	for index, item := range wordList {
		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item.Title,
			"url":      "https://www.douyin.com/search/" + item.Title,
			"hotValue": fmt.Sprintf("%.2f万", item.HotVaule/10000),
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "抖音",
		"icon":    "https://lf1-cdn-tos.bytegoofy.com/goofy/ies/douyin_web/public/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api
}
