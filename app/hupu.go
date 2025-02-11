package app

import (
	"api/utils"
	"io"
	"net/http"
)

func Hupu() map[string]interface{} {
	url := "https://www.hupu.com/"
	resp, err := http.Get(url)
	// fmt.Println(resp)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	pattern := `<a\s+href="([^"]+)"[^>]+>\s*<div[^>]+>\s*<div[^>]+>\d+</div>\s*<div[^>]+>(.*?)</div>`
	matches := utils.ExtractMatches(string(pageBytes), pattern)

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "虎扑"

	var obj []map[string]interface{}

	for index, item := range matches {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item[2]
		result["url"] = item[1]
		obj = append(obj, result)
	}
	api["obj"] = obj
	api["icon"] = "https://www.hupu.com/favicon.ico" // 32 x 32
	return api
}
