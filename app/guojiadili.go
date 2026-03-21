package app

import (
	"api/utils"
	"fmt"
)

func Guojiadili() (map[string]interface{}, error) {
	url := "https://www.dili360.com/"

	pageContent, err := utils.GetHTML(url)
	if err != nil {
		return nil, fmt.Errorf("GetHTML error: %w", err)
	}

	pattern := `<li>\s*<span>\d*</span>\s*<h3><a href="(.*?)" target="_blank">(.*?)</a>`
	matched := utils.ExtractMatches(pageContent, pattern)

	if len(matched) == 0 {
		return utils.BuildErrorResponse("国家地理", "https://www.dili360.com/favicon.ico",
			"未匹配到数据，可能页面结构已变更"), nil
	}

	obj := make([]map[string]interface{}, 0, len(matched))
	for index, item := range matched {
		if len(item) >= 3 {
			urlPath := item[1]
			fullURL := "http://www.dili360.com" + urlPath
			if len(urlPath) > 4 && urlPath[:4] == "http" {
				fullURL = urlPath
			}
			obj = append(obj, utils.BuildItem(index+1, item[2], fullURL))
		}
	}

	return utils.BuildSuccessResponse("国家地理", "http://www.dili360.com/favicon.ico", obj), nil
}
