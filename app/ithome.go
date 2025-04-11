package app

import (
	"api/utils"
	"io"
	"net/http"
)

func Ithome() map[string]interface{} {
	url := "https://m.ithome.com/rankm/"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	pattern := `<p class="plc-title">(.*?)<\/p>.*?<a href="(.*?)"`
	matches := utils.ExtractMatches(string(pageBytes), pattern)

	var obj []map[string]interface{}
	for index, item := range matches[:12] {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item[1],
			"url":   item[2],
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "IT之家",
		"icon":    "https://www.ithome.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api
}
