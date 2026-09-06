package app

import (
	"api/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type kr36Response struct {
	Code int        `json:"code"`
	Data kr36Data   `json:"data"`
}

type kr36Data struct {
	HotRankList []kr36Item `json:"hotRankList"`
}

type kr36Item struct {
	ItemId           int64              `json:"itemId"`
	TemplateMaterial kr36TemplateMaterial `json:"templateMaterial"`
}

type kr36TemplateMaterial struct {
	WidgetTitle  string `json:"widgetTitle"`
	WidgetImage  string `json:"widgetImage"`
	AuthorName   string `json:"authorName"`
	StatCollect  int    `json:"statCollect"`
	StatPraise   int    `json:"statPraise"`
}

func Kr36() (map[string]interface{}, error) {
	url := "https://gateway.36kr.com/api/mis/nav/home/nav/rank/hot"

	payload := map[string]interface{}{
		"partner_id": "wap",
		"param": map[string]interface{}{
			"siteId":     1,
			"platformId": 2,
		},
		"timestamp": time.Now().UnixMilli(),
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("JSON序列化失败: %w", err)
	}

	client := utils.DefaultClient()
	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	var result kr36Response
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("JSON解析失败: %w", err)
	}

	if result.Code != 0 {
		return utils.BuildErrorResponse("36氪", "https://36kr.com/favicon.ico",
			fmt.Sprintf("36Kr API返回错误: code=%d", result.Code)), nil
	}

	items := result.Data.HotRankList
	if len(items) == 0 {
		return utils.BuildErrorResponse("36氪", "https://36kr.com/favicon.ico",
			"返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(items))
	for index, item := range items {
		itemId := fmt.Sprintf("%d", item.ItemId)
		title := item.TemplateMaterial.WidgetTitle
		link := fmt.Sprintf("https://www.36kr.com/p/%s", itemId)
		hotValue := fmt.Sprintf("%d", item.TemplateMaterial.StatCollect)

		entry := utils.BuildItem(index+1, title, link,
			map[string]string{"hotValue": hotValue})
		obj = append(obj, entry)
	}

	return utils.BuildSuccessResponse("36氪", "https://36kr.com/favicon.ico", obj), nil
}
