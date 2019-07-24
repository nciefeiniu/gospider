package spidercore

import (
	"fmt"
	"io"
	"net/http"
)

func buildHeaders(methods string, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(methods, url, body)
	if err != nil {
		fmt.Println("http.NewRequest error")
		return nil, err
	}
	return request, nil
}

func BuildGetHeaders(url string) (*http.Request, error) {
	request, err := buildHeaders("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	request.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	request.Header.Add("authority", url)
	request.Header.Add("method", "GET")
	request.Header.Add("referer", "https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu&wd=v2ex&oq="+
		"go%2520http.Get%25E8%25AE%25BE%25E7%25BD%25AE%25E8%25AF%25B7%25E6%25B1%2582%25E5%25A4%25B4&rsv_pq=e981f1"+
		"82002388a0&rsv_t=9aa6EdB28ccCWuFn5vENP5ix%2Bn%2BsL7Tglh%2F9pFF%2BOzC%2FcVClQ6vSXgj0t8Q&rqlang=cn&rsv_en"+
		"ter=1&rsv_sug3=5&rsv_sug1=5&rsv_sug7=100&bs=go%20http.Get%E8%AE%BE%E7%BD%AE%E8%AF%B7%E6%B1%82%E5%A4%B4")
	request.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 "+
		"(KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36")
	request.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,"+
		"image/apng,*/*;q=0.8,application/signed-exchange;v=b3")

	return request, nil
}
