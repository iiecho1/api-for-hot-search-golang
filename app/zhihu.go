package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
)

type zhResponse struct {
	Response response `json:"recommend_queries"`
}
type response struct {
	Data []zhData `json:"queries"`
}
type zhData struct {
	Title string `json:"query"`
}

func Zhihu() map[string]interface{} {
	url := "https://www.zhihu.com/api/v4/search/recommend_query/v2"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap zhResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal")

	data := resultMap.Response.Data

	var obj []map[string]interface{}
	for index, item := range data {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Title,
			"url":   "https://www.zhihu.com/search?q=" + item.Title,
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "知乎",
		"icon":    "https://static.zhihu.com/static/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api
}
