package app

import (
	"api/utils"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type csdnResponse struct {
	Data []csdnData `json:"data"`
}

type csdnData struct {
	Title    string `json:"articleTitle"`
	URL      string `json:"articleDetailUrl"`
	HotValue string `json:"pcHotRankScore"`
}

func CSDN() (map[string]interface{}, error) {
	url := "https://blog.csdn.net/phoenix/web/blog/hotRank?&pageSize=100"

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 10 * 1e9,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP请求失败，状态码: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %w", err)
	}

	var resultMap csdnResponse
	if err := json.Unmarshal(body, &resultMap); err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %w", err)
	}

	obj := make([]map[string]interface{}, 0, len(resultMap.Data))
	for index, item := range resultMap.Data {
		obj = append(obj, utils.BuildItem(index+1, item.Title, item.URL,
			map[string]string{"hotValue": item.HotValue}))
	}

	return utils.BuildSuccessResponse("CSDN", "https://csdnimg.cn/public/favicon.ico", obj), nil
}
