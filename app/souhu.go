package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type souhuResponse struct {
	Data []souhuArticle `json:"newsArticles"`
}

type souhuArticle struct {
	Title string `json:"title"`
	URL   string `json:"h5Link"`
	Hot   string `json:"score"`
}

func fetchSouhuPage(page int) ([]souhuArticle, error) {
	url := fmt.Sprintf("https://3g.k.sohu.com/api/channel/hotchart/hotnews.go?p1=NjY2NjY2&page=%d", page)

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch page %d: http.Get error: %w", page, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch page %d: HTTP请求失败，状态码: %d", page, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fetch page %d: io.ReadAll error: %w", page, err)
	}

	var resultMap souhuResponse
	if err := json.Unmarshal(body, &resultMap); err != nil {
		return nil, fmt.Errorf("fetch page %d: json.Unmarshal error: %w", page, err)
	}

	return resultMap.Data, nil
}

func Souhu() (map[string]interface{}, error) {
	var wordList []souhuArticle
	var wg sync.WaitGroup
	var mu sync.Mutex
	var fetchErrors []error

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			data, err := fetchSouhuPage(page)
			if err != nil {
				mu.Lock()
				fetchErrors = append(fetchErrors, err)
				mu.Unlock()
				return
			}
			mu.Lock()
			wordList = append(wordList, data...)
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	if len(fetchErrors) > 0 && len(wordList) == 0 {
		return nil, fmt.Errorf("获取数据失败: %v", fetchErrors)
	}

	if len(wordList) == 0 {
		return utils.BuildErrorResponse("搜狐新闻", "https://3g.k.sohu.com/favicon.ico",
			"API返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(wordList))
	for index, item := range wordList {
		hot, _ := strconv.ParseFloat(item.Hot, 64)
		obj = append(obj, utils.BuildItem(index+1, item.Title, item.URL,
			map[string]string{"hotValue": fmt.Sprintf("%.2f万", hot)}))
	}

	return utils.BuildSuccessResponse("搜狐新闻", "https://3g.k.sohu.com/favicon.ico", obj), nil
}
