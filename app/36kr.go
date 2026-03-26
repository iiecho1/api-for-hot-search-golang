package app

import (
	"api/utils"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type krRSS struct {
	Channel krChannel `xml:"channel"`
}

type krChannel struct {
	Items []krItem `xml:"item"`
}

type krItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

func Kr36() (map[string]interface{}, error) {
	url := "https://36kr.com/feed-newsflash"

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	var rss krRSS
	if err := xml.Unmarshal(body, &rss); err != nil {
		return nil, fmt.Errorf("XML解析失败: %w", err)
	}

	if len(rss.Channel.Items) == 0 {
		return utils.BuildErrorResponse("36氪", "https://36kr.com/favicon.ico",
			"RSS返回数据为空"), nil
	}

	// 清理描述中的HTML标签
	cleanHTML := func(s string) string {
		re := regexp.MustCompile(`<[^>]+>`)
		return utils.NormalizeTitle(re.ReplaceAllString(s, ""))
	}

	obj := make([]map[string]interface{}, 0, len(rss.Channel.Items))
	for index, item := range rss.Channel.Items {
		desc := cleanHTML(item.Description)
		if len(desc) > 100 {
			desc = desc[:100] + "..."
		}

		entry := utils.BuildItem(index+1, item.Title, item.Link,
			map[string]string{"desc": desc})
		obj = append(obj, entry)
	}

	return utils.BuildSuccessResponse("36氪", "https://36kr.com/favicon.ico", obj), nil
}
