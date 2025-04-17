package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type wbResponse struct {
	Item wbItem `json:"data"`
}
type wbItem struct {
	Data []wbData `json:"realtime"`
}
type wbData struct {
	Title    string  `json:"word"`
	HotValue float64 `json:"num"`
}

func WeiboHot() map[string]interface{} {
	url := "https://weibo.com/ajax/side/hotSearch"
	// 1.去网站拿数据
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get error")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll error")
	var resultMap wbResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal error")

	realtimeList := resultMap.Item.Data

	obj := []map[string]interface{}{}
	for index, item := range realtimeList {
		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item.Title,
			"url":      "https://s.weibo.com/weibo?q=" + item.Title,
			"hotValue": fmt.Sprintf("%.1f万", item.HotValue/10000),
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "微博",
		"icon":    "https://www.weibo.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api
}
