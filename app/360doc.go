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

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "360doc"

	var obj []map[string]interface{}

	for index, item := range matched {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item[2]
		result["url"] = item[1]
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
