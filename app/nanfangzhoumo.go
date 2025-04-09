package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Response struct {
	NfzmData Data `json:"data"`
}
type Data struct {
	HotContents []contents `json:"hot_contents"`
}
type contents struct {
	Title string  `json:"subject"`
	ID    float64 `json:"id"`
}

func Nanfangzhoumo() map[string]interface{} {
	url := "https://www.infzm.com/hot_contents?format=json"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap Response
	_ = json.Unmarshal(pageBytes, &resultMap)

	wordList := resultMap.NfzmData.HotContents

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "南方周末"

	var obj []map[string]interface{}
	for index, item := range wordList {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.Title
		result["url"] = "https://www.infzm.com/contents/" + strconv.FormatFloat(item.ID, 'f', -1, 64)
		obj = append(obj, result)
	}
	api["obj"] = obj
	api["icon"] = "https://www.infzm.com/favicon.ico" // 32 x 32
	return api
}
