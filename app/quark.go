package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type quarkResponse struct {
	Data quarkData `json:"data"`
}
type quarkData struct {
	HotNews hotNews `json:"hotNews"`
}
type hotNews struct {
	Item []quarkItem `json:"item"`
}
type quarkItem struct {
	URL      string `json:"url"`
	Title    string `json:"title"`
	HotValue string `json:"hot"`
}

func Quark() map[string]interface{} {
	url := "https://biz.quark.cn/api/trending/ranking/getNewsRanking?modules=hotNews&uc_param_str=dnfrpfbivessbtbmnilauputogpintnwmtsvcppcprsnnnchmicckpgixsnx"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get error")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll error")
	var resultMap quarkResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal error")

	data := resultMap.Data.HotNews.Item
	obj := make([]map[string]interface{}, 0, len(data))

	for i, item := range data {
		hot, err := strconv.ParseFloat(item.HotValue, 64)
		if err != nil {
			hot = 0
		}

		obj = append(obj, map[string]interface{}{
			"index":    i + 1,
			"title":    item.Title,
			"url":      item.URL,
			"hotValue": fmt.Sprintf("%.1f万", hot/10000),
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "夸克",
		"obj":     obj,
		"icon":    "https://gw.alicdn.com/imgextra/i3/O1CN018r2tKf28YP7ev0fPF_!!6000000007944-2-tps-48-48.png",
	}
	return api
}
