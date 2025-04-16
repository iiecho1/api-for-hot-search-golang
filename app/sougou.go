package app

import (
	"api/utils"
	"io"
	"net/http"
)

func Sougou() map[string]interface{} {
	url := "https://www.sogou.com/web?query=%E6%90%9C%E7%8B%97%E7%83%AD%E6%90%9C"
	req, err := http.NewRequest("GET", url, nil)
	utils.HandleError(err, "http.NewRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)

	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")

	pattern := `<span [^>]*>[\s\S]*?<p>\s*<a href="([^"]+)"[^>]*>(.*?)</a>\s*</p>[\s\S]*?</span>\s*<span class="hot-rank-right">(.*?)</span>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	var obj []map[string]interface{}
	for index, item := range matched {
		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item[2],
			"url":      item[1],
			"hotValue": item[3],
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "搜狗",
		"icon":    "https://www.sogou.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api
}
