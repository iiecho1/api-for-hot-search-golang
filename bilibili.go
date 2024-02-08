package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Bilibili() map[string]interface{} {
	url := "https://api.bilibili.com/x/web-interface/ranking/v2?rid=0&type=all"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)

	HandleError(err, "http.Get")
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	var resultMap map[string]interface{}
	err = json.Unmarshal(pageBytes, &resultMap)
	HandleError(err, "json.Unmarshal error")

	data := resultMap["data"].(map[string]interface{})["list"]
	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "哔哩哔哩日榜"
	var obj []map[string]interface{}
	for index, item := range data.([]interface{}) {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["title"]
		result["url"] = "https://www.bilibili.com/video/" + fmt.Sprint(item.(map[string]interface{})["bvid"])
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
