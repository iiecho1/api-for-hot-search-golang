package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type qqResponse struct {
	IdList []idListItem `json:"idlist"`
}

type idListItem struct {
	IdsHash  string     `json:"ids_hash"`
	NewsList []newsItem `json:"newslist"`
}

type newsItem struct {
	Title    string   `json:"title"`
	Url      string   `json:"url"`
	Time     string   `json:"time"`
	HotEvent hotEvent `json:"hotEvent"`
}

type hotEvent struct {
	HotScore float64 `json:"hotScore"`
}

func Qqnews() map[string]interface{} {
	url := "https://r.inews.qq.com/gw/event/hot_ranking_list?page_size=51"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")

	// 使用结构体解析响应
	var result qqResponse
	_ = json.Unmarshal(pageBytes, &result)

	// 获取新闻列表数据
	newsListData := result.IdList[0].NewsList

	var obj []map[string]interface{}
	for index, item := range newsListData {
		if index == 0 {
			continue
		}
		hot := item.HotEvent.HotScore / 10000
		hotValue := fmt.Sprintf("%.1f万", hot)

		obj = append(obj, map[string]interface{}{
			"index":    index,
			"title":    item.Title,
			"url":      item.Url,
			"time":     item.Time,
			"hotValue": hotValue,
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "腾讯新闻",
		"icon":    "https://mat1.gtimg.com/qqcdn/qqindex2021/favicon.ico", // 96 x 96
		"obj":     obj,
	}
	return api
}
