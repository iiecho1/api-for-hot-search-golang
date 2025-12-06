package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func WeiboHot() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	url := "https://s.weibo.com/top/summary"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest error: %w", err)
	}
	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Cookie", "SUB=_2AkMasdasdqadTy2Pna4Rl77p7cJZAXC")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Referer", "https://s.weibo.com/")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.Client.Do error: %w", err)
	}
	defer resp.Body.Close()
	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP请求失败,状态码: %d", resp.StatusCode)
	}
	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	pageContent := string(pageBytes)

	// 方法1：使用正则表达式提取热搜数据
	var obj []map[string]interface{}

	// 正则表达式匹配热搜条目
	pattern := `<a href="(/weibo\?q=[^"]+)"[^>]*target="_blank">([^<]+)</a>\s*<span>([^<]*)?</span>`
	matched := utils.ExtractMatches(pageContent, pattern)
	// 预编译正则表达式，避免重复编译
	nonDigitRegexp := regexp.MustCompile(`[^\d]`)
	for index, item := range matched {
		if len(item) >= 3 {
			title := strings.TrimSpace(item[2])
			url := "https://s.weibo.com" + item[1]
			hotValue := ""
			if len(item) >= 4 {
				hotValue = strings.TrimSpace(item[3])
			}

			// 清理热度值中的非数字字符
			if hotValue != "" {
				hotValue = strings.TrimSpace(nonDigitRegexp.ReplaceAllString(hotValue, ""))
			}

			obj = append(obj, map[string]interface{}{
				"index":    index + 1,
				"title":    title,
				"url":      url,
				"hotValue": hotValue,
			})
		}
	}

	// 如果正则匹配失败，尝试备用方法
	if len(obj) == 0 {
		obj = extractWeiboHotSearchFallback(pageContent)
	}
	// 检查是否获取到数据
	if len(obj) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "无法提取热搜数据，页面结构可能已变更",
			"icon":    "https://weibo.com/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "微博热搜",
		"icon":    "https://weibo.com/favicon.ico",
		"obj":     obj,
	}
	return api, nil
}

// 备用提取方法
func extractWeiboHotSearchFallback(content string) []map[string]interface{} {
	var obj []map[string]interface{}

	// 尝试匹配更简单的模式
	patterns := []string{
		`<a href="(/weibo\?q=[^"]+)"[^>]*>([^<]+)</a>`,
		`class="td-02".*?<a href="(/weibo\?q=[^"]+)"[^>]*>([^<]+)</a>`,
	}

	for _, pattern := range patterns {
		matched := utils.ExtractMatches(content, pattern)
		for index, item := range matched {
			if len(item) >= 3 {
				title := strings.TrimSpace(item[2])
				url := "https://s.weibo.com" + item[1]

				obj = append(obj, map[string]interface{}{
					"index":    index + 1,
					"title":    title,
					"url":      url,
					"hotValue": "", // 备用方法可能无法获取热度值
				})
			}
		}
		if len(obj) > 0 {
			break
		}
	}

	return obj
}
