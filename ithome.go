package utils

import (
	"io"
	"net/http"
)

func Ithome() map[string]interface{} {
	url := "https://m.ithome.com/rankm/"
	resp, err := http.Get(url)
	HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	HandleError(err, "io.ReadAll")
	pattern := `<p class="plc-title">(.*?)<\/p>.*?<a href="(.*?)"`
	matches := ExtractMatches(string(pageBytes), pattern)

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "It之家日榜"
	var obj []map[string]interface{}
	for index, item := range matches[:12] {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item[1]
		result["url"] = item[2]
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
