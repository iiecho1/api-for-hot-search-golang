package app

import (
	"api/utils"
	"fmt"
)

func Hupu() (map[string]interface{}, error) {
	url := "https://www.hupu.com/"

	pageContent, err := utils.GetHTML(url)
	if err != nil {
		return nil, fmt.Errorf("GetHTML error: %w", err)
	}

	pattern := `<a\s+href="([^"]+)"[^>]+>\s*<div[^>]+>\s*<div[^>]+>\d+</div>\s*<div[^>]+>(.*?)</div>`
	matched := utils.ExtractMatches(pageContent, pattern)

	if len(matched) == 0 {
		return utils.BuildErrorResponse("虎扑", "https://www.hupu.com/favicon.ico",
			"未匹配到数据，可能页面结构已变更"), nil
	}

	obj := make([]map[string]interface{}, 0, len(matched))
	for index, item := range matched {
		if len(item) >= 3 {
			link := item[1]
			if len(link) > 0 && link[0] == '/' {
				link = "https://www.hupu.com" + link
			}
			obj = append(obj, utils.BuildItem(index+1, item[2], link))
		}
	}

	return utils.BuildSuccessResponse("虎扑", "https://www.hupu.com/favicon.ico", obj), nil
}
