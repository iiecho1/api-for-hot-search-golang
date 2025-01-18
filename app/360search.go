package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func Search360() map[string]interface{} {
	url := "https://ranks.hao.360.com/mbsug-api/hotnewsquery?type=news&realhot_limit=50"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get error")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll error")
	var resultSlice []map[string]interface{}
	err = json.Unmarshal(pageBytes, &resultSlice)
	utils.HandleError(err, "json.Unmarshal error")

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "360搜索"
	var obj []map[string]interface{}
	for _, item := range resultSlice {
		result := make(map[string]interface{})
		result["index"] = item["rank"]

		if item["long_title"] == "" {
			result["title"] = item["title"]
		} else {
			result["title"] = item["long_title"]
		}

		hot, err := strconv.ParseFloat(item["score"].(string), 64)
		utils.HandleError(err, "strconv.ParseFloat")

		result["hotValue"] = fmt.Sprintf("%.1f", hot/10000) + "万"
		result["url"] = item["url"]
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
