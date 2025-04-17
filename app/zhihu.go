package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type zhResponse struct {
	Data []zhData `json:"data"`
}
type zhData struct {
	Target zhTarget `json:"target"`
}
type zhTarget struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func Zhihu() map[string]interface{} {
	url := "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap zhResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal")

	data := resultMap.Data

	var obj []map[string]interface{}
	for index, item := range data {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Target.Title,
			"url":   strings.Replace(item.Target.URL, "api.zhihu.com/questions", "www.zhihu.com/question", 1),
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
