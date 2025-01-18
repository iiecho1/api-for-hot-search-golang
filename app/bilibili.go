package app

import (
	"api/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
	var resultMap map[string]interface{}
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal error")

	// 检查 resultMap["data"] 是否存在且类型正确
	data, ok := resultMap["data"]
	if !ok || data == nil {
		return map[string]interface{}{
			"code":    500,
			"message": "API 返回的数据格式不正确",
		}
	}

	// 检查 data 是否为 map[string]interface{} 类型
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return map[string]interface{}{
			"code":    500,
			"message": "API 返回的 data 字段格式不正确",
		}
	}

	// 检查 dataMap["list"] 是否存在且类型正确
	list, ok := dataMap["list"]
	if !ok || list == nil {
		return map[string]interface{}{
			"code":    500,
			"message": "API 返回的 list 字段格式不正确",
		}
	}

	// 检查 list 是否为 []interface{} 类型
	listSlice, ok := list.([]interface{})
	if !ok {
		return map[string]interface{}{
			"code":    500,
			"message": "API 返回的 list 字段不是数组类型",
		}
	}

	// data := resultMap["data"].(map[string]interface{})["list"]
	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "哔哩哔哩"
	var obj []map[string]interface{}
	for index, item := range listSlice {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["title"]
		result["url"] = "https://www.bilibili.com/video/" + fmt.Sprint(item.(map[string]interface{})["bvid"])
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
