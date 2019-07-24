package v2ex

import (
	"fmt"
	"spider/analysis"
	"spider/spidercore"
	"spider/public"
	"time"
)

func RunV2ex(url string) {
	fmt.Println(url)
	content := spidercore.Request("GET", url)
	if content == "" {
		fmt.Println("爬取数据为空")
		// 睡眠5分钟
		time.Sleep(time.Duration(300) * time.Second)
		RunV2ex(url)
	}

	data, err := analysis.AnalysisHot(content)
	if err != nil {
		fmt.Println(err)
		// 睡眠5分钟
		time.Sleep(time.Duration(300) * time.Second)
		RunV2ex(url)
	}

	public.Savedata2RD(data, "v2ex")
}