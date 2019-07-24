package v2ex

import (
	"fmt"
	"spider/analysis"
	"spider/spidercore"
	"spider/public"
)

func RunV2ex(url string) {
	fmt.Println(url)
	content := spidercore.Request("GET", url)
	if content == "" {
		fmt.Println("爬取数据为空")
		return
	}

	data, err := analysis.AnalysisHot(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	public.Savedata2RD(data, "v2ex")
}