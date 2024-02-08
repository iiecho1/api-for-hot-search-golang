package utils

import (
	"io"
	"net/http"
)

func Xinjingbao() map[string]interface{} {
	url := "https://www.bjnews.com.cn/"
	resp, err := http.Get(url)
	HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")

	pattern := `<h3>\s*<a class="link" href="([^"]+)"[^>]*>\s*<span[^>]*>\d*</span>\s*(.*?)</a>\s*</h3>[\s\S]*?</i>(.*?)</span>`
	matched := ExtractMatches(string(pageBytes), pattern)

	api := make(map[string]interface{})
	api["code"] = 200
	var obj []map[string]interface{}
	for index, item := range matched {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item[2]
		result["url"] = item[1]
		result["hotValue"] = item[3]
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
