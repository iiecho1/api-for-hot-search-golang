package app

import (
	"api/utils"
	"fmt"
	"net/url"
	"strings"
)

type doubanItem struct {
	Score float64 `json:"score"`
	Name  string  `json:"name"`
	URI   string  `json:"uri"`
}

func Douban() (map[string]interface{}, error) {
	apiURL := "https://m.douban.com/rexxar/api/v2/chart/hot_search_board?count=10&start=0"

	headers := map[string]string{
		"Referer": "https://www.douban.com/gallery/",
	}

	var items []doubanItem
	if err := utils.FetchJSON(apiURL, &items, headers); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(items) == 0 {
		return utils.BuildErrorResponse("豆瓣", "https://www.douban.com/favicon.ico",
			"API返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(items))
	for index, item := range items {
		hotValue := ""
		if item.Score > 0 {
			hotValue = fmt.Sprintf("%.2f万", item.Score/10000)
		}
		// douban:// 协议或无效链接转为 HTTPS 搜索链接
		link := item.URI
		if strings.HasPrefix(link, "douban://") || !strings.HasPrefix(link, "http") {
			// 尝试从 douban:// 链接提取 q= 参数
			if strings.HasPrefix(link, "douban://") {
				if u, err := url.Parse(link); err == nil {
					if q := u.Query().Get("q"); q != "" {
						link = "https://www.douban.com/search?q=" + url.QueryEscape(q)
					}
				}
			}
			// 如果仍然是 douban:// 或无效链接，使用名称作为搜索词
			if !strings.HasPrefix(link, "http") {
				link = "https://www.douban.com/search?q=" + url.QueryEscape(item.Name)
			}
		}
		obj = append(obj, utils.BuildItem(index+1, item.Name, link,
			map[string]string{"hotValue": hotValue}))
	}

	return utils.BuildSuccessResponse("豆瓣", "https://www.douban.com/favicon.ico", obj), nil
}
