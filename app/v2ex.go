package app

import (
	"api/utils"
	"fmt"
)

func V2ex() (map[string]interface{}, error) {
	url := "https://www.v2ex.com"

	pageContent, err := utils.GetHTML(url)
	if err != nil {
		return nil, fmt.Errorf("GetHTML error: %w", err)
	}

	pattern := `<span class="item_hot_topic_title">\s*<a href="(.*?)">(.*?)<\/a>\s*<\/span>`
	matched := utils.ExtractMatches(pageContent, pattern)

	if len(matched) == 0 {
		return utils.BuildErrorResponse("V2EX",
			"https://www.v2ex.com/static/img/icon_rayps_64.png",
			"未匹配到数据，可能页面结构已变更"), nil
	}

	obj := make([]map[string]interface{}, 0, len(matched))
	for index, item := range matched {
		obj = append(obj, utils.BuildItem(index+1, item[2], url+item[1]))
	}

	return utils.BuildSuccessResponse("V2EX",
		"https://www.v2ex.com/static/img/icon_rayps_64.png", obj), nil
}
