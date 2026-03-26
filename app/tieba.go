package app

import (
	"api/utils"
	"fmt"
)

type tiebaResponse struct {
	Data tiebaData `json:"data"`
}

type tiebaData struct {
	HotTopicList []tiebaTopic `json:"hot_topic_list"`
}

type tiebaTopic struct {
	TopicName  string `json:"topic_name"`
	TopicID    int64  `json:"topic_id"`
	DiscussNum int64  `json:"discuss_num"`
}

func Tieba() (map[string]interface{}, error) {
	url := "https://tieba.baidu.com/c/f/pc/homeSidebarRight?subapp_type=pc&_client_type=20&sign=e9b101df871c39eedcf9a232c2d26ec8"

	var resultMap tiebaResponse
	if err := utils.FetchJSON(url, &resultMap, nil); err != nil {
		return nil, fmt.Errorf("FetchJSON error: %w", err)
	}

	if len(resultMap.Data.HotTopicList) == 0 {
		return utils.BuildErrorResponse("百度贴吧", "https://tieba.baidu.com/favicon.ico",
			"API返回数据为空"), nil
	}

	obj := make([]map[string]interface{}, 0, len(resultMap.Data.HotTopicList))
	for index, item := range resultMap.Data.HotTopicList {
		hot := float64(item.DiscussNum) / 10000
		url := fmt.Sprintf("https://tieba.baidu.com/f?kw=%s&ie=utf-8", item.TopicName)

		obj = append(obj, utils.BuildItem(index+1, item.TopicName, url,
			map[string]string{"hotValue": fmt.Sprintf("%.1f万", hot)}))
	}

	return utils.BuildSuccessResponse("百度贴吧", "https://tieba.baidu.com/favicon.ico", obj), nil
}
