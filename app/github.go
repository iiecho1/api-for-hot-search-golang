package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func Github() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://github.com/trending"
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	pattern := `<span\s+data-view-component="true"\s+class="text-normal">\s*([^<]+)\s*<\/span>\s*([^<]+)<\/a>\s*<\/h2>\s*<p\sclass="col-9 color-fg-muted my-1 pr-4">\s*([^<]+)\s*<\/p>`
	// 注意：这里假设 utils.ExtractMatches 没有改变，如果它也返回 error，需要修改
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	// 检查是否匹配到数据
	if len(matched) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "未匹配到 trending 数据，可能页面结构已变更",
			"icon":    "https://github.githubassets.com/favicons/favicon.png",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	var obj []map[string]interface{}
	for index, item := range matched {
		// 添加边界检查
		if len(item) >= 4 {
			// 清理用户名和仓库名
			user := strings.TrimSpace(item[1])
			repo := strings.TrimSpace(item[2])
			// 合并并移除所有空格
			trimed := strings.ReplaceAll(user+repo, " ", "")

			desc := ""
			if len(item) >= 4 {
				desc = strings.TrimSpace(item[3])
			}

			obj = append(obj, map[string]interface{}{
				"index": index + 1,
				"title": trimed,
				"desc":  desc,
				"url":   "https://github.com/" + trimed,
			})
		}
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "GitHub",
		"icon":    "https://github.githubassets.com/favicons/favicon.png", // 32 x 32
		"obj":     obj,
	}
	return api, nil
}
