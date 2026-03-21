package app

import (
	"api/utils"
	"fmt"
	"strings"
)

func Baidu() (map[string]interface{}, error) {
	url := "https://top.baidu.com/board?tab=realtime"

	pageContent, err := utils.GetHTML(url)
	if err != nil {
		return nil, fmt.Errorf("GetHTML error: %w", err)
	}

	pattern := `<div\sclass="c-single-text-ellipsis">(.*?)</div?`
	matched := utils.ExtractMatches(pageContent, pattern)

	obj := make([]map[string]interface{}, 0, len(matched))
	for index, item := range matched {
		if len(item) >= 2 {
			title := strings.TrimSpace(item[1])
			obj = append(obj, utils.BuildItem(index+1, title,
				"https://www.baidu.com/s?wd="+title))
		}
	}

	return utils.BuildSuccessResponse("百度", "https://www.baidu.com/favicon.ico", obj), nil
}
