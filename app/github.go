package app

import (
	"api/utils"
	"fmt"
	"regexp"
	"strings"
)

func Github() (map[string]interface{}, error) {
	url := "https://github.com/trending"

	pageContent, err := utils.GetHTML(url)
	if err != nil {
		return nil, fmt.Errorf("GetHTML error: %w", err)
	}

	// 按 <article> 标签拆分每个仓库
	articleRe := regexp.MustCompile(`(?s)<article[^>]*>(.*?)</article>`)
	articles := articleRe.FindAllStringSubmatch(pageContent, -1)

	if len(articles) == 0 {
		return utils.BuildErrorResponse("GitHub",
			"https://github.githubassets.com/favicons/favicon.png",
			"未匹配到 trending 数据，可能页面结构已变更"), nil
	}

	obj := make([]map[string]interface{}, 0, len(articles))
	for index, art := range articles {
		article := art[1]

		// 从 <h2> 中提取仓库路径 (格式: /user/repo)
		linkRe := regexp.MustCompile(`(?s)<h2[^>]*>\s*<a[^>]*href="(/[a-zA-Z0-9_./-]+)"[^>]*>`)
		linkMatch := linkRe.FindStringSubmatch(article)
		if len(linkMatch) < 2 {
			continue
		}
		repoPath := strings.TrimSpace(linkMatch[1])

		// 从 <p> 中提取描述
		descRe := regexp.MustCompile(`(?s)<p[^>]*class="[^"]*col-9[^"]*"[^>]*>(.*?)</p>`)
		descMatch := descRe.FindStringSubmatch(article)
		desc := ""
		if len(descMatch) >= 2 {
			desc = strings.TrimSpace(regexp.MustCompile(`<[^>]+>`).ReplaceAllString(descMatch[1], ""))
			desc = strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(desc, " "))
		}

		obj = append(obj, utils.BuildItem(index+1, repoPath[1:],
			"https://github.com"+repoPath,
			map[string]string{"desc": desc}))
	}

	return utils.BuildSuccessResponse("GitHub",
		"https://github.githubassets.com/favicons/favicon.png", obj), nil
}
