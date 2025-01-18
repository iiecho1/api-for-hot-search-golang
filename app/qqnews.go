package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Qqnews() map[string]interface{} {
	url := "https://r.inews.qq.com/gw/event/hot_ranking_list?page_size=51"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	resultMap := make(map[string]interface{})
	_ = json.Unmarshal(pageBytes, &resultMap)

	newslist := resultMap["idlist"].([]interface{})[0].(map[string]interface{})["newslist"].([]interface{})
	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "腾讯新闻"

	var obj []map[string]interface{}

	for index, item := range newslist {
		if index > 0 {
			result := make(map[string]interface{})
			result["index"] = index
			result["title"] = item.(map[string]interface{})["title"]
			result["url"] = item.(map[string]interface{})["url"]
			result["time"] = item.(map[string]interface{})["time"]
			hot := item.(map[string]interface{})["hotEvent"].(map[string]interface{})["hotScore"].(float64) / 10000
			result["hotValue"] = fmt.Sprintf("%.1f", hot) + "万"
			obj = append(obj, result)
		}
	}
	api["obj"] = obj
	return api
}
