package main

import (
	"fmt"
	"spider/v2ex"
	"time"
)

func run(url string, keyName string, ch chan int) {
	i := 0
	for {
		go v2ex.RunV2ex(url, keyName)

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
			ch <- 1
		}
	}
}

func main() {
	//i := 0
	//for {
	ch := make(chan int, 2)

	go run("https://www.v2ex.com/?tab=hot", "v2ex-hot", ch)
	go run("https://www.v2ex.com/?tab=jobs", "v2ex-jobs", ch)

	for i := 0; i < 2; i++ {
		<-ch
	}
	fmt.Println("over")
}
