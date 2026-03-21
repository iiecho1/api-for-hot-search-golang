package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
)

type cctvResponse struct {
	Data cctvData `json:"data"`
}

type cctvData struct {
	List []cctvItem `json:"list"`
}

type cctvItem struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func CCTV() (map[string]interface{}, error) {
	url := "https://news.cctv.com/2019/07/gaiban/cmsdatainterface/page/world_1.jsonp"

	// JSONP 响应需要手动处理
	pageContent, err := utils.GetHTML(url)
	if err != nil {
		return nil, fmt.Errorf("GetHTML error: %w", err)
	}

	if len(pageContent) <= 6 {
		return nil, fmt.Errorf("API返回数据长度不足")
	}

	// 去除 JSONP 回调包装
	var resultMap cctvResponse
	if err := json.Unmarshal([]byte(pageContent[6:len(pageContent)-1]), &resultMap); err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	if len(resultMap.Data.List) == 0 {
		return utils.BuildErrorResponse("CCTV新闻", "https://news.cctv.com/favicon.ico",
			"API返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(resultMap.Data.List))
	for index, item := range resultMap.Data.List {
		obj = append(obj, utils.BuildItem(index+1, item.Title, item.URL))
	}

	return utils.BuildSuccessResponse("CCTV新闻", "https://news.cctv.com/favicon.ico", obj), nil
}
