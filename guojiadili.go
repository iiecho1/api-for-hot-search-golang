package utils

import (
	"io"
	"net/http"
)

func Guojiadili() map[string]interface{} {
	url := "http://www.dili360.com/"
	resp, err := http.Get(url)
	HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	pattern := `<li>\s*<span>\d*</span>\s*<h3><a href="(.*?)" target="_blank">(.*?)</a>`
	matched := ExtractMatches(string(pageBytes), pattern)

	api := make(map[string]interface{})
	api["code"] = 200

	var obj []map[string]interface{}

	for index, item := range matched {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item[2]
		result["url"] = "http://www.dili360.com" + item[1]
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
