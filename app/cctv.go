package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
)

type cctvResponse struct {
	Data cctvData `json:"data"`
}
type cctvData struct {
	List []cctvList `json:"list"`
}

type cctvList struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func CCTV() map[string]interface{} {
	url := "https://news.cctv.com/2019/07/gaiban/cmsdatainterface/page/world_1.jsonp"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap cctvResponse
	// 删除多余字符，解析json
	_ = json.Unmarshal(pageBytes[6:len(pageBytes)-1], &resultMap)

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "CCTV"
	var obj []map[string]interface{}

	for index, item := range resultMap.Data.List {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.Title
		result["url"] = item.URL
		obj = append(obj, result)
	}
	api["obj"] = obj
	api["icon"] = "https://news.cctv.com/favicon.ico" // 16 x 16
	return api
}
