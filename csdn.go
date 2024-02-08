package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func CSDN() map[string]interface{} {
	url := "https://blog.csdn.net/phoenix/web/blog/hotRank?&pageSize=100"
	resp, err := http.Get(url)
	HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	var resultMap map[string]interface{}
	err = json.Unmarshal(pageBytes, &resultMap)
	HandleError(err, "json.Umarshal")
	data := resultMap["data"]

	api := make(map[string]interface{})
	api["code"] = 200
	var obj []map[string]interface{}

	for index, item := range data.([]interface{}) {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["articleTitle"]
		result["url"] = item.(map[string]interface{})["articleDetailUrl"]
		hot, err := strconv.ParseFloat(item.(map[string]interface{})["hotRankScore"].(string), 64)
		HandleError(err, "strconv.ParseFloat")
		result["hotValue"] = fmt.Sprintf("%.1f", hot/10000) + "万"
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
