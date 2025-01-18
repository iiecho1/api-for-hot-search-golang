package app

import (
	"api/utils"
	"encoding/json"
	"io"
	"net/http"
)

func Quark() map[string]interface{} {
	url := "https://biz.quark.cn/api/trending/ranking/getNewsRanking?modules=hotNews&uc_param_str=dnfrpfbivessbtbmnilauputogpintnwmtsvcppcprsnnnchmicckpgixsnx"
	resp, err := http.Get(url)
	utils.HandleError(err, "http.Get error")
	defer resp.Body.Close()
	pageBytes, err := io.ReadAll(resp.Body)
	utils.HandleError(err, "io.ReadAll error")
	var resultMap map[string]interface{}
	err = json.Unmarshal(pageBytes, &resultMap)
	utils.HandleError(err, "json.Unmarshal error")

	data := resultMap["data"].(map[string]interface{})["hotNews"].(map[string]interface{})["item"]

	api := make(map[string]interface{})
	api["code"] = 200
	api["message"] = "夸克"
	var obj []map[string]interface{}

	for index, item := range data.([]interface{}) {
		result := make(map[string]interface{})
		result["index"] = index + 1
		result["title"] = item.(map[string]interface{})["title"]
		result["url"] = item.(map[string]interface{})["url"]
		result["hotValue"] = item.(map[string]interface{})["hot"]
		obj = append(obj, result)
	}
	api["obj"] = obj
	return api
}
