package app

import (
	"api/utils"
	"fmt"
	"strconv"
)

func WangyiNews() (map[string]interface{}, error) {
	url := "https://news.163.com/"

	pageContent, err := utils.GetHTML(url)
	if err != nil {
		return nil, fmt.Errorf("GetHTML error: %w", err)
	}

	pattern := `<em>\d*</em>\s*<a href="([^"]+)"[^>]*>(.*?)</a>\s*<span>(\d*)</span>`
	matched := utils.ExtractMatches(pageContent, pattern)

	if len(matched) == 0 {
		return utils.BuildErrorResponse("网易新闻", "https://news.163.com/favicon.ico",
			"未匹配到数据，可能页面结构已变更"), nil
	}

	obj := make([]map[string]interface{}, 0, len(matched))
	for index, item := range matched {
		hot, _ := strconv.ParseFloat(item[3], 64)
		obj = append(obj, utils.BuildItem(index+1, item[2], item[1],
			map[string]string{"hotValue": fmt.Sprintf("%.1f万", hot/10000)}))
	}

	return utils.BuildSuccessResponse("网易新闻", "https://news.163.com/favicon.ico", obj), nil
}
