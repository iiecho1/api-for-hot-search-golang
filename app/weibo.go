package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func WeiboHot() map[string]interface{} {
	url := "https://weibo.com/ajax/side/hotSearch"
	// 1.去网站拿数据
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get error")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll error")
	var resultMap map[string]interface{}
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal error")

	realtimeList := resultMap["data"].(map[string]interface{})["realtime"].([]interface{})
	// 遍历结果
	json := make(map[string]interface{})
	json["code"] = 200
	json["success"] = "success"
	json["message"] = "微博"

	obj := []map[string]interface{}{}
	for key, value := range realtimeList {
		result := make(map[string]interface{})
		result["id"] = key + 1
		result["title"] = value.(map[string]interface{})["note"]
		result["url"] = "https://s.weibo.com/weibo?q=" + strings.Replace(fmt.Sprint(result["title"]), " ", "%20", -1)
		result["hotValue"] = value.(map[string]interface{})["raw_hot"]
		obj = append(obj, result)
	}
	json["obj"] = obj
	return json
}
