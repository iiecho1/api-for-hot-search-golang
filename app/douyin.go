package app

import (
	"api/utils"
	"fmt"
)

type douyinResponse struct {
	WordList []douyinData `json:"word_list"`
}

type douyinData struct {
	Title    string  `json:"word"`
	HotValue float64 `json:"hot_value"`
}

func Douyin() (map[string]interface{}, error) {
	urlStr := "https://www.iesdouyin.com/web/api/v2/hotsearch/billboard/word/"

	var resultMap douyinResponse
	if err := utils.FetchJSON(urlStr, &resultMap, nil); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(resultMap.WordList) == 0 {
		return utils.BuildErrorResponse("抖音",
			"https://lf1-cdn-tos.bytegoofy.com/goofy/ies/douyin_web/public/favicon.ico",
			"API返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(resultMap.WordList))
	for index, item := range resultMap.WordList {
		hotValue := ""
		if item.HotValue > 0 {
			hotValue = fmt.Sprintf("%.2f万", item.HotValue/10000)
		}
		obj = append(obj, utils.BuildItem(index+1, item.Title,
			"https://www.douyin.com/search/"+item.Title,
			map[string]string{"hotValue": hotValue}))
	}

	return utils.BuildSuccessResponse("抖音",
		"https://lf1-cdn-tos.bytegoofy.com/goofy/ies/douyin_web/public/favicon.ico",
		obj), nil
}
