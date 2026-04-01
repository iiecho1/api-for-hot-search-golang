# 🔥 Hot Search API

> ⚠️ **重要通知**：360doc 将于 2026年5月1日闭站，自 2026年4月1日起不再发布新内容。该数据源4月1日起停止更新。

基于 **Gin** 框架的高性能热搜聚合 API 服务，支持 30+ 中外主流平台实时数据抓取。为 [hot_searches_for_apps](https://github.com/iiecho1/hot_searches_for_apps) 提供数据接口。

## ✨ 特性

- **32个数据源** — 百度、微博、抖音、知乎、GitHub 等主流平台
- **高并发聚合** — `/all` 接口 goroutine 并发抓取，8 秒超时
- **统一工具库** — 共享 HTTP 客户端、JSON 解析、响应构建
- **零外部依赖** — 编译为单个二进制，开箱即用
- **健康检查** — 内置 `/health` 端点

## 🚀 快速开始

```bash
# 运行
go run main.go

# 或编译后运行
go build -o hot-search-api .
./hot-search-api
```

## 📡 API 端点

| 路径 | 平台 | 路径 | 平台 |
|------|------|------|------|
| `/baidu` | 百度 | `/weibo` | 微博 |
| `/douyin` | 抖音 | `/zhihu` | 知乎 |
| `/bilibili` | 哔哩哔哩 | `/github` | GitHub |
| `/toutiao` | 今日头条 | `/csdn` | CSDN |
| `/v2ex` | V2EX | `/douban` | 豆瓣 |
| `/hupu` | 虎扑 | `/ithome` | IT之家 |
| `/sougou` | 搜狗 | `/qqnews` | 腾讯新闻 |
| `/pengpai` | 澎湃新闻 | `/cctv` | CCTV |
| `/renmin` | 人民网 | `/wangyinews` | 网易新闻 |
| `/acfun` | AcFun | `/dongqiudi` | 懂球帝 |
| `/tieba` | 百度贴吧 | `/36kr` | 36氪 |
| `/lishipin` | 梨视频 | `/shaoshupai` | 少数派 |
| `/souhu` | 搜狐 | `/quark` | 夸克 |
| `/xinjingbao` | 新京报 | `/nanfang` | 南方周末 |
| `/guojiadili` | 国家地理 | `/history` | 历史上的今天 |
| `/360search` | 360搜索 | `/360doc` | 360doc ⚠️ *将于2026年5月1日闭站，不再更新* |
| **`/all`** | **聚合所有源** | `/health` | 健康检查 |

### 响应格式

```json
{
  "code": 200,
  "message": "百度",
  "icon": "https://www.baidu.com/favicon.ico",
  "obj": [
    { "index": 1, "title": "热搜标题", "url": "...", "hotValue": "12345" }
  ]
}
```

## ⚙️ 配置

| 环境变量 | 默认值 | 说明 |
|----------|--------|------|
| `PORT` | `1111` | 监听端口 |
| `RELEASE` | `false` | Gin Release 模式 |
| `ENV` | `development` | 环境标识 |

## 📦 项目结构

```
├── main.go        # 入口 + 路由注册
├── all/all.go     # 聚合逻辑（并发 + 超时控制）
├── app/           # 30+ 数据源实现
└── utils/utils.go # 共享工具（HTTP/JSON/响应构建）
```

## 相关项目

- [hot_searches_for_apps](https://github.com/iiecho1/hot_searches_for_apps) — 热搜归档脚本（每小时定时拉取）

🤖 本项目由 OpenClaw AI 助手完成代码重构与优化。（AI太强了，旧代码在old分支）