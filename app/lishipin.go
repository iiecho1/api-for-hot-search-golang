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

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "梨视频"

	var obj []map[string]interface{}

	for index, item := range matched {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item[2]
		result["url"] = "https://www.pearvideo.com/" + fmt.Sprint(item[1])
		result["desc"] = item[3]
		obj = append(obj, result)
	}
	api["obj"] = obj
	api["icon"] = "https://page.pearvideo.com/webres/img/logo.png" // 76 x 98
	return api
}
