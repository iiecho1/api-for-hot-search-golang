package utils

import (
	"io"
	"net/http"
)

func V2ex() map[string]interface{} {
	url := "https://www.v2ex.com"
	resp, err := http.Get(url)
	HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	// fmt.Println(string(pageBytes))
	pattern := `<span class="item_hot_topic_title">\s<a href="(.*?)">(.*?)<\/a>\s<\/span>`
	matched := ExtractMatches(string(pageBytes), pattern)

	api := make(map[string]interface{})
	api["code"] = 200

	var obj []map[string]interface{}

	for index, item := range matched {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item[2]
		result["url"] = url + item[1]
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
