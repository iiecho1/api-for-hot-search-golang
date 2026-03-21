package app

import (
	"api/utils"
	"fmt"
)

type zhResponse struct {
	Response zhData `json:"recommend_queries"`
}

type zhData struct {
	Queries []zhItem `json:"queries"`
}

type zhItem struct {
	Title string `json:"query"`
}

func Zhihu() (map[string]interface{}, error) {
	urlStr := "https://www.zhihu.com/api/v4/search/recommend_query/v2"

	headers := map[string]string{
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Referer":         "https://www.zhihu.com/",
	}

	var resultMap zhResponse
	if err := utils.FetchJSON(urlStr, &resultMap, headers); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(resultMap.Response.Queries) == 0 {
		return utils.BuildErrorResponse("知乎",
			"https://static.zhihu.com/static/favicon.ico",
			"API返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(resultMap.Response.Queries))
	for index, item := range resultMap.Response.Queries {
		obj = append(obj, utils.BuildItem(index+1, item.Title,
			"https://www.zhihu.com/search?q="+item.Title))
	}

	return utils.BuildSuccessResponse("知乎", "https://static.zhihu.com/static/favicon.ico", obj), nil
}
