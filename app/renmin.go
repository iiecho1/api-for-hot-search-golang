package app

import (
	"api/utils"
	"fmt"
	"regexp"
)

func Renminwang() (map[string]interface{}, error) {
	url := "http://www.people.com.cn/GB/59476/index.html"

	pageContent, err := utils.GetHTML(url)
	if err != nil {
		return nil, fmt.Errorf("GetHTML error: %w", err)
	}

	// 提取 id="ta_1" 的今日要闻区域
	sectionRe := regexp.MustCompile(`(?s)id="ta_1"[^>]*>(.*?)(?:</table>|id="ta_2")`)
	sectionMatch := sectionRe.FindStringSubmatch(pageContent)
	if len(sectionMatch) < 2 {
		return utils.BuildErrorResponse("人民网", "http://www.people.com.cn/favicon.ico",
			"未找到今日要闻区域"), nil
	}
	section := sectionMatch[1]

	// 提取文章链接（排除根域名的导航链接）
	linkRe := regexp.MustCompile(`(?s)<a\s+href="(http://[^"]+)"[^>]*>([^<]{3,})</a>`)
	matched := linkRe.FindAllStringSubmatch(section, -1)

	obj := make([]map[string]interface{}, 0, len(matched))
	for _, item := range matched {
		link := item[1]
		title := item[2]

		// 过滤根域名链接（导航条目）
		if link == "http://www.people.com.cn/" || link == "http://www.people.com.cn" {
			continue
		}
		obj = append(obj, utils.BuildItem(len(obj)+1, title, link))
	}

	if len(obj) == 0 {
		return utils.BuildErrorResponse("人民网", "http://www.people.com.cn/favicon.ico",
			"今日要闻为空"), nil
	}

	return utils.BuildSuccessResponse("人民网", "http://www.people.com.cn/favicon.ico", obj), nil
}
