package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
)

type acfunResponse struct {
	RankList []acfunData `json:"rankList"`
}
type acfunData struct {
	Title string `json:"contentTitle"`
	URL   string `json:"shareUrl"`
}

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
	var resultMap acfunResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal error")

	var obj []map[string]interface{}

	for index, item := range resultMap.RankList {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item.Title,
			"url":   item.URL,
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "AcFun",
		"icon":    "https://cdn.aixifan.com/ico/favicon.ico", // 32 x 32
		"obj":     obj,
	}

	return api
}
