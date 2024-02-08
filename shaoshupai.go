package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Shaoshupai() map[string]interface{} {
	url := "https://sspai.com/api/v1/article/tag/page/get?limit=100000&tag=%E7%83%AD%E9%97%A8%E6%96%87%E7%AB%A0"
	resp, err := http.Get(url)
	HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(pageBytes, &resultMap)

	data := resultMap["data"].([]interface{})
	api := make(map[string]interface{})
	api["code"] = 200

	var obj []map[string]interface{}

	for index, item := range data {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["title"]
		result["url"] = "https://sspai.com/post/" + fmt.Sprint(item.(map[string]interface{})["id"])
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
