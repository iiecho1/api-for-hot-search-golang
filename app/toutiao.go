package app

import (
	"api/utils"
	"fmt"
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

func Toutiao() (map[string]interface{}, error) {
	url := "https://www.toutiao.com/hot-event/hot-board/?origin=toutiao_pc"

	var resultMap ttResponse
	if err := utils.FetchJSON(url, &resultMap, nil); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(resultMap.Data) == 0 {
		return utils.BuildErrorResponse("今日头条",
			"https://lf3-static.bytednsdoc.com/obj/eden-cn/pipieh7nupabozups/toutiao_web_pc/tt-icon.png",
			"API返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(resultMap.Data))
	for index, item := range resultMap.Data {
		hot, err := strconv.ParseFloat(item.HotValue, 64)
		if err != nil {
			hot = 0
		}
		obj = append(obj, utils.BuildItem(index+1, item.Title, item.URL,
			map[string]string{"hotValue": fmt.Sprintf("%.1f万", hot/10000)}))
	}

	return utils.BuildSuccessResponse("今日头条",
		"https://lf3-static.bytednsdoc.com/obj/eden-cn/pipieh7nupabozups/toutiao_web_pc/tt-icon.png",
		obj), nil
}
