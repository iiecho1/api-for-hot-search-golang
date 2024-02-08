package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func Dongqiudi() map[string]interface{} {
	url := "https://dongqiudi.com/api/v3/archive/pc/index/getIndex"
	resp, err := http.Get(url)
	HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	var resultMap map[string]interface{}
	err = json.Unmarshal(pageBytes, &resultMap)
	HandleError(err, "json.Unmarshal")

	data := resultMap["data"].(map[string]interface{})["new_list"]

	api := make(map[string]interface{})
	api["code"] = 200
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
