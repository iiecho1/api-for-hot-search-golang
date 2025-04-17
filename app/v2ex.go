package app

import (
	"api/utils"
	"io"
	"net/http"
)

func V2ex() map[string]interface{} {
	url := "https://www.v2ex.com"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	pattern := `<span class="item_hot_topic_title">\s*<a href="(.*?)">(.*?)<\/a>\s*<\/span>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	var obj []map[string]interface{}
	for index, item := range matched {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item[2]
		result["url"] = url + item[1]
		obj = append(obj, result)
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "V2EX",
		"icon":    "https://www.v2ex.com/static/img/icon_rayps_64.png", // 64 x 64
		"obj":     obj,
	}
	return api
}
