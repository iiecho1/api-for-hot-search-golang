package app

import (
	"api/utils"
	"io"
	"net/http"
	"strings"
)

func Github() map[string]interface{} {
	url := "https://github.com/trending"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")

	pattern := `<span\s+data-view-component="true"\s+class="text-normal">\s*([^<]+)\s*<\/span>\s*([^<]+)<\/a>\s*<\/h2>\s*<p\sclass="col-9 color-fg-muted my-1 pr-4">\s*([^<]+)\s*<\/p>`
	matched := utils.ExtractMatches(string(pageBytes), pattern)

	var obj []map[string]interface{}
	for index, item := range matched {
		trimed := strings.ReplaceAll(strings.TrimSpace(item[1])+strings.TrimSpace(item[2]), " ", "")
		obj = append(obj, map[string]interface{}{
			"index": index + 1,
			"title": trimed,
			"desc":  strings.TrimSpace(item[3]),
			"url":   "https://github.com/" + trimed,
		})
	}
	api := map[string]interface{}{
		"code":    200,
		"message": "GitHub",
		"icon":    "https://github.githubassets.com/favicons/favicon.png", // 32 x 32
		"obj":     obj,
	}
	return api
}
