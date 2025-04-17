package app

import (
	"api/utils"
	"io"
	"net/http"
)

func Xinjingbao() map[string]interface{} {
	url := "https://www.bjnews.com.cn/"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")

	pattern := `<h3>\s*<a class="link" href="([^"]+)"[^>]*>\s*<span[^>]*>\d*</span>\s*(.*?)</a>\s*</h3>[\s\S]*?</i>(.*?)</span>`
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
		"message": "新京报",
		"icon":    "https://www.bjnews.com.cn/favicon.ico", // 20 x 20
		"obj":     obj,
	}
	return api
}
