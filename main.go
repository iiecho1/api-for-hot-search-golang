package main

import (
	"api/all"
	"api/app"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 创建路由映射，使用包装函数
	routes := map[string]func(c *gin.Context){
		"/bilibili":   createHandler(app.Bilibili),
		"/360search":  createHandler(app.Search360),
		"/acfun":      createHandler(app.Acfun),
		"/csdn":       createHandler(app.CSDN),
		"/dongqiudi":  createHandler(app.Dongqiudi),
		"/douban":     createHandler(app.Douban),
		"/douyin":     createHandler(app.Douyin),
		"/github":     createHandler(app.Github),
		"/guojiadili": createHandler(app.Guojiadili),
		"/history":    createHandler(app.History),
		"/hupu":       createHandler(app.Hupu),
		"/ithome":     createHandler(app.Ithome),
		"/lishipin":   createHandler(app.Lishipin),
		"/pengpai":    createHandler(app.Pengpai),
		"/qqnews":     createHandler(app.Qqnews),
		"/shaoshupai": createHandler(app.Shaoshupai),
		"/sougou":     createHandler(app.Sougou),
		"/toutiao":    createHandler(app.Toutiao),
		"/v2ex":       createHandler(app.V2ex),
		"/wangyinews": createHandler(app.WangyiNews),
		"/weibo":      createHandler(app.WeiboHot),
		"/xinjingbao": createHandler(app.Xinjingbao),
		"/zhihu":      createHandler(app.Zhihu),
		"/kuake":      createHandler(app.Quark),
		"/souhu":      createHandler(app.Souhu),
		"/baidu":      createHandler(app.Baidu),
		"/renmin":     createHandler(app.Renminwang),
		"/nanfang":    createHandler(app.Nanfangzhoumo),
		"/360doc":     createHandler(app.Doc360),
		"/cctv":       createHandler(app.CCTV),
		"/all":        allHandler, // 单独处理 all，因为它签名不同
	}

	// 注册路由
	for path, handler := range routes {
		r.GET(path, handler)
	}

	r.Run(":1111")
}

// 创建通用处理器函数，处理返回 error 的函数
func createHandler(fn func() (map[string]interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := fn()
		if err != nil {
			// 发生错误时返回错误响应
			c.JSON(500, map[string]interface{}{
				"code":    500,
				"message": "服务器内部错误: " + err.Error(),
				"obj":     []map[string]interface{}{},
			})
			return
		}

		// 检查 API 返回的 code
		if code, ok := result["code"].(int); ok && code != 200 {
			// API 返回业务错误
			c.JSON(code, result)
			return
		}

		// 成功返回
		c.JSON(200, result)
	}
}

// allHandler 专门处理 all.All() 函数
func allHandler(c *gin.Context) {
	result := all.All()

	// 检查 API 返回的 code
	if code, ok := result["code"].(int); ok && code != 200 {
		// API 返回业务错误
		c.JSON(code, result)
		return
	}

	// 成功返回
	c.JSON(200, result)
}
