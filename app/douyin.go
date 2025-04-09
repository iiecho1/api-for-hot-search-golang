package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Douyinresponse struct {
	WordList []Douyindata `json:"word_list"`
}

type Douyindata struct {
	Title    string  `json:"word"`
	HotVaule float64 `json:"hot_value"`
}

func Douyin() map[string]interface{} {
	url := "https://www.iesdouyin.com/web/api/v2/hotsearch/billboard/word/"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")

	var resultMap Douyinresponse
	_ = json.Unmarshal(pageBytes, &resultMap)

	wordList := resultMap.WordList

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "抖音"

	var obj []map[string]interface{}
	for index, item := range wordList {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.Title
		hot := item.HotVaule / 10000
		result["hotValue"] = fmt.Sprint(hot) + "万"
		result["url"] = "https://www.douyin.com/search/" + item.Title
		obj = append(obj, result)
	}
	api["obj"] = obj
	api["icon"] = "https://lf1-cdn-tos.bytegoofy.com/goofy/ies/douyin_web/public/favicon.ico" // 32 x 32
	return api
}
