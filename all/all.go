package all

import (
	"api/app"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const defaultTimeout = 8 * time.Second

// sources 定义所有数据源
var sources = map[string]func() (map[string]interface{}, error){
	"360搜索":   app.Search360,
	"36氪":     app.Kr36,
	"哔哩哔哩":   app.Bilibili,
	"AcFun":    app.Acfun,
	"CSDN":     app.CSDN,
	"懂球帝":    app.Dongqiudi,
	"豆瓣":      app.Douban,
	"抖音":      app.Douyin,
	"GitHub":   app.Github,
	"国家地理":   app.Guojiadili,
	"历史上的今天": app.History,
	"虎扑":      app.Hupu,
	"IT之家":   app.Ithome,
	"梨视频":    app.Lishipin,
	"澎湃新闻":   app.Pengpai,
	"腾讯新闻":   app.Qqnews,
	"少数派":    app.Shaoshupai,
	"搜狗":      app.Sougou,
	"今日头条":   app.Toutiao,
	"V2EX":     app.V2ex,
	"网易新闻":   app.WangyiNews,
	"微博":      app.WeiboHot,
	"新京报":    app.Xinjingbao,
	"知乎":      app.Zhihu,
	"夸克":      app.Quark,
	"搜狐":      app.Souhu,
	"百度":      app.Baidu,
	"百度贴吧":   app.Tieba,
	"人民网":    app.Renminwang,
	"南方周末":   app.Nanfangzhoumo,
	"360doc":   app.Doc360,
	"CCTV新闻":  app.CCTV,
}

func All() map[string]interface{} {
	start := time.Now()
	totalSources := len(sources)

	allResult := make(map[string]interface{})
	var mu sync.Mutex
	var wg sync.WaitGroup

	sem := make(chan struct{}, 10)

	for key, fn := range sources {
		wg.Add(1)
		sem <- struct{}{}
		go func(k string, f func() (map[string]interface{}, error)) {
			defer wg.Done()
			defer func() { <-sem }()

			done := make(chan struct{}, 1)
			var result map[string]interface{}
			var err error

			go func() {
				result, err = f()
				done <- struct{}{}
			}()

			select {
			case <-time.After(defaultTimeout):
				fmt.Printf("⚠️  %s 请求超时\n", k)
			case <-done:
				if err != nil {
					fmt.Printf("❌ %s 请求失败: %v\n", k, err)
					return
				}
				if code, ok := result["code"].(int); ok && code == 200 {
					mu.Lock()
					allResult[k] = result["obj"]
					mu.Unlock()
				}
			}
		}(key, fn)
	}

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("✅ 完成 /all 聚合: %d/%d 个源成功, 耗时 %v\n", len(allResult), totalSources, elapsed)

	return map[string]interface{}{
		"code": 200,
		"obj":  allResult,
		"meta": map[string]interface{}{
			"total":   totalSources,
			"success": len(allResult),
			"elapsed": elapsed.String(),
		},
	}
}

// AllHandler 返回 /all 路由的 Gin 处理器
func AllHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := All()
		c.JSON(200, result)
	}
}
