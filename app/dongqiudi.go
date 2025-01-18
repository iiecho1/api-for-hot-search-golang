package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
)

func Dongqiudi() map[string]interface{} {
	url := "https://dongqiudi.com/api/v3/archive/pc/index/getIndex"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap map[string]interface{}
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal")

	data := resultMap["data"].(map[string]interface{})["new_list"]

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "懂球帝"
	var obj []map[string]interface{}

	for index, item := range data.([]interface{}) {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["title"]
		result["url"] = item.(map[string]interface{})["url"]
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
