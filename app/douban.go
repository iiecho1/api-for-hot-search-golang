package app

import (
	"api/utils"
	"fmt"
)

type doubanItem struct {
	Score float64 `json:"score"`
	Name  string  `json:"name"`
	URI   string  `json:"uri"`
}

func Douban() (map[string]interface{}, error) {
	url := "https://m.douban.com/rexxar/api/v2/chart/hot_search_board?count=10&start=0"

	headers := map[string]string{
		"Referer": "https://www.douban.com/gallery/",
	}

	var items []doubanItem
	if err := utils.FetchJSON(url, &items, headers); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(items) == 0 {
		return utils.BuildErrorResponse("豆瓣", "https://www.douban.com/favicon.ico",
			"API返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(items))
	for index, item := range items {
		hotValue := ""
		if item.Score > 0 {
			hotValue = fmt.Sprintf("%.2f万", item.Score/10000)
		}
		obj = append(obj, utils.BuildItem(index+1, item.Name, item.URI,
			map[string]string{"hotValue": hotValue}))
	}

	return utils.BuildSuccessResponse("豆瓣", "https://www.douban.com/favicon.ico", obj), nil
}
