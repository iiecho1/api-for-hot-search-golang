package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
)

type ppResponse struct {
	Data ppData `json:"data"`
}
type ppData struct {
	HotNews []news `json:"hotNews"`
}
type news struct {
	Title  string `json:"name"`
	ContId string `json:"contId"`
}

func Pengpai() map[string]interface{} {
	url := "https://cache.thepaper.cn/contentapi/wwwIndex/rightSidebar"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap ppResponse
	_ = json.Unmarshal(pageBytes, &resultMap)

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "澎湃新闻"

	data := resultMap.Data.HotNews

	var obj []map[string]interface{}

	for index, item := range data {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.Title
		result["url"] = "https://www.thepaper.cn/newsDetail_forward_" + item.ContId
		obj = append(obj, result)
	}
	api["obj"] = obj
	api["icon"] = "https://www.thepaper.cn/favicon.ico" // 32 x 32
	return api
}
