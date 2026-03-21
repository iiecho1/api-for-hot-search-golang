package app

import (
	"api/utils"
	"fmt"
)

type bilibiliResponse struct {
	Data bilibiliList `json:"data"`
}

type bilibiliList struct {
	List []bilibiliData `json:"list"`
}

type bilibiliData struct {
	Title string `json:"title"`
	Bvid  string `json:"bvid"`
}

func Bilibili() (map[string]interface{}, error) {
	url := "https://api.bilibili.com/x/web-interface/ranking/v2?rid=0&type=all"

	var resultMap bilibiliResponse
	if err := utils.FetchJSON(url, &resultMap, nil); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(resultMap.Data.List) == 0 {
		return utils.BuildErrorResponse("哔哩哔哩",
			"https://www.bilibili.com/favicon.ico",
			"API返回数据为空或格式不正确，实际返回数据："+fmt.Sprintf("%+v", resultMap.Data)), nil
	}

	obj := make([]map[string]interface{}, 0, len(resultMap.Data.List))
	for index, item := range resultMap.Data.List {
		obj = append(obj, utils.BuildItem(index+1, item.Title,
			"https://www.bilibili.com/video/"+item.Bvid))
	}

	return utils.BuildSuccessResponse("哔哩哔哩", "https://www.bilibili.com/favicon.ico", obj), nil
}
