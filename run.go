package main

import (
	"spider/v2ex"
	"time"
)

func main() {
	i := 0
	for {
		go v2ex.RunV2ex("https://www.v2ex.com/?tab=hot")

		// 睡眠15分钟
		time.Sleep(time.Duration(900) * time.Second)

		if i > 10 {
			break
		}
	}
}
