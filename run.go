package main

import (
	"spider/v2ex"
	"time"
)

func main() {
	i := 0
	for {
		go v2ex.RunV2ex("https://www.v2ex.com/?tab=hot", "v2ex-hot")

		go v2ex.RunV2ex("https://www.v2ex.com/?tab=jobs", "v2ex-jobs")

		nowTime := time.Now().Hour()
		if nowTime >= 0 || nowTime < 9 {
			// 夜晚睡眠60分钟
			time.Sleep(time.Duration(3600) * time.Second)
		} else if nowTime >= 9 || nowTime < 11 {
			time.Sleep(time.Duration(900) * time.Second)
		} else if nowTime >= 11 || nowTime < 22 {
			time.Sleep(time.Duration(480) * time.Second)
		} else {
			// 睡眠15分钟
			time.Sleep(time.Duration(900) * time.Second)
		}

		if i > 10 {
			break
		}
	}
}
