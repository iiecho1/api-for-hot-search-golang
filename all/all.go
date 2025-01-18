package all

import (
	"api/app"
	"fmt"
	"sync"
)

func All() map[string]interface{} {
	// 定义一个函数列表，方便循环调用
	funcs := map[string]func() map[string]interface{}{
		"360搜索":  app.Search360,
		"哔哩哔哩":   app.Bilibili,
		"AcFun":  app.Acfun,
		"CSDN":   app.CSDN,
		"懂球帝":    app.Dongqiudi,
		"豆瓣":     app.Douban,
		"抖音":     app.Douyin,
		"GitHub": app.Github,
		"国家地理":   app.Guojiadili,
		"历史上的今天": app.History,
		"虎扑":     app.Hupu,
		"IT之家":   app.Ithome,
		"梨视频":    app.Lishipin,
		"澎湃新闻":   app.Pengpai,
		"腾讯新闻":   app.Qqnews,
		"少数派":    app.Shaoshupai,
		"搜狗":     app.Sougou,
		"今日头条":   app.Toutiao,
		"V2EX":   app.V2ex,
		"网易新闻":   app.WangyiNews,
		"微博":     app.WeiboHot,
		"新京报":    app.Xinjingbao,
		"知乎":     app.Zhihu,
		"夸克":     app.Quark,
		"搜狐":     app.Souhu,
		"百度":     app.Baidu,
		"人民网":    app.Renminwang,
	}

	allResult := make(map[string]interface{})

	// 使用并发来调用各个函数
	var wg sync.WaitGroup
	var mu sync.Mutex

	for key, fn := range funcs {
		wg.Add(1)
		go func(k string, f func() map[string]interface{}) {
			defer wg.Done()
			result := f()

			// 检查是否调用成功（假设成功时 code 为 200）
			if result["code"] == 200 {
				mu.Lock()
				allResult[k] = result["obj"]
				mu.Unlock()
			}
			// 如果失败，则不添加到 allResult 中
		}(key, fn)
	}

	wg.Wait()
	fmt.Println(len(allResult))
	return map[string]interface{}{
		"code": 200,
		"obj":  allResult,
	}
}
