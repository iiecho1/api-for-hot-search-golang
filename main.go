package main

import (
	"api/all"
	"api/app"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 定义路由映射
	routes := map[string]func() map[string]interface{}{
		"/bilibili":   app.Bilibili,
		"/360search":  app.Search360,
		"/acfun":      app.Acfun,
		"/csdn":       app.CSDN,
		"/dongqiudi":  app.Dongqiudi,
		"/douban":     app.Douban,
		"/douyin":     app.Douyin,
		"/github":     app.Github,
		"/guojiadili": app.Guojiadili,
		"/history":    app.History,
		"/hupu":       app.Hupu,
		"/ithome":     app.Ithome,
		"/lishipin":   app.Lishipin,
		"/pengpai":    app.Pengpai,
		"/qqnews":     app.Qqnews,
		"/shaoshupai": app.Shaoshupai,
		"/sougou":     app.Sougou,
		"/toutiao":    app.Toutiao,
		"/v2ex":       app.V2ex,
		"/wangyinews": app.WangyiNews,
		"/weibo":      app.WeiboHot,
		"/xinjingbao": app.Xinjingbao,
		"/zhihu":      app.Zhihu,
		"/kuake":      app.Quark,
		"/souhu":      app.Souhu,
		"/baidu":      app.Baidu,
		"/renmin":     app.Renminwang,
		"/all":        all.All,
	}

	// 注册路由
	for path, handler := range routes {
		r.GET(path, func(c *gin.Context) {
			c.JSON(200, handler())
		})
	}

	r.Run(":1111")
}
