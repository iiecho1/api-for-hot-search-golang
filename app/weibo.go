package app

import (
	"api/utils"
	"fmt"
	"regexp"
	"strings"
)

func WeiboHot() (map[string]interface{}, error) {
	url := "https://s.weibo.com/top/summary"

	headers := map[string]string{
		"Cookie":         "SUB=_2AkMasdasdqadTy2Pna4Rl77p7cJZAXC",
		"Accept":         "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Referer":        "https://s.weibo.com/",
	}

	pageContent, err := utils.FetchHTML(url, headers)
	if err != nil {
		return nil, fmt.Errorf("FetchHTML error: %w", err)
	}

	pattern := `<a href="(/weibo\?q=[^"]+)"[^>]*target="_blank">([^<]+)</a>\s*<span>([^<]*)?</span>`
	matched := utils.ExtractMatches(pageContent, pattern)

	nonDigitRegexp := regexp.MustCompile(`[^\d]`)
	var obj []map[string]interface{}
	for index, item := range matched {
		if len(item) >= 3 {
			title := strings.TrimSpace(item[2])
			link := "https://s.weibo.com" + item[1]
			hotValue := ""
			if len(item) >= 4 {
				hotValue = strings.TrimSpace(item[3])
			}
			if hotValue != "" {
				hotValue = strings.TrimSpace(nonDigitRegexp.ReplaceAllString(hotValue, ""))
			}
			obj = append(obj, utils.BuildItem(index+1, title, link,
				map[string]string{"hotValue": hotValue}))
		}
	}

	// 备用提取方法
	if len(obj) == 0 {
		obj = extractWeiboHotSearchFallback(pageContent)
	}

	if len(obj) == 0 {
		return utils.BuildErrorResponse("微博热搜", "https://weibo.com/favicon.ico",
			"无法提取热搜数据，页面结构可能已变更"), nil
	}

	return utils.BuildSuccessResponse("微博热搜", "https://weibo.com/favicon.ico", obj), nil
}

func extractWeiboHotSearchFallback(content string) []map[string]interface{} {
	var obj []map[string]interface{}
	patterns := []string{
		`<a href="(/weibo\?q=[^"]+)"[^>]*>([^<]+)</a>`,
		`class="td-02".*?<a href="(/weibo\?q=[^"]+)"[^>]*>([^<]+)</a>`,
	}

	for _, pattern := range patterns {
		matched := utils.ExtractMatches(content, pattern)
		for index, item := range matched {
			if len(item) >= 3 {
				title := strings.TrimSpace(item[2])
				link := "https://s.weibo.com" + item[1]
				obj = append(obj, utils.BuildItem(index+1, title, link))
			}
		}
		if len(obj) > 0 {
			break
		}
	}
	return obj
}
