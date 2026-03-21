package app

import (
	"api/utils"
	"fmt"
)

func Doc360() (map[string]interface{}, error) {
	url := "http://www.360doc.com/"

	pageContent, err := utils.GetHTML(url)
	if err != nil {
		return nil, fmt.Errorf("GetHTML error: %w", err)
	}

	pattern := `<div class=" num\d* yzphlist hei"><a href="(.*?)".*?>(?:<span class="icon_yuan2"></span>)?(.*?)</a></div>`
	matched := utils.ExtractMatches(pageContent, pattern)

	obj := make([]map[string]interface{}, 0, len(matched))
	for index, item := range matched {
		if len(item) >= 3 {
			obj = append(obj, utils.BuildItem(index+1, item[2], item[1]))
		}
	}

	return utils.BuildSuccessResponse("360doc", "http://www.360doc.com/favicon.ico", obj), nil
}
