package app

import (
	"api/utils"
	"fmt"
	"strconv"
)

type nfResponse struct {
	Data nfData `json:"data"`
}

type nfData struct {
	HotContents []nfContent `json:"hot_contents"`
}

type nfContent struct {
	Title string  `json:"subject"`
	ID    float64 `json:"id"`
}

func Nanfangzhoumo() (map[string]interface{}, error) {
	url := "https://www.infzm.com/hot_contents?format=json"

	var resultMap nfResponse
	if err := utils.FetchJSON(url, &resultMap, nil); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(resultMap.Data.HotContents) == 0 {
		return utils.BuildErrorResponse("南方周末", "https://www.infzm.com/favicon.ico",
			"API返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(resultMap.Data.HotContents))
	for index, item := range resultMap.Data.HotContents {
		obj = append(obj, utils.BuildItem(index+1, item.Title,
			"https://www.infzm.com/contents/"+strconv.FormatFloat(item.ID, 'f', -1, 64)))
	}

	return utils.BuildSuccessResponse("南方周末", "https://www.infzm.com/favicon.ico", obj), nil
}
