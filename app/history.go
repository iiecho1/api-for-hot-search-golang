package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func stripHTML(htmlString string) string {
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		return htmlString
	}
	var result strings.Builder
	var visit func(n *html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.TextNode {
			result.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(doc)
	return result.String()
}

func History() (map[string]interface{}, error) {
	currentTime := time.Now()
	month := fmt.Sprintf("%02d", currentTime.Month())
	day := fmt.Sprintf("%02d", currentTime.Day())
	url := "https://baike.baidu.com/cms/home/eventsOnHistory/" + month + ".json"

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var resultMap map[string]interface{}
	if err := json.Unmarshal(body, &resultMap); err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	monthData, ok := resultMap[month].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("API返回数据格式异常: 月份数据不存在")
	}

	date := month + day
	dateListInterface, ok := monthData[date]
	if !ok {
		return nil, fmt.Errorf("今天(%s月%s日)没有历史事件数据", month, day)
	}

	dateList, ok := dateListInterface.([]interface{})
	if !ok {
		return nil, fmt.Errorf("API返回数据格式异常: 日期列表格式不正确")
	}

	if len(dateList) == 0 {
		return utils.BuildErrorResponse("历史上的今天", "https://baike.baidu.com/favicon.ico",
			"今天("+month+"月"+day+"日)没有历史事件数据"), nil
	}

	obj := make([]map[string]interface{}, 0, len(dateList))
	for index, item := range dateList {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		titleInterface, ok := itemMap["title"]
		if !ok {
			continue
		}
		title, ok := titleInterface.(string)
		if !ok {
			continue
		}
		urlStr := ""
		if linkInterface, ok := itemMap["link"]; ok {
			if link, ok := linkInterface.(string); ok {
				urlStr = link
			}
		}
		obj = append(obj, utils.BuildItem(index+1, stripHTML(title), urlStr))
	}

	return utils.BuildSuccessResponse("历史上的今天", "https://baike.baidu.com/favicon.ico", obj), nil
}
