package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func Toutiao() map[string]interface{} {
	url := "https://www.toutiao.com/hot-event/hot-board/?origin=toutiao_pc"
	resp, err := http.Get(url)
	HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	resultMap := make(map[string]interface{})
	err = json.Unmarshal(pageBytes, &resultMap)

	data := resultMap["data"].([]interface{})
	api := make(map[string]interface{})
	api["code"] = 200
	var obj []map[string]interface{}

	for index, item := range data {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["Title"]
		result["url"] = item.(map[string]interface{})["Url"]
		hot, err := strconv.ParseFloat(item.(map[string]interface{})["HotValue"].(string), 64)
		HandleError(err, "strconv.ParseFloat")
		result["hotValue"] = fmt.Sprint(fmt.Sprintf("%.1f", hot/10000)) + "万"
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
