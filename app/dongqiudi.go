package app

import (
	"api/utils"
	"fmt"
)

type dqdResponse struct {
	Data dqdList `json:"data"`
}

type dqdList struct {
	NewList []dqdData `json:"new_list"`
}

type dqdData struct {
	Title string `json:"title"`
	URL   string `json:"share"`
}

func Dongqiudi() (map[string]interface{}, error) {
	url := "https://dongqiudi.com/api/v3/archive/pc/index/getIndex"

	var resultMap dqdResponse
	if err := utils.FetchJSON(url, &resultMap, nil); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(resultMap.Data.NewList) == 0 {
		return utils.BuildErrorResponse("懂球帝",
			"https://www.dongqiudi.com/images/dqd-logo.png",
			"API返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(resultMap.Data.NewList))
	for index, item := range resultMap.Data.NewList {
		obj = append(obj, utils.BuildItem(index+1, item.Title, item.URL))
	}

	return utils.BuildSuccessResponse("懂球帝", "https://www.dongqiudi.com/images/dqd-logo.png", obj), nil
}
