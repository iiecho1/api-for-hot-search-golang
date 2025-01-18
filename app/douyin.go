package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
)

func Douyin() map[string]interface{} {
	url := "https://www.iesdouyin.com/web/api/v2/hotsearch/billboard/word/"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap map[string]interface{}
	_ = json.Unmarshal(pageBytes, &resultMap)

	wordList := resultMap["word_list"]

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "抖音"

	var obj []map[string]interface{}
	for index, item := range wordList.([]interface{}) {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["word"]
		hot := item.(map[string]interface{})["hot_value"].(float64) / 10000
		result["hotValue"] = fmt.Sprint(math.Round(hot*10)/10) + "万"
		result["url"] = "https://www.douyin.com/search/" + item.(map[string]interface{})["word"].(string)
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
