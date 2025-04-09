package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
)

type response struct {
	Data dongqiudiData `json:"data"`
}
type dongqiudiData struct {
	NewList []data `json:"new_list"`
}
type data struct {
	Title string `json:"title"`
	URL   string `json:"share"`
}

func Dongqiudi() map[string]interface{} {
	url := "https://dongqiudi.com/api/v3/archive/pc/index/getIndex"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap response
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal")

	data := resultMap.Data.NewList

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "懂球帝"
	var obj []map[string]interface{}

	for index, item := range data {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.Title
		result["url"] = item.URL
		obj = append(obj, result)
	}
	api["obj"] = obj
	api["icon"] = "https://www.dongqiudi.com/images/dqd-logo.png" // 800 x 206
	return api
}
