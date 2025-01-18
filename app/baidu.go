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

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "百度"

	var obj []map[string]interface{}

	for index, item := range matched {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = strings.TrimSpace(item[1])
		result["url"] = "https://www.baidu.com/s?wd=" + result["title"].(string)
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
