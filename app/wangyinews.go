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

	var obj []map[string]interface{}
	for index, item := range matched {
		hot, err := strconv.ParseFloat(item[3], 64)
		utils.HandleError(err, "strconv.ParseFloat")
		obj = append(obj, map[string]interface{}{
			"index":    index + 1,
			"title":    item[2],
			"url":      item[1],
			"hotValue": fmt.Sprintf("%.1f万", hot/10000),
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "网易新闻",
		"icon":    "https://news.163.com/favicon.ico", // 16 x 16
		"obj":     obj,
	}
	return api
}
