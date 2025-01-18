package app

import (
	"api/utils"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func WangyiNews() map[string]interface{} {
	url := "https://news.163.com/"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	pageBytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	pattern := `<em>\d*</em>\s*<a href="([^"]+)"[^>]*>(.*?)</a>\s*<span>(\d*)</span>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "网易新闻"

	var obj []map[string]interface{}

	for index, item := range matched {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item[2]
		result["url"] = item[1]
		hot, err := strconv.ParseFloat(item[3], 64)
		utils.HandleError(err, "strconv.ParseFloat")

		result["hotValue"] = fmt.Sprintf("%.1f万", hot/10000)
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
