package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
)

type csdbResponse struct {
	Data []csdnData `json:"data"`
}
type csdnData struct {
	Title    string `json:"articleTitle"`
	URL      string `json:"articleDetailUrl"`
	HotValue string `json:"pcHotRankScore"`
}

func CSDN() map[string]interface{} {
	url := "https://blog.csdn.net/phoenix/web/blog/hotRank?&pageSize=100"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap csdbResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Umarshal")
	data := resultMap.Data

	var obj []map[string]interface{}
	for index, item := range data {
		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item.Title,
			"url":      item.URL,
			"hotValue": item.HotValue,
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "CSDN",
		"icon":    "https://csdnimg.cn/public/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api
}
