package app

import (
	"api/utils"
	"io"
	"net/http"
)

func Guojiadili() map[string]interface{} {
	url := "http://www.dili360.com/"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	pattern := `<li>\s*<span>\d*</span>\s*<h3><a href="(.*?)" target="_blank">(.*?)</a>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	var obj []map[string]interface{}
	for index, item := range matched {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item[2],
			"url":   "http://www.dili360.com" + item[1],
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "国家地理",
		"icon":    "http://www.dili360.com/favicon.ico", // 32 x 32
		"obj":     obj,
	}
	return api
}
