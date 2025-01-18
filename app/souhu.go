package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
)

func Souhu() map[string]interface{} {
	url := "https://3g.k.sohu.com/api/channel/hotchart/hotnews.go?page=1"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap map[string]interface{}
	_ = json.Unmarshal(pageBytes, &resultMap)

	wordList := resultMap["newsArticles"]

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "搜狐新闻"

	var obj []map[string]interface{}
	for index, item := range wordList.([]interface{}) {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["title"]
		hot, _ := strconv.ParseFloat(item.(map[string]interface{})["score"].(string), 64)
		result["hotValue"] = fmt.Sprint(math.Round(hot/1000)/10) + "万"
		result["url"] = "https://3g.k.sohu.com/t/n" + strconv.FormatFloat(item.(map[string]interface{})["newsId"].(float64), 'f', -1, 64) + "?serialId=94c07e6320f9c8a0b40c55059f0cedef"
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
