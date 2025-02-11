package app

import (
	"api/utils"
	"io"
	"net/http"
)

func Ithome() map[string]interface{} {
	url := "https://m.ithome.com/rankm/"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	pattern := `<p class="plc-title">(.*?)<\/p>.*?<a href="(.*?)"`
	matches := utils.ExtractMatches(string(pageBytes), pattern)

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "It之家"
	var obj []map[string]interface{}
	for index, item := range matches[:12] {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item[1]
		result["url"] = item[2]
		obj = append(obj, result)
	}
	api["obj"] = obj
	api["icon"] = "https://www.ithome.com/favicon.ico" // 32 x 32
	return api
}
