package app

import (
	"api/utils"
	"io"
	"net/http"
	"strings"
)

func Baidu() map[string]interface{} {
	url := "https://top.baidu.com/board?tab=realtime"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	pattern := `<div\sclass="c-single-text-ellipsis">(.*?)</div?`
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	var obj []map[string]interface{}
	for index, item := range matched {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": strings.TrimSpace(item[1]),
			"url":   "https://www.baidu.com/s?wd=" + strings.TrimSpace(item[1]),
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "百度",
		"icon":    "https://www.baidu.com/favicon.ico", // 64 x 64
		"obj":     obj,
	}
	return api
}
