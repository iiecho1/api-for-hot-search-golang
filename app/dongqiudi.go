package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
)

type dqdresponse struct {
	Data dqdList `json:"data"`
}
type dqdList struct {
	NewList []dqddata `json:"new_list"`
}
type dqddata struct {
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
	var resultMap dqdresponse
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal")

	data := resultMap.Data.NewList

	var obj []map[string]interface{}
	for index, item := range data {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Title,
			"url":   item.URL,
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "懂球帝",
		"icon":    "https://www.dongqiudi.com/images/dqd-logo.png", // 800 x 206
		"obj":     obj,
	}
	return api
}
