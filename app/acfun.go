package app

import (
	"api/utils"
	"fmt"
)

type acfunResponse struct {
	RankList []acfunData `json:"rankList"`
}

type acfunData struct {
	Title string `json:"contentTitle"`
	URL   string `json:"shareUrl"`
}

func Acfun() (map[string]interface{}, error) {
	url := "https://www.acfun.cn/rest/pc-direct/rank/channel?channelId=&subChannelId=&rankLimit=30&rankPeriod=DAY"

	var resultMap acfunResponse
	if err := utils.FetchJSON(url, &resultMap, nil); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	obj := make([]map[string]interface{}, 0, len(resultMap.RankList))
	for index, item := range resultMap.RankList {
		obj = append(obj, utils.BuildItem(index+1, item.Title, item.URL))
	}

	return utils.BuildSuccessResponse("AcFun", "https://cdn.aixifan.com/ico/favicon.ico", obj), nil
}
