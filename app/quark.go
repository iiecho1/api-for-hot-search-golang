package app

import (
	"api/utils"
	"fmt"
	"strconv"
)

type quarkResponse struct {
	Data quarkData `json:"data"`
}

type quarkData struct {
	HotNews quarkHotNews `json:"hotNews"`
}

type quarkHotNews struct {
	Item []quarkItem `json:"item"`
}

type quarkItem struct {
	URL      string `json:"url"`
	Title    string `json:"title"`
	HotValue string `json:"hot"`
}

func Quark() (map[string]interface{}, error) {
	url := "https://biz.quark.cn/api/trending/ranking/getNewsRanking?modules=hotNews&uc_param_str=dnfrpfbivessbtbmnilauputogpintnwmtsvcppcprsnnnchmicckpgixsnx"

	var resultMap quarkResponse
	if err := utils.FetchJSON(url, &resultMap, nil); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(resultMap.Data.HotNews.Item) == 0 {
		return utils.BuildErrorResponse("夸克",
			"https://gw.alicdn.com/imgextra/i3/O1CN018r2tKf28YP7ev0fPF_!!6000000007944-2-tps-48-48.png",
			"API返回数据为空"), nil
	}

	data := resultMap.Data.HotNews.Item
	obj := make([]map[string]interface{}, 0, len(data))
	for i, item := range data {
		hot, _ := strconv.ParseFloat(item.HotValue, 64)
		obj = append(obj, utils.BuildItem(i+1, item.Title, item.URL,
			map[string]string{"hotValue": fmt.Sprintf("%.1f万", hot/10000)}))
	}

	return utils.BuildSuccessResponse("夸克",
		"https://gw.alicdn.com/imgextra/i3/O1CN018r2tKf28YP7ev0fPF_!!6000000007944-2-tps-48-48.png",
		obj), nil
}
