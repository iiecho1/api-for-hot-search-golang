package app

import (
	"api/utils"
	"fmt"
)

type ppResponse struct {
	Data ppData `json:"data"`
}

type ppData struct {
	HotNews []ppNews `json:"hotNews"`
}

type ppNews struct {
	Title  string `json:"name"`
	ContId string `json:"contId"`
}

func Pengpai() (map[string]interface{}, error) {
	url := "https://cache.thepaper.cn/contentapi/wwwIndex/rightSidebar"

	var resultMap ppResponse
	if err := utils.FetchJSON(url, &resultMap, nil); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(resultMap.Data.HotNews) == 0 {
		return utils.BuildErrorResponse("澎湃新闻", "https://www.thepaper.cn/favicon.ico",
			"API返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(resultMap.Data.HotNews))
	for index, item := range resultMap.Data.HotNews {
		if item.ContId == "" {
			continue
		}
		obj = append(obj, utils.BuildItem(index+1, item.Title,
			"https://www.thepaper.cn/newsDetail_forward_"+item.ContId))
	}

	if len(obj) == 0 {
		return utils.BuildErrorResponse("澎湃新闻", "https://www.thepaper.cn/favicon.ico",
			"API返回数据格式异常，缺少必要字段"), nil
	}

	return utils.BuildSuccessResponse("澎湃新闻", "https://www.thepaper.cn/favicon.ico", obj), nil
}
