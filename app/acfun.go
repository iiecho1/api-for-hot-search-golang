package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
)

func Acfun() map[string]interface{} {
	url := "https://www.acfun.cn/rest/pc-direct/rank/channel?channelId=&subChannelId=&rankLimit=30&rankPeriod=DAY"
	// 创建一个自定义请求
	req, err := http.NewRequest("GET", url, nil)
	utils.HandleError(err, "http.NewRequest")
	// 设置 Headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	utils.HandleError(err, "http.DefaultClient.Do")
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll error")
	var resultMap map[string]interface{}
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal error")
	rankList := resultMap["rankList"]

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "AcFun"
	var obj []map[string]interface{}

	for index, item := range rankList.([]interface{}) {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["contentTitle"]
		result["url"] = item.(map[string]interface{})["shareUrl"]
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
