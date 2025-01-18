package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Pengpai() map[string]interface{} {
	url := "https://cache.thepaper.cn/contentapi/wwwIndex/rightSidebar"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	resultMap := make(map[string]interface{})
	_ = json.Unmarshal(pageBytes, &resultMap)

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "澎湃新闻"

	data := resultMap["data"].(map[string]interface{})["hotNews"].([]interface{})

	var obj []map[string]interface{}

	for index, item := range data {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["name"]
		result["url"] = "https://www.thepaper.cn/newsDetail_forward_" + fmt.Sprint(item.(map[string]interface{})["contId"])
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
