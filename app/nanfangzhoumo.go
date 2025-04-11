package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type nfResponse struct {
	NfzmData nfData `json:"data"`
}
type nfData struct {
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
	var resultMap nfResponse
	_ = json.Unmarshal(pageBytes, &resultMap)

	wordList := resultMap.NfzmData.HotContents

	var obj []map[string]interface{}
	for index, item := range wordList {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Title,
			"url":   "https://www.infzm.com/contents/" + strconv.FormatFloat(item.ID, 'f', -1, 64),
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "南方周末",
		"icon":    "https://www.infzm.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api
}
