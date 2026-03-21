package app

import (
	"api/utils"
	"fmt"
)

func Renminwang() (map[string]interface{}, error) {
	url := "http://www.people.com.cn/GB/59476/index.html"

	pageContent, err := utils.GetHTML(url)
	if err != nil {
		return nil, fmt.Errorf("GetHTML error: %w", err)
	}

	// 匹配 <a href="..." ...>标题</a>，兼容 target="_blank" rel="noopener" 和 target=_blank 等变体
	pattern := `<a\s+href="(http://[^"]+\.people\.com\.cn/[^"]+)"[^>]*>([^<]{5,})</a>`
	matched := utils.ExtractMatches(pageContent, pattern)

	if len(matched) == 0 {
		return utils.BuildErrorResponse("人民网", "http://www.people.com.cn/favicon.ico",
			"未匹配到数据，可能页面结构已变更"), nil
	}

	obj := make([]map[string]interface{}, 0, len(matched))
	for index, item := range matched {
		obj = append(obj, utils.BuildItem(index+1, item[2], item[1]))
	}

	return utils.BuildSuccessResponse("人民网", "http://www.people.com.cn/favicon.ico", obj), nil
}
