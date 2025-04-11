package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
)

func Lishipin() map[string]interface{} {
	url := "https://www.pearvideo.com/popular"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")

	pattern := `<a\shref="(.*?)".*?>\s*<h2\sclass="popularem-title">(.*?)</h2>\s*<p\sclass="popularem-abs padshow">(.*?)</p>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	var obj []map[string]interface{}
	for index, item := range matched {
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": item[2],
			"url":   "https://www.pearvideo.com/" + fmt.Sprint(item[1]),
			"desc":  item[3],
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "梨视频",
		"icon":    "https://page.pearvideo.com/webres/img/logo.png", // 76 x 98
		"obj":     obj,
	}
	return api
}
