package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type bilibiliResponse struct {
	Data bilibiliList `json:"data"`
}

type bilibiliList struct {
	List []bilibiliData `json:"list"`
}
type bilibiliData struct {
	Title string `json:"title"`
	Bvid  string `json:"bvid"`
}

func Bilibili() map[string]interface{} {
	url := "https://api.bilibili.com/x/web-interface/ranking/v2?rid=0&type=all"

	req, err := http.NewRequest("GET", url, nil)
	utils.HandleError(err, "http.Get")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)

	utils.HandleError(err, "http.Get")
	defer resp.Body.Close()

	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll")
	var resultMap bilibiliResponse
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal error")

	// 检查 resultMap["data"] 是否存在且类型正确
	if len(resultMap.Data.List) == 0 {
		return map[string]interface{}{
			"code":    500,
			"message": "API返回数据为空或格式不正确，实际返回数据：" + fmt.Sprintf("%+v", resultMap.Data),
		}
	}

	// data := resultMap["data"].(map[string]interface{})["list"]
	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "哔哩哔哩"
	var obj []map[string]interface{}
	for index, item := range resultMap.Data.List {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.Title
		result["url"] = "https://www.bilibili.com/video/" + fmt.Sprint(item.Bvid)
		obj = append(obj, result)
	}
	api["obj"] = obj
	api["icon"] = "https://www.bilibili.com/favicon.ico" // 32 x 32
	return api
}
