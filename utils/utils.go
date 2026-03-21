package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// Item 表示一个热搜条目
type Item struct {
	Index    int    `json:"index"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	HotValue string `json:"hotValue,omitempty"`
	Desc     string `json:"desc,omitempty"`
}

// ExtractMatches 使用正则表达式提取匹配内容
func ExtractMatches(text, pattern string) [][]string {
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(text, -1)
	return matches
}

// NewClient 创建一个带指定超时的 HTTP 客户端
func NewClient(timeout time.Duration) *http.Client {
	if timeout == 0 {
		timeout = 10 * time.Second
	}
	return &http.Client{Timeout: timeout}
}

// DefaultClient 返回默认的 HTTP 客户端（10秒超时，优先使用IPv4）
func DefaultClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				DualStack: false, // 禁用 IPv6，避免部分 API 在 IPv6 下返回空数据
			}).DialContext,
		},
	}
}

// GetHTML 获取网页内容，支持可选的自定义客户端
func GetHTML(url string, client ...*http.Client) (string, error) {
	c := DefaultClient()
	if len(client) > 0 && client[0] != nil {
		c = client[0]
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; HotSearchBot/1.0)")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	resp, err := c.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP请求失败: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %w", err)
	}
	return string(body), nil
}

// GetJSON 获取 JSON 数据并解析为 map
func GetJSON(url string, client ...*http.Client) (map[string]interface{}, error) {
	c := DefaultClient()
	if len(client) > 0 && client[0] != nil {
		c = client[0]
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; HotSearchBot/1.0)")
	req.Header.Set("Accept", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP请求失败: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode)
	}
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("JSON解析失败: %w", err)
	}
	return data, nil
}

// FetchHTML 获取网页内容，支持自定义请求头
func FetchHTML(url string, headers map[string]string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := DefaultClient().Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP请求失败: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %w", err)
	}
	return string(body), nil
}

// FetchJSON 获取 JSON 数据并解码到指定的目标对象，支持自定义请求头
func FetchJSON(url string, target interface{}, headers map[string]string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := DefaultClient().Do(req)
	if err != nil {
		return fmt.Errorf("HTTP请求失败: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("JSON解析失败: %w", err)
	}
	return nil
}

// CleanHotValue 清理热度值（移除非数字字符）
func CleanHotValue(s string) string {
	nonDigit := regexp.MustCompile(`[^\d]`)
	return strings.TrimSpace(nonDigit.ReplaceAllString(s, ""))
}

// NormalizeTitle 标准化标题（压缩空白字符）
func NormalizeTitle(s string) string {
	whitespace := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(whitespace.ReplaceAllString(s, " "))
}

// BuildItem 创建一个热搜条目
func BuildItem(index int, title, url string, opts ...map[string]string) map[string]interface{} {
	item := map[string]interface{}{
		"index": index,
		"title": title,
		"url":   url,
	}
	if len(opts) > 0 {
		if hotValue, ok := opts[0]["hotValue"]; ok && hotValue != "" {
			item["hotValue"] = hotValue
		}
		if desc, ok := opts[0]["desc"]; ok && desc != "" {
			item["desc"] = desc
		}
	}
	return item
}

// BuildSuccessResponse 构建成功的标准响应
func BuildSuccessResponse(name, icon string, obj interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    200,
		"message": name,
		"icon":    icon,
		"obj":     obj,
	}
}

// BuildErrorResponse 构建错误的标准响应
func BuildErrorResponse(name, icon, errMsg string) map[string]interface{} {
	return map[string]interface{}{
		"code":    500,
		"message": errMsg,
		"icon":    icon,
		"obj":     []map[string]interface{}{},
	}
}
