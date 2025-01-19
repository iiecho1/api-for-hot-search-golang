package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
)

func CCTV() map[string]interface{} {
	url := "https://news.cctv.com/2019/07/gaiban/cmsdatainterface/page/world_1.jsonp"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap map[string]interface{}
	// 删除多余字符，解析json
	_ = json.Unmarshal(pageBytes[6:len(pageBytes)-1], &resultMap)

	wordList := resultMap["data"].(map[string]interface{})["list"]

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "CCTV"
	var obj []map[string]interface{}

	for index, item := range wordList.([]interface{}) {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["title"]
		result["url"] = item.(map[string]interface{})["url"]
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
