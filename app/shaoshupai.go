package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Shaoshupai() map[string]interface{} {
	url := "https://sspai.com/api/v1/article/tag/page/get?limit=100000&tag=%E7%83%AD%E9%97%A8%E6%96%87%E7%AB%A0"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	resultMap := make(map[string]interface{})
	_ = json.Unmarshal(pageBytes, &resultMap)

	data := resultMap["data"].([]interface{})
	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "少数派"

	var obj []map[string]interface{}

	for index, item := range data {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["title"]
		result["url"] = "https://sspai.com/post/" + fmt.Sprint(item.(map[string]interface{})["id"])
		obj = append(obj, result)
	}
	api["obj"] = obj
	api["icon"] = "https://cdn-static.sspai.com/favicon/sspai.ico" // 64 x 64
	return api
}
