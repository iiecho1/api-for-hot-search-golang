package app

import (
	"api/utils"
	"fmt"
)

func Sougou() (map[string]interface{}, error) {
	url := "https://www.sogou.com/web?query=%E6%90%9C%E7%8B%97%E7%83%AD%E6%90%9C"

	headers := map[string]string{
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
	}

	pageContent, err := utils.FetchHTML(url, headers)
	if err != nil {
		return nil, fmt.Errorf("FetchHTML error: %w", err)
	}

	pattern := `<span [^>]*>[\s\S]*?<p>\s*<a href="([^"]+)"[^>]*>(.*?)</a>\s*</p>[\s\S]*?</span>\s*<span class="hot-rank-right">(.*?)</span>`
	matched := utils.ExtractMatches(pageContent, pattern)

	if len(matched) == 0 {
		return utils.BuildErrorResponse("搜狗", "https://www.sogou.com/favicon.ico",
			"未匹配到数据，可能页面结构已变更"), nil
	}

	obj := make([]map[string]interface{}, 0, len(matched))
	for index, item := range matched {
		obj = append(obj, utils.BuildItem(index+1, item[2], item[1],
			map[string]string{"hotValue": item[3]}))
	}

	return utils.BuildSuccessResponse("搜狗", "https://www.sogou.com/favicon.ico", obj), nil
}
