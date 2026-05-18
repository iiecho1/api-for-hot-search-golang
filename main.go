package main

import (
	"api/all"
	"api/app"
	"api/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := getEnv("PORT", "1111")
	releaseMode := getEnv("RELEASE", "false") == "true"

	if releaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	registerRoutes(r)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
			"time":   time.Now().Unix(),
			"port":   port,
		})
	})

	addr := ":" + port
	log.Printf("服务器启动: http://localhost%s (环境: %s)", addr, getEnv("ENV", "development"))
	if err := r.Run(addr); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

func registerRoutes(r *gin.Engine) {
	handler := func(cacheKey string, fn func() (map[string]interface{}, error)) gin.HandlerFunc {
		return func(c *gin.Context) {
			result, err := utils.WithCache(cacheKey, utils.DefaultCacheTTL, fn)
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

	routes := map[string]gin.HandlerFunc{
		"/bilibili":   handler("bilibili", app.Bilibili),
		"/360search":  handler("360search", app.Search360),
		"/acfun":      handler("acfun", app.Acfun),
		"/csdn":       handler("csdn", app.CSDN),
		"/dongqiudi":  handler("dongqiudi", app.Dongqiudi),
		"/douban":     handler("douban", app.Douban),
		"/douyin":     handler("douyin", app.Douyin),
		"/github":     handler("github", app.Github),
		"/guojiadili": handler("guojiadili", app.Guojiadili),
		"/history":    handler("history", app.History),
		"/hupu":       handler("hupu", app.Hupu),
		"/ithome":     handler("ithome", app.Ithome),
		"/lishipin":   handler("lishipin", app.Lishipin),
		"/pengpai":    handler("pengpai", app.Pengpai),
		"/qqnews":     handler("qqnews", app.Qqnews),
		"/shaoshupai": handler("shaoshupai", app.Shaoshupai),
		"/sougou":     handler("sougou", app.Sougou),
		"/toutiao":    handler("toutiao", app.Toutiao),
		"/v2ex":       handler("v2ex", app.V2ex),
		"/wangyinews": handler("wangyinews", app.WangyiNews),
		"/weibo":      handler("weibo", app.WeiboHot),
		"/xinjingbao": handler("xinjingbao", app.Xinjingbao),
		"/zhihu":      handler("zhihu", app.Zhihu),
		"/quark":      handler("quark", app.Quark),
		"/souhu":      handler("souhu", app.Souhu),
		"/baidu":      handler("baidu", app.Baidu),
		"/renmin":     handler("renmin", app.Renminwang),
		"/nanfang":    handler("nanfang", app.Nanfangzhoumo),
		"/36kr":       handler("36kr", app.Kr36),
		"/cctv":       handler("cctv", app.CCTV),
		"/tieba":      handler("tieba", app.Tieba),
		"/all":        allHandler(),
	}

	for path, h := range routes {
		r.GET(path, h)
	}
}

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
