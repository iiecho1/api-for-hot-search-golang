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

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "Github"

	var obj []map[string]interface{}

	for index, item := range matched {
		result := make(map[string]interface{})
		result["index"] = index + 1
		trimed := strings.ReplaceAll(strings.TrimSpace(item[1])+strings.TrimSpace(item[2]), " ", "")
		result["title"] = trimed
		result["desc"] = strings.TrimSpace(item[3])
		result["url"] = "https://github.com/" + trimed
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
