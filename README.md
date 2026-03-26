🤖 本项目由 [OpenClaw](https://github.com/openclaw/openclaw) AI 助手完成代码重构与优化。（AI太强了，旧代码在old分支）

# api-for-hot-search-golang

基于 Gin 框架的热搜聚合 API，支持 30+ 平台。为 [hot_searches_for_apps](https://github.com/iiecho1/hot_searches_for_apps) 提供数据接口。

## 快速开始

```bash
go mod tidy
go run main.go
# 服务启动在 http://localhost:1111
```

## API 接口

每个平台一个 GET 路径，返回 JSON 格式热榜数据。

| 分类 | 路径 | 说明 |
|:--|:--|:--|
| 搜索/门户 | `/baidu` | 百度热搜 |
| | `/sougou` | 搜狗热榜 |
| | `/360search` | 360搜索热榜 |
| | `/souhu` | 搜狐热榜 |
| | `/kuake` | 夸克热榜 |
| 社交/社区 | `/weibo` | 微博热搜 |
| | `/zhihu` | 知乎热榜 |
| | `/v2ex` | V2EX |
| | `/hupu` | 虎扑热帖 |
| | `/douban` | 豆瓣热榜 |
| | `/acfun` | AcFun |
| | `/tieba` | 百度贴吧 |
| 新闻资讯 | `/toutiao` | 今日头条 |
| | `/pengpai` | 澎湃新闻 |
| | `/xinjingbao` | 新京报 |
| | `/wangyinews` | 网易新闻 |
| | `/qqnews` | 腾讯新闻 |
| | `/renmin` | 人民网 |
| | `/nanfang` | 南方周末 |
| | `/cctv` | CCTV新闻 |
| 科技 | `/csdn` | CSDN |
| | `/github` | GitHub 热门 |
| | `/ithome` | IT之家 |
| | `/36kr` | 36氪快讯 |
| 视频/娱乐 | `/bilibili` | B站热门 |
| | `/douyin` | 抖音热榜 |
| | `/lishipin` | 梨视频 |
| | `/dongqiudi` | 懂车帝 |
| 其他 | `/shaoshupai` | 少数派 |
| | `/guojiadili` | 国家地理 |
| | `/history` | 历史上的今天 |
| | `/360doc` | 360doc |
| 聚合 | `/all` | 全部平台聚合 |

返回格式统一：

```json
{
  "code": 200,
  "message": "成功",
  "data": [
    { "title": "...", "url": "..." }
  ]
}
```

## 环境变量

| 变量 | 默认值 | 说明 |
|:--|:--|:--|
| `PORT` | 1111 | 服务端口 |
| `RELEASE` | false | 生产模式 |

## 相关项目

- [hot_searches_for_apps](https://github.com/iiecho1/hot_searches_for_apps) — 热搜归档脚本（每日定时拉取）
