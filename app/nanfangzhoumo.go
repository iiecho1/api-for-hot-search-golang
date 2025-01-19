package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func Nanfangzhoumo() map[string]interface{} {
	url := "https://www.infzm.com/hot_contents?format=json"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap map[string]interface{}
	_ = json.Unmarshal(pageBytes, &resultMap)

	wordList := resultMap["data"].(map[string]interface{})["hot_contents"]

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "南方周末"

	var obj []map[string]interface{}
	for index, item := range wordList.([]interface{}) {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["subject"]
		result["url"] = "https://www.infzm.com/contents/" + strconv.FormatFloat(item.(map[string]interface{})["id"].(float64), 'f', -1, 64)
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
