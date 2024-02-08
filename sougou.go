package utils

import (
	"io"
	"net/http"
)

func Sougou() map[string]interface{} {
	url := "https://www.sogou.com/web?query=%E6%90%9C%E7%8B%97%E7%83%AD%E6%90%9C"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)

	HandleError(err, "http.Get")
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")

	pattern := `<span [^>]*>[\s\S]*?<p>\s*<a href="([^"]+)"[^>]*>(.*?)</a>\s*</p>[\s\S]*?</span>\s*<span class="hot-rank-right">(.*?)</span>`
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
