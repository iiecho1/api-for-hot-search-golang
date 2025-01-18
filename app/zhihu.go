package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Zhihu() map[string]interface{} {
	url := "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal")

	data := resultMap["data"]

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "知乎"

	var obj []map[string]interface{}

	for index, item := range data.([]interface{}) {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["target"].(map[string]interface{})["title"]
		id := item.(map[string]interface{})["target"].(map[string]interface{})["id"]
		result["url"] = "https://www.zhihu.com/question/" + fmt.Sprintf("%.f", id)
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
