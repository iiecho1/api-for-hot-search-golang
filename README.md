🤖 本项目由 [OpenClaw](https://github.com/openclaw/openclaw) AI 助手完成代码重构与优化。（AI太强了，旧代码在old分支）

# 热搜 API（Go）

基于 Gin 框架的热搜聚合 API，支持 30+ 平台。为 [hot_searches_for_apps](https://github.com/iiecho1/hot_searches_for_apps) 提供数据接口。

## 支持平台（30+）

搜索/门户：百度、搜狗、360搜索、搜狐、夸克
社交/社区：微博、知乎、V2EX、虎扑、豆瓣、AcFun、百度贴吧
新闻资讯：今日头条、澎湃新闻、新京报、网易新闻、腾讯新闻、人民网、南方周末、CCTV新闻
科技：CSDN、GitHub、IT之家
视频/娱乐：哔哩哔哩、抖音、梨视频
其他：少数派、懂球帝、国家地理、历史上的今天、360doc

## 快速开始

```bash
go mod tidy
go run main.go
# 服务启动在 http://localhost:1111
```

## API 接口

| 路径 | 说明 |
|:--|:--|
| `/weibo` | 微博热搜 |
| `/zhihu` | 知乎热榜 |
| `/baidu` | 百度热搜 |
| `/bilibili` | B站热门 |
| `/douyin` | 抖音热榜 |
| `/36kr` | 36氪快讯 |
| `/tieba` | 百度贴吧 |
| `/github` | GitHub 热门 |
| `/all` | 全部平台聚合 |
| `/` | 平台列表 |

## 环境变量

| 变量 | 默认值 | 说明 |
|:--|:--|:--|
| `PORT` | 1111 | 服务端口 |
| `RELEASE` | false | 生产模式 |
