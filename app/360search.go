package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type search360Item struct {
	LongTitle string `json:"long_title"`
	Title     string `json:"title"`
	Score     string `json:"score"`
	Rank      string `json:"rank"`
}

func Search360() (map[string]interface{}, error) {
	url := "https://ranks.hao.360.com/mbsug-api/hotnewsquery?type=news&realhot_limit=50"

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var resultSlice []search360Item
	if err := json.Unmarshal(body, &resultSlice); err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	obj := make([]map[string]interface{}, 0, len(resultSlice))
	for i, item := range resultSlice {
		title := item.Title
		if item.LongTitle != "" {
			title = item.LongTitle
		}
		hot, _ := strconv.ParseFloat(item.Score, 64)
		entry := utils.BuildItem(i+1, title,
			"https://www.so.com/s?q="+title,
			map[string]string{"hotValue": fmt.Sprintf("%.1f万", hot/10000)})
		obj = append(obj, entry)
	}

	return utils.BuildSuccessResponse("360搜索",
		"https://ss.360tres.com/static/121a1737750aa53d.ico", obj), nil
}
