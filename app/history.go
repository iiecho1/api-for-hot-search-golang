package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func stripHTML(htmlString string) string {
	// 使用 html.Parse 解析 HTML 字符串
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		// 解析失败时返回原始字符串
		return htmlString
	}

	// 使用一个 buffer 保存结果
	var result strings.Builder

	// 定义一个递归函数来遍历 HTML 树
	var visit func(n *html.Node)
	visit = func(n *html.Node) {
		// 如果当前节点是文本节点，将文本内容追加到结果中
		if n.Type == html.TextNode {
			result.WriteString(n.Data)
		}
		// 递归处理子节点
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}

	// 调用递归函数开始遍历 HTML 树
	visit(doc)

	// 返回结果的字符串形式
	return result.String()
}

func History() (map[string]interface{}, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	currentTime := time.Now()
	month := fmt.Sprintf("%02d", currentTime.Month())
	day := fmt.Sprintf("%02d", currentTime.Day())
	url := "https://baike.baidu.com/cms/home/eventsOnHistory/" + month + ".json"

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	// 检查数据结构
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

	// 检查数据是否为空
	if len(dateList) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "今天(" + month + "月" + day + "日)没有历史事件数据",
			"icon":    "https://baike.baidu.com/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}

	var obj []map[string]interface{}
	for index, item := range dateList {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue // 跳过格式不正确的项
		}

		// 提取标题
		titleInterface, ok := itemMap["title"]
		if !ok {
			continue
		}
		title, ok := titleInterface.(string)
		if !ok {
			continue
		}

		// 提取链接
		urlInterface, ok := itemMap["link"]
		urlStr := ""
		if ok {
			if link, ok := urlInterface.(string); ok {
				urlStr = link
			}
		}

		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": stripHTML(title),
			"url":   urlStr,
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "历史上的今天",
		"icon":    "https://baike.baidu.com/favicon.ico", // 64 x 64
		"obj":     obj,
	}
	return api, nil
}
