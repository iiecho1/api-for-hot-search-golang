package app

import (
	"api/utils"
	"fmt"
)

func Lishipin() (map[string]interface{}, error) {
	url := "https://www.pearvideo.com/popular"

	pageContent, err := utils.GetHTML(url)
	if err != nil {
		return nil, fmt.Errorf("GetHTML error: %w", err)
	}

	pattern := `<a\shref="(.*?)".*?>\s*<h2\sclass="popularem-title">(.*?)</h2>\s*<p\sclass="popularem-abs padshow">(.*?)</p>`
	matched := utils.ExtractMatches(pageContent, pattern)

	if len(matched) == 0 {
		return utils.BuildErrorResponse("梨视频",
			"https://page.pearvideo.com/webres/img/logo.png",
			"未匹配到数据，可能页面结构已变更"), nil
	}

	obj := make([]map[string]interface{}, 0, len(matched))
	for index, item := range matched {
		obj = append(obj, utils.BuildItem(index+1, item[2],
			"https://www.pearvideo.com/"+item[1],
			map[string]string{"desc": item[3]}))
	}

	return utils.BuildSuccessResponse("梨视频",
		"https://page.pearvideo.com/webres/img/logo.png", obj), nil
}
