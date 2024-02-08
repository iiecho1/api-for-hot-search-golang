package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
)

func Douyin() map[string]interface{} {
	url := "https://www.iesdouyin.com/web/api/v2/hotsearch/billboard/word/"
	resp, err := http.Get(url)
	HandleError(err, "http.Get")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	var resultMap map[string]interface{}
	err = json.Unmarshal(pageBytes, &resultMap)

	wordList := resultMap["word_list"]

	api := make(map[string]interface{})
	api["code"] = 200

	var obj []map[string]interface{}
	for index, item := range wordList.([]interface{}) {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["word"]
		hot := item.(map[string]interface{})["hot_value"].(float64) / 10000
		result["hotValue"] = fmt.Sprint(math.Round(hot*10)/10) + "万"
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
