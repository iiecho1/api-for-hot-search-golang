package app

import (
	"api/utils"
	"io"
	"net/http"
)

func Douban() map[string]interface{} {
	url := "https://www.douban.com/gallery/"

	req, err := http.NewRequest("GET", url, nil)
	utils.HandleError(err, "http.NewRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)

	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")

	pattern := `<a href="([^"]+)"[^>]*>(.*?)</a>\s*<span>(.*?)次浏览</span>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "豆瓣"

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
	api["icon"] = "https://www.douban.com/favicon.ico" // 32 x 32
	return api
}
