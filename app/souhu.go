package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
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
	url := fmt.Sprintf("https://3g.k.sohu.com/api/channel/hotchart/hotnews.go?p1=NjY2NjY2&page=%d", page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var resultMap souhuResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	if err != nil {
		return nil, err
	}

	return resultMap.Data, nil
}

func Souhu() map[string]interface{} {
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

	// 如果有错误发生，返回错误信息
	if len(fetchErrors) > 0 {
		return map[string]interface{}{
			"code":    500,
			"message": fmt.Sprintf("获取数据时出错: %v", fetchErrors),
		}
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
	return api
}
