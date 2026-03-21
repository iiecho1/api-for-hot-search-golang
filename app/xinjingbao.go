package app

import (
	"api/utils"
	"fmt"
)

func Xinjingbao() (map[string]interface{}, error) {
	url := "https://www.bjnews.com.cn/"

	pageContent, err := utils.GetHTML(url)
	if err != nil {
		return nil, fmt.Errorf("GetHTML error: %w", err)
	}

	pattern := `<h3>\s*<a class="link" href="([^"]+)"[^>]*>\s*<span[^>]*>\d*</span>\s*(.*?)</a>\s*</h3>[\s\S]*?</i>(.*?)</span>`
	matched := utils.ExtractMatches(pageContent, pattern)

	if len(matched) == 0 {
		return utils.BuildErrorResponse("新京报", "https://www.bjnews.com.cn/favicon.ico",
			"未匹配到数据，可能页面结构已变更"), nil
	}

	obj := make([]map[string]interface{}, 0, len(matched))
	for index, item := range matched {
		obj = append(obj, utils.BuildItem(index+1, item[2], item[1],
			map[string]string{"hotValue": item[3]}))
	}

	return utils.BuildSuccessResponse("新京报", "https://www.bjnews.com.cn/favicon.ico", obj), nil
}
