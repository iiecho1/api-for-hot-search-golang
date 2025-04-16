package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type sspResponse struct {
	Data []sspData `json:"data"`
}
type sspData struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
}

func Shaoshupai() map[string]interface{} {
	url := "https://sspai.com/api/v1/article/tag/page/get?limit=100000&tag=%E7%83%AD%E9%97%A8%E6%96%87%E7%AB%A0"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap sspResponse
	_ = json.Unmarshal(pageBytes, &resultMap)

	data := resultMap.Data

	var obj []map[string]interface{}
	for index, item := range data {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Title,
			"url":   "https://sspai.com/post/" + fmt.Sprint(item.ID),
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "少数派",
		"icon":    "https://cdn-static.sspai.com/favicon/sspai.ico", // 64 x 64
		"obj":     obj,
	}
	return api
}
