package app

import (
	"api/utils"
	"fmt"
)

type sspResponse struct {
	Data  []sspItem `json:"data"`
	Total int       `json:"total"`
}

type sspItem struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
}

func Shaoshupai() (map[string]interface{}, error) {
	url := "https://sspai.com/api/v1/article/tag/page/get?limit=100000&tag=%E7%83%AD%E9%97%A8%E6%96%87%E7%AB%A0"

	var resultMap sspResponse
	if err := utils.FetchJSON(url, &resultMap, nil); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(resultMap.Data) == 0 {
		return utils.BuildErrorResponse("少数派",
			"https://cdn-static.sspai.com/favicon/sspai.ico",
			"API返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(resultMap.Data))
	for index, item := range resultMap.Data {
		obj = append(obj, utils.BuildItem(index+1, item.Title,
			fmt.Sprintf("https://sspai.com/post/%d", item.ID)))
	}

	return utils.BuildSuccessResponse("少数派",
		"https://cdn-static.sspai.com/favicon/sspai.ico", obj), nil
}
