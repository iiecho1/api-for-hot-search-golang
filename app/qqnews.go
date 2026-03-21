package app

import (
	"api/utils"
	"fmt"
)

type qqResponse struct {
	IdList []qqIdListItem `json:"idlist"`
}

type qqIdListItem struct {
	NewsList []qqNewsItem `json:"newslist"`
}

type qqNewsItem struct {
	Title    string      `json:"title"`
	Url      string      `json:"url"`
	Time     string      `json:"time"`
	HotEvent qqHotEvent  `json:"hotEvent"`
}

type qqHotEvent struct {
	HotScore float64 `json:"hotScore"`
}

func Qqnews() (map[string]interface{}, error) {
	url := "https://r.inews.qq.com/gw/event/hot_ranking_list?page_size=51"

	var result qqResponse
	if err := utils.FetchJSON(url, &result, nil); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(result.IdList) == 0 || len(result.IdList[0].NewsList) == 0 {
		return utils.BuildErrorResponse("腾讯新闻",
			"https://mat1.gtimg.com/qqcdn/qqindex2021/favicon.ico",
			"API返回数据为空"), nil
	}

	newsListData := result.IdList[0].NewsList
	obj := make([]map[string]interface{}, 0, len(newsListData))
	for index, item := range newsListData {
		if index == 0 {
			continue
		}
		hot := item.HotEvent.HotScore / 10000
		hotValue := fmt.Sprintf("%.1f万", hot)

		entry := utils.BuildItem(index, item.Title, item.Url,
			map[string]string{"hotValue": hotValue})
		entry["time"] = item.Time
		obj = append(obj, entry)
	}

	return utils.BuildSuccessResponse("腾讯新闻",
		"https://mat1.gtimg.com/qqcdn/qqindex2021/favicon.ico", obj), nil
}
