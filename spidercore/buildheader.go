package spidercore

import (
	"fmt"
	"io"
	"net/http"
)

type Header struct {
	methods string
	header
}

func buildHeaders(methods string, url string, body io.Reader) (*Request, error)  {
	request, err := http.NewRequest(methods, url, body)
	if err != nil {
		fmt.Println("http.NewRequest error")
		return nil, err
	}
	return request, nil
}

func BuildGetHeaders(methods string, url string) {

}