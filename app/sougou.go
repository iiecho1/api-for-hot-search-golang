package app

import (
	"api/utils"
	"fmt"
	"strconv"
)

type sougouResponse struct {
	Main []sougouItem `json:"main"`
}

type sougouItem struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Score string `json:"score"`
}

func Sougou() (map[string]interface{}, error) {
	apiURL := "https://hotlist.imtt.qq.com/Fetch"

	var result sougouResponse
	if err := utils.FetchJSON(apiURL, &result, nil); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(result.Main) == 0 {
		return utils.BuildErrorResponse("搜狗", "https://www.sogou.com/favicon.ico",
			"API返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(result.Main))
	for index, item := range result.Main {
		hotValue := ""
		if score, err := strconv.ParseFloat(item.Score, 64); err == nil {
			hotValue = fmt.Sprintf("%.1f万", score)
		}
		obj = append(obj, utils.BuildItem(index+1, item.Title, item.URL,
			map[string]string{"hotValue": hotValue}))
	}

	return utils.BuildSuccessResponse("搜狗", "https://www.sogou.com/favicon.ico", obj), nil
}
