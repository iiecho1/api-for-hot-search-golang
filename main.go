package main

import (
	"api/all"
	"api/app"
	"api/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	port := getEnv("PORT", "1111")
	releaseMode := getEnv("RELEASE", "false") == "true"

	if releaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	registerRoutes(r)

	// 健康检查端点
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
			"time":   time.Now().Unix(),
			"port":   port,
		})
	})

	addr := ":" + port
	log.Printf("🔥 服务器启动: http://localhost%s (环境: %s)", addr, getEnv("ENV", "development"))
	if err := r.Run(addr); err != nil {
		log.Fatalf("❌ 服务器启动失败: %v", err)
	}
}

func registerRoutes(r *gin.Engine) {
	// 通用处理器工厂
	handler := func(fn func() (map[string]interface{}, error)) gin.HandlerFunc {
		return func(c *gin.Context) {
			result, err := fn()
			if err != nil {
				c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
					"服务器内部错误", "", err.Error()))
				return
			}
			if code, ok := result["code"].(int); ok && code != 200 {
				c.JSON(code, result)
				return
			}
			c.JSON(http.StatusOK, result)
		}
	}

	// 注册所有应用路由
	routes := map[string]func(c *gin.Context){
		"/bilibili":   handler(app.Bilibili),
		"/360search":  handler(app.Search360),
		"/acfun":      handler(app.Acfun),
		"/csdn":       handler(app.CSDN),
		"/dongqiudi":  handler(app.Dongqiudi),
		"/douban":     handler(app.Douban),
		"/douyin":     handler(app.Douyin),
		"/github":     handler(app.Github),
		"/guojiadili": handler(app.Guojiadili),
		"/history":    handler(app.History),
		"/hupu":       handler(app.Hupu),
		"/ithome":     handler(app.Ithome),
		"/lishipin":   handler(app.Lishipin),
		"/pengpai":    handler(app.Pengpai),
		"/qqnews":     handler(app.Qqnews),
		"/shaoshupai": handler(app.Shaoshupai),
		"/sougou":     handler(app.Sougou),
		"/toutiao":    handler(app.Toutiao),
		"/v2ex":       handler(app.V2ex),
		"/wangyinews": handler(app.WangyiNews),
		"/weibo":      handler(app.WeiboHot),
		"/xinjingbao": handler(app.Xinjingbao),
		"/zhihu":      handler(app.Zhihu),
		"/kuake":      handler(app.Quark),
		"/souhu":      handler(app.Souhu),
		"/baidu":      handler(app.Baidu),
		"/renmin":     handler(app.Renminwang),
		"/nanfang":    handler(app.Nanfangzhoumo),
		"/360doc":     handler(app.Doc360),
		"/cctv":       handler(app.CCTV),
		"/all":        allHandler(),
	}

	for path, h := range routes {
		r.GET(path, h)
	}
}

// allHandler 创建 /all 的聚合处理器
func allHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := all.All()
		if code, ok := result["code"].(int); ok && code != 200 {
			c.JSON(code, result)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
