package app

import (
	"api/utils"
	"fmt"
)

func Ithome() (map[string]interface{}, error) {
	url := "https://m.ithome.com/rankm/"

	pageContent, err := utils.GetHTML(url)
	if err != nil {
		return nil, fmt.Errorf("GetHTML error: %w", err)
	}

	pattern := `<a href="(https://m\.ithome\.com/html/\d+\.htm)"[^>]*>[\s\S]*?<p class="plc-title">([^<]+)</p>`
	matched := utils.ExtractMatches(pageContent, pattern)

	if len(matched) == 0 {
		return utils.BuildErrorResponse("IT之家", "https://www.ithome.com/favicon.ico",
			"未匹配到数据，可能页面结构已变更"), nil
	}

	count := len(matched)
	if count > 12 {
		count = 12
	}

	obj := make([]map[string]interface{}, 0, count)
	for index := 0; index < count; index++ {
		item := matched[index]
		if len(item) >= 3 {
			obj = append(obj, utils.BuildItem(index+1, item[2], item[1]))
		}
	}

	if len(obj) == 0 {
		return utils.BuildErrorResponse("IT之家", "https://www.ithome.com/favicon.ico",
			"处理后的数据为空"), nil
	}

	return utils.BuildSuccessResponse("IT之家", "https://www.ithome.com/favicon.ico", obj), nil
}
