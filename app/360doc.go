package app

import (
	"api/utils"
	"io"
	"net/http"
)

func Doc360() map[string]interface{} {
	url := "http://www.360doc.com/"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")

	pattern := `<div class=" num\d* yzphlist hei"><a href="(.*?)".*?>(?:<span class="icon_yuan2"></span>)?(.*?)</a></div>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	var obj []map[string]interface{}
	for index, item := range matched {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item[2],
			"url":   item[1],
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "360doc",
		"icon":    "http://www.360doc.com/favicon.ico", // 16 x 16
		"obj":     obj,
	}
	return api
}
