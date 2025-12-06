package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type souhuResponse struct {
	Data []newsArticles `json:"newsArticles"`
}
type newsArticles struct {
	Title string `json:"title"`
	URL   string `json:"h5Link"`
	Hot   string `json:"score"`
}

func fetchSouhuPage(page int) ([]newsArticles, error) {
	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf("https://3g.k.sohu.com/api/channel/hotchart/hotnews.go?p1=NjY2NjY2&page=%d", page)
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch page %d: http.Get error: %w", page, err)
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch page %d: HTTP请求失败，状态码: %d", page, resp.StatusCode)
	}

	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fetch page %d: io.ReadAll error: %w", page, err)
	}

	var resultMap souhuResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, fmt.Errorf("fetch page %d: json.Unmarshal error: %w", page, err)
	}

	return resultMap.Data, nil
}

func Souhu() (map[string]interface{}, error) {
	var wordList []newsArticles
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var fetchErrors []error

	// 并发获取两个页面的数据
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			data, err := fetchSouhuPage(page)
			if err != nil {
				mutex.Lock()
				fetchErrors = append(fetchErrors, err)
				mutex.Unlock()
				return
			}
			mutex.Lock()
			wordList = append(wordList, data...)
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	// 如果有错误发生，检查是否至少获取到一些数据
	if len(fetchErrors) > 0 {
		// 如果完全没有获取到数据，返回错误
		if len(wordList) == 0 {
			return nil, fmt.Errorf("获取数据失败: %v", fetchErrors)
		}
		// 如果获取到部分数据，继续处理但记录错误
		fmt.Printf("部分页面获取失败，但继续处理已获取的数据: %v\n", fetchErrors)
	}
	// 检查数据是否为空
	if len(wordList) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空",
			"icon":    "https://3g.k.sohu.com/favicon.ico",
			"obj":     []map[string]interface{}{},
		}, nil
	}
	var obj []map[string]interface{}
	for index, item := range wordList {
		hotValue, err := strconv.ParseFloat(item.Hot, 64)
		if err != nil {
			hotValue = 0 // 如果解析失败，设置为0
		}
		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item.Title,
			"url":      item.URL,
			"hotValue": fmt.Sprintf("%.2f万", hotValue),
		})
	}

	api := map[string]interface{}{
		"code":    200,
		"message": "搜狐新闻",
		"icon":    "https://3g.k.sohu.com/favicon.ico", // 48 x 48
		"obj":     obj,
	}
	return api, nil
}
