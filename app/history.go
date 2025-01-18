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
	// 使用 html.Parse 解析 HTML 字符串
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		fmt.Println("解析HTML时出错:", err)
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
func History() map[string]interface{} {
	currentTime := time.Now()
	month := fmt.Sprintf("%02d", currentTime.Month())
	day := fmt.Sprintf("%02d", currentTime.Day())
	url := "https://baike.baidu.com/cms/home/eventsOnHistory/" + fmt.Sprint(month) + ".json"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll error")
	var resultMap map[string]interface{}
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "io.json.Unmarshal error")

	date := fmt.Sprintf("%v%v", month, day)
	dateList := resultMap[month].(map[string]interface{})[date]

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "历史上的今天"

	var obj []map[string]interface{}
	for index, item := range dateList.([]interface{}) {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = stripHTML(item.(map[string]interface{})["title"].(string))
		result["url"] = item.(map[string]interface{})["link"]
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
