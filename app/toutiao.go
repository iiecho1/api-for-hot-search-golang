package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type ttResponse struct {
	Data []ttData `json:"data"`
}
type ttData struct {
	Title    string `json:"Title"`
	URL      string `json:"Url"`
	HotValue string `json:"HotValue"`
}

func Toutiao() map[string]interface{} {
	url := "https://www.toutiao.com/hot-event/hot-board/?origin=toutiao_pc"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap ttResponse
	_ = json.Unmarshal(pageBytes, &resultMap)

	data := resultMap.Data

	var obj []map[string]interface{}
	for index, item := range data {
		hot, err := strconv.ParseFloat(item.HotValue, 64)
		utils.HandleError(err, "strconv.ParseFloat")
		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item.Title,
			"url":      item.URL,
			"hotValue": fmt.Sprintf("%.1f万", hot/10000),
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "今日头条",
		"icon":    "https://lf3-static.bytednsdoc.com/obj/eden-cn/pipieh7nupabozups/toutiao_web_pc/tt-icon.png", // 144 x 144
		"obj":     obj,
	}
	return api
}
