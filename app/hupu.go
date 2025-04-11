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

	var obj []map[string]interface{}
	for index, item := range matches {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item[2],
			"url":   item[1],
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "虎扑",
		"icon":    "https://www.hupu.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api
}
